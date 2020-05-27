// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeNetworkPoliciesActor struct {
	NetworkPoliciesBySpaceStub        func(string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)
	networkPoliciesBySpaceMutex       sync.RWMutex
	networkPoliciesBySpaceArgsForCall []struct {
		arg1 string
	}
	networkPoliciesBySpaceReturns struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	networkPoliciesBySpaceReturnsOnCall map[int]struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	NetworkPoliciesBySpaceAndAppNameStub        func(string, string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)
	networkPoliciesBySpaceAndAppNameMutex       sync.RWMutex
	networkPoliciesBySpaceAndAppNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	networkPoliciesBySpaceAndAppNameReturns struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	networkPoliciesBySpaceAndAppNameReturnsOnCall map[int]struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpace(arg1 string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error) {
	fake.networkPoliciesBySpaceMutex.Lock()
	ret, specificReturn := fake.networkPoliciesBySpaceReturnsOnCall[len(fake.networkPoliciesBySpaceArgsForCall)]
	fake.networkPoliciesBySpaceArgsForCall = append(fake.networkPoliciesBySpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("NetworkPoliciesBySpace", []interface{}{arg1})
	fake.networkPoliciesBySpaceMutex.Unlock()
	if fake.NetworkPoliciesBySpaceStub != nil {
		return fake.NetworkPoliciesBySpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.networkPoliciesBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceCallCount() int {
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	return len(fake.networkPoliciesBySpaceArgsForCall)
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceCalls(stub func(string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)) {
	fake.networkPoliciesBySpaceMutex.Lock()
	defer fake.networkPoliciesBySpaceMutex.Unlock()
	fake.NetworkPoliciesBySpaceStub = stub
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceArgsForCall(i int) string {
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	argsForCall := fake.networkPoliciesBySpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceReturns(result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.networkPoliciesBySpaceMutex.Lock()
	defer fake.networkPoliciesBySpaceMutex.Unlock()
	fake.NetworkPoliciesBySpaceStub = nil
	fake.networkPoliciesBySpaceReturns = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceReturnsOnCall(i int, result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.networkPoliciesBySpaceMutex.Lock()
	defer fake.networkPoliciesBySpaceMutex.Unlock()
	fake.NetworkPoliciesBySpaceStub = nil
	if fake.networkPoliciesBySpaceReturnsOnCall == nil {
		fake.networkPoliciesBySpaceReturnsOnCall = make(map[int]struct {
			result1 []cfnetworkingaction.Policy
			result2 cfnetworkingaction.Warnings
			result3 error
		})
	}
	fake.networkPoliciesBySpaceReturnsOnCall[i] = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppName(arg1 string, arg2 string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error) {
	fake.networkPoliciesBySpaceAndAppNameMutex.Lock()
	ret, specificReturn := fake.networkPoliciesBySpaceAndAppNameReturnsOnCall[len(fake.networkPoliciesBySpaceAndAppNameArgsForCall)]
	fake.networkPoliciesBySpaceAndAppNameArgsForCall = append(fake.networkPoliciesBySpaceAndAppNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("NetworkPoliciesBySpaceAndAppName", []interface{}{arg1, arg2})
	fake.networkPoliciesBySpaceAndAppNameMutex.Unlock()
	if fake.NetworkPoliciesBySpaceAndAppNameStub != nil {
		return fake.NetworkPoliciesBySpaceAndAppNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.networkPoliciesBySpaceAndAppNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameCallCount() int {
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	return len(fake.networkPoliciesBySpaceAndAppNameArgsForCall)
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameCalls(stub func(string, string) ([]cfnetworkingaction.Policy, cfnetworkingaction.Warnings, error)) {
	fake.networkPoliciesBySpaceAndAppNameMutex.Lock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.Unlock()
	fake.NetworkPoliciesBySpaceAndAppNameStub = stub
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameArgsForCall(i int) (string, string) {
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	argsForCall := fake.networkPoliciesBySpaceAndAppNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameReturns(result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.networkPoliciesBySpaceAndAppNameMutex.Lock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.Unlock()
	fake.NetworkPoliciesBySpaceAndAppNameStub = nil
	fake.networkPoliciesBySpaceAndAppNameReturns = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) NetworkPoliciesBySpaceAndAppNameReturnsOnCall(i int, result1 []cfnetworkingaction.Policy, result2 cfnetworkingaction.Warnings, result3 error) {
	fake.networkPoliciesBySpaceAndAppNameMutex.Lock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.Unlock()
	fake.NetworkPoliciesBySpaceAndAppNameStub = nil
	if fake.networkPoliciesBySpaceAndAppNameReturnsOnCall == nil {
		fake.networkPoliciesBySpaceAndAppNameReturnsOnCall = make(map[int]struct {
			result1 []cfnetworkingaction.Policy
			result2 cfnetworkingaction.Warnings
			result3 error
		})
	}
	fake.networkPoliciesBySpaceAndAppNameReturnsOnCall[i] = struct {
		result1 []cfnetworkingaction.Policy
		result2 cfnetworkingaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworkPoliciesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.networkPoliciesBySpaceMutex.RLock()
	defer fake.networkPoliciesBySpaceMutex.RUnlock()
	fake.networkPoliciesBySpaceAndAppNameMutex.RLock()
	defer fake.networkPoliciesBySpaceAndAppNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworkPoliciesActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.NetworkPoliciesActor = new(FakeNetworkPoliciesActor)
