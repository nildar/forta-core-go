// Code generated by go-merge-types. DO NOT EDIT.

package contract_stake_allocator

import (
	import_fmt "fmt"
	import_sync "sync"


	stakeallocator010 "github.com/forta-network/forta-core-go/contracts/generated/contract_stake_allocator_0_1_0"



	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"math/big"

	"github.com/ethereum/go-ethereum/common"

)

// StakeAllocatorCaller is a new type which can multiplex calls to different implementation types.
type StakeAllocatorCaller struct {

	typ0 *stakeallocator010.StakeAllocatorCaller

	currTag string
	mu import_sync.RWMutex
	unsafe bool // default: false
}

// NewStakeAllocatorCaller creates a new merged type.
func NewStakeAllocatorCaller(address common.Address, caller bind.ContractCaller) (*StakeAllocatorCaller, error) {
	var (
		mergedType StakeAllocatorCaller
		err error
	)
	mergedType.currTag = "0.1.0"


	mergedType.typ0, err = stakeallocator010.NewStakeAllocatorCaller(address, caller)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize stakeallocator010.StakeAllocatorCaller: %v", err)
	}


	return &mergedType, nil
}

// IsKnownTag tells if given tag is a known tag.
func IsKnownTag(tag string) bool {

	if tag == "0.1.0" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *StakeAllocatorCaller) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTag(tag) {
		tag = "0.1.0"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *StakeAllocatorCaller) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *StakeAllocatorCaller) Safe() {
	merged.unsafe = false
}




// AllocatedDelegatorsStakePerManaged multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) AllocatedDelegatorsStakePerManaged(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AllocatedDelegatorsStakePerManaged(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.AllocatedDelegatorsStakePerManaged not implemented (tag=%s)", merged.currTag)
	return
}



// AllocatedManagedStake multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) AllocatedManagedStake(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AllocatedManagedStake(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.AllocatedManagedStake not implemented (tag=%s)", merged.currTag)
	return
}



// AllocatedOwnStakePerManaged multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) AllocatedOwnStakePerManaged(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AllocatedOwnStakePerManaged(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.AllocatedOwnStakePerManaged not implemented (tag=%s)", merged.currTag)
	return
}



// AllocatedStakeFor multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) AllocatedStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AllocatedStakeFor(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.AllocatedStakeFor not implemented (tag=%s)", merged.currTag)
	return
}



// AllocatedStakePerManaged multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) AllocatedStakePerManaged(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.AllocatedStakePerManaged(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.AllocatedStakePerManaged not implemented (tag=%s)", merged.currTag)
	return
}



// GetDelegatedSubjectType multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) GetDelegatedSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetDelegatedSubjectType(opts, subjectType)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.GetDelegatedSubjectType not implemented (tag=%s)", merged.currTag)
	return
}



// GetDelegatorSubjectType multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) GetDelegatorSubjectType(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetDelegatorSubjectType(opts, subjectType)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.GetDelegatorSubjectType not implemented (tag=%s)", merged.currTag)
	return
}



// GetSubjectTypeAgency multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) GetSubjectTypeAgency(opts *bind.CallOpts, subjectType uint8) (retVal uint8, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.GetSubjectTypeAgency(opts, subjectType)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.GetSubjectTypeAgency not implemented (tag=%s)", merged.currTag)
	return
}



// IsTrustedForwarder multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (retVal bool, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.IsTrustedForwarder(opts, forwarder)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.IsTrustedForwarder not implemented (tag=%s)", merged.currTag)
	return
}



// ProxiableUUID multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) ProxiableUUID(opts *bind.CallOpts) (retVal [32]byte, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.ProxiableUUID(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.ProxiableUUID not implemented (tag=%s)", merged.currTag)
	return
}



// RewardsDistributor multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) RewardsDistributor(opts *bind.CallOpts) (retVal common.Address, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.RewardsDistributor(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.RewardsDistributor not implemented (tag=%s)", merged.currTag)
	return
}



// UnallocatedStakeFor multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) UnallocatedStakeFor(opts *bind.CallOpts, subjectType uint8, subject *big.Int) (retVal *big.Int, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.UnallocatedStakeFor(opts, subjectType, subject)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.UnallocatedStakeFor not implemented (tag=%s)", merged.currTag)
	return
}



// Version multiplexes to different implementations of the method.
func (merged *StakeAllocatorCaller) Version(opts *bind.CallOpts) (retVal string, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.0" {
		val, methodErr := merged.typ0.Version(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("StakeAllocatorCaller.Version not implemented (tag=%s)", merged.currTag)
	return
}