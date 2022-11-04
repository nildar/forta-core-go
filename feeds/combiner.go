package feeds

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/forta-network/forta-core-go/clients/graphql"
	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var (
	ErrCombinerStopReached   = fmt.Errorf("combiner stop reached")
	DefaultRatelimitDuration = time.Minute
)

type cfHandler struct {
	Handler func(evt *domain.AlertEvent) error
	ErrCh   chan<- error
}
type combinerFeed struct {
	start     uint64
	end       uint64
	ctx       context.Context
	started   bool
	rateLimit *time.Ticker

	lastAlert health.MessageTracker

	alertCh chan *domain.AlertEvent
	client  *graphql.Client

	botSubscriptions map[string][]string
	botsMu           sync.RWMutex
	// bloom filter could be a better choice but the miss rate was too high
	// and we don't know the expected item count
	alertCache *cache.Cache

	handlers   []cfHandler
	handlersMu sync.Mutex
}

// Subscriptions returns alert/combiner bot pairs,
func (cf *combinerFeed) Subscriptions() map[string][]string {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()
	return cf.botSubscriptions
}
func (cf *combinerFeed) SubscribedBots() (bots []string) {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()

	for sub := range cf.botSubscriptions {
		bots = append(bots, sub)
	}
	return
}
func (cf *combinerFeed) AddSubscription(subscription *protocol.CombinerBotSubscription, subscriber string) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()
	subscriptionString := encodeSubscription(subscription)
	for _, s := range cf.botSubscriptions[subscriptionString] {
		if s == subscriber {
			return
		}
	}

	cf.botSubscriptions[subscriptionString] = append(cf.botSubscriptions[subscriptionString], subscriber)
}
func (cf *combinerFeed) RemoveSubscription(subscription *protocol.CombinerBotSubscription, subscriber string) {
	subscriptionString := encodeSubscription(subscription)
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for i, s := range cf.botSubscriptions[subscriptionString] {
		if s == subscriber {
			cf.botSubscriptions[subscriptionString] = append(cf.botSubscriptions[subscriptionString][:i], cf.botSubscriptions[subscriptionString][i+1:]...)
		}
	}
	if len(cf.botSubscriptions[encodeSubscription(subscription)]) == 0 {
		delete(cf.botSubscriptions, subscriptionString)
	}
}
func (cf *combinerFeed) RegisterHandler(alertHandler func(evt *domain.AlertEvent) error) <-chan error {
	cf.handlersMu.Lock()
	defer cf.handlersMu.Unlock()

	errCh := make(chan error)
	cf.handlers = append(
		cf.handlers, cfHandler{
			Handler: alertHandler,
			ErrCh:   errCh,
		},
	)
	return errCh
}

func encodeSubscription(s *protocol.CombinerBotSubscription) string {
	return strings.Join([]string{s.BotId}, ",")
}

type CombinerFeedConfig struct {
	RateLimit *time.Ticker
	APIUrl    string
	Start     uint64
	End       uint64
}

func (cf *combinerFeed) IsStarted() bool {
	return cf.started
}

func (cf *combinerFeed) Start() {
	if !cf.started {
		go cf.loop()
	}
}
func (cf *combinerFeed) initialize() error {
	if cf.start == 0 {
		cf.start = uint64(time.Now().Add(time.Minute * -10).UnixMilli())
	}
	cf.rateLimit = time.NewTicker(DefaultRatelimitDuration)
	return nil
}
func (cf *combinerFeed) StartRange(start uint64, end uint64, rate int64) {
	if !cf.started {
		cf.start = start
		cf.end = end
		go cf.loop()
	}
}

func (cf *combinerFeed) ForEachAlert(alertHandler func(evt *domain.AlertEvent) error) error {
	return cf.forEachAlert([]cfHandler{{Handler: alertHandler}})
}

func (cf *combinerFeed) forEachAlert(alertHandlers []cfHandler) error {
	currentTimestp := cf.start
	for {

		if cf.ctx.Err() != nil {
			return cf.ctx.Err()
		}

		logger := log.WithFields(
			log.Fields{
				"currentTimestamp": currentTimestp,
			},
		)

		if cf.end != 0 && currentTimestp > cf.end {
			logger.Info("end timestamp reached - exiting")
			return ErrCombinerStopReached
		}

		// skip query if there are no alert subscriptions
		if len(cf.SubscribedBots()) == 0 {
			continue
		}
		createdSince := time.Now().UnixMilli() - int64(currentTimestp)
		// TODO: needs changes in alerts-api
		// createdBefore := createdSince - graphql.DefaultLastNMinutes.Milliseconds()

		var alerts []*protocol.AlertEvent
		bo := backoff.NewExponentialBackOff()
		err := backoff.Retry(
			func() error {
				var cErr error
				alerts, cErr = cf.client.GetAlerts(
					cf.ctx,
					&graphql.AlertsInput{
						Bots:         cf.SubscribedBots(),
						CreatedSince: uint(createdSince),
						// CreatedBefore: uint(createdBefore), needs changes in alerts-api
					},
				)
				if cErr != nil {
					log.WithError(cErr).Warn("error retrieving alerts")
					return cErr
				}

				return nil
			}, bo,
		)
		if err != nil {
			return err
		}

		for _, alert := range alerts {
			if _, exists := cf.alertCache.Get(alert.Alert.Hash); exists {
				continue
			}

			cf.alertCache.Set(
				alert.Alert.Hash, struct{}{}, cache.DefaultExpiration,
			)

			alertCA, err := time.Parse(time.RFC3339, alert.Alert.CreatedAt)
			if err != nil {
				return err
			}

			// just for local purposes
			if cf.end != 0 && alertCA.UnixMilli() > int64(cf.end) {
				continue
			}

			evt := &domain.AlertEvent{
				Event: alert,
				Timestamps: &domain.TrackingTimestamps{
					Feed:        time.Now().UTC(),
					SourceAlert: alertCA,
				},
			}

			for _, alertHandler := range alertHandlers {
				if err := alertHandler.Handler(evt); err != nil {
					return err
				}
			}

		}

		currentTimestp += uint64(DefaultRatelimitDuration.Milliseconds())

		if cf.rateLimit != nil {
			<-cf.rateLimit.C
		}
	}
}

// Name returns the name of this implementation.
func (cf *combinerFeed) Name() string {
	return "alert-feed"
}

// Health implements the health.Reporter interface.
func (cf *combinerFeed) Health() health.Reports {
	return health.Reports{
		cf.lastAlert.GetReport("last-alert"),
	}
}

func (cf *combinerFeed) loop() {
	if err := cf.initialize(); err != nil {
		log.WithError(err).Panic("failed to initialize")
	}
	cf.started = true
	defer func() {
		cf.started = false
	}()

	err := cf.forEachAlert(cf.handlers)
	if err == nil {
		return
	}
	if err != ErrCombinerStopReached {
		log.WithError(err).Warn("failed while processing blocks")
	}
	for _, handler := range cf.handlers {
		if handler.ErrCh != nil {
			handler.ErrCh <- err
		}
	}
}

func NewCombinerFeed(ctx context.Context, cfg CombinerFeedConfig) (AlertFeed, error) {
	ac := graphql.NewClient(cfg.APIUrl)
	alerts := make(chan *domain.AlertEvent, 10)

	bf := &combinerFeed{
		start:            cfg.Start,
		end:              cfg.End,
		ctx:              ctx,
		client:           ac,
		rateLimit:        cfg.RateLimit,
		alertCh:          alerts,
		botSubscriptions: map[string][]string{},
		alertCache:       cache.New(graphql.DefaultLastNMinutes*2, time.Minute),
	}
	return bf, nil
}
