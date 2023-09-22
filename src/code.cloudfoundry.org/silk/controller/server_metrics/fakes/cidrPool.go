// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type CIDRPool struct {
	BlockPoolSizeStub        func() int
	blockPoolSizeMutex       sync.RWMutex
	blockPoolSizeArgsForCall []struct{}
	blockPoolSizeReturns     struct {
		result1 int
	}
	blockPoolSizeReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CIDRPool) BlockPoolSize() int {
	fake.blockPoolSizeMutex.Lock()
	ret, specificReturn := fake.blockPoolSizeReturnsOnCall[len(fake.blockPoolSizeArgsForCall)]
	fake.blockPoolSizeArgsForCall = append(fake.blockPoolSizeArgsForCall, struct{}{})
	fake.recordInvocation("BlockPoolSize", []interface{}{})
	fake.blockPoolSizeMutex.Unlock()
	if fake.BlockPoolSizeStub != nil {
		return fake.BlockPoolSizeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.blockPoolSizeReturns.result1
}

func (fake *CIDRPool) BlockPoolSizeCallCount() int {
	fake.blockPoolSizeMutex.RLock()
	defer fake.blockPoolSizeMutex.RUnlock()
	return len(fake.blockPoolSizeArgsForCall)
}

func (fake *CIDRPool) BlockPoolSizeReturns(result1 int) {
	fake.BlockPoolSizeStub = nil
	fake.blockPoolSizeReturns = struct {
		result1 int
	}{result1}
}

func (fake *CIDRPool) BlockPoolSizeReturnsOnCall(i int, result1 int) {
	fake.BlockPoolSizeStub = nil
	if fake.blockPoolSizeReturnsOnCall == nil {
		fake.blockPoolSizeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.blockPoolSizeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *CIDRPool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockPoolSizeMutex.RLock()
	defer fake.blockPoolSizeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CIDRPool) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}