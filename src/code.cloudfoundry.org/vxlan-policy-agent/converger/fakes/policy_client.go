// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type PolicyClient struct {
	GetPoliciesLastUpdatedStub        func() (int, error)
	getPoliciesLastUpdatedMutex       sync.RWMutex
	getPoliciesLastUpdatedArgsForCall []struct {
	}
	getPoliciesLastUpdatedReturns struct {
		result1 int
		result2 error
	}
	getPoliciesLastUpdatedReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PolicyClient) GetPoliciesLastUpdated() (int, error) {
	fake.getPoliciesLastUpdatedMutex.Lock()
	ret, specificReturn := fake.getPoliciesLastUpdatedReturnsOnCall[len(fake.getPoliciesLastUpdatedArgsForCall)]
	fake.getPoliciesLastUpdatedArgsForCall = append(fake.getPoliciesLastUpdatedArgsForCall, struct {
	}{})
	stub := fake.GetPoliciesLastUpdatedStub
	fakeReturns := fake.getPoliciesLastUpdatedReturns
	fake.recordInvocation("GetPoliciesLastUpdated", []interface{}{})
	fake.getPoliciesLastUpdatedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PolicyClient) GetPoliciesLastUpdatedCallCount() int {
	fake.getPoliciesLastUpdatedMutex.RLock()
	defer fake.getPoliciesLastUpdatedMutex.RUnlock()
	return len(fake.getPoliciesLastUpdatedArgsForCall)
}

func (fake *PolicyClient) GetPoliciesLastUpdatedCalls(stub func() (int, error)) {
	fake.getPoliciesLastUpdatedMutex.Lock()
	defer fake.getPoliciesLastUpdatedMutex.Unlock()
	fake.GetPoliciesLastUpdatedStub = stub
}

func (fake *PolicyClient) GetPoliciesLastUpdatedReturns(result1 int, result2 error) {
	fake.getPoliciesLastUpdatedMutex.Lock()
	defer fake.getPoliciesLastUpdatedMutex.Unlock()
	fake.GetPoliciesLastUpdatedStub = nil
	fake.getPoliciesLastUpdatedReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) GetPoliciesLastUpdatedReturnsOnCall(i int, result1 int, result2 error) {
	fake.getPoliciesLastUpdatedMutex.Lock()
	defer fake.getPoliciesLastUpdatedMutex.Unlock()
	fake.GetPoliciesLastUpdatedStub = nil
	if fake.getPoliciesLastUpdatedReturnsOnCall == nil {
		fake.getPoliciesLastUpdatedReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.getPoliciesLastUpdatedReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *PolicyClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPoliciesLastUpdatedMutex.RLock()
	defer fake.getPoliciesLastUpdatedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PolicyClient) recordInvocation(key string, args []interface{}) {
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
