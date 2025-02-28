// Code generated by counterfeiter. DO NOT EDIT.
package authfakes

import (
	"sync"

	"github.com/whoyao/protocol/auth"
)

type FakeTokenVerifier struct {
	IdentityStub        func() string
	identityMutex       sync.RWMutex
	identityArgsForCall []struct {
	}
	identityReturns struct {
		result1 string
	}
	identityReturnsOnCall map[int]struct {
		result1 string
	}
	VerifyStub        func(interface{}) (*auth.ClaimGrants, error)
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 interface{}
	}
	verifyReturns struct {
		result1 *auth.ClaimGrants
		result2 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 *auth.ClaimGrants
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenVerifier) Identity() string {
	fake.identityMutex.Lock()
	ret, specificReturn := fake.identityReturnsOnCall[len(fake.identityArgsForCall)]
	fake.identityArgsForCall = append(fake.identityArgsForCall, struct {
	}{})
	stub := fake.IdentityStub
	fakeReturns := fake.identityReturns
	fake.recordInvocation("Identity", []interface{}{})
	fake.identityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTokenVerifier) IdentityCallCount() int {
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	return len(fake.identityArgsForCall)
}

func (fake *FakeTokenVerifier) IdentityCalls(stub func() string) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = stub
}

func (fake *FakeTokenVerifier) IdentityReturns(result1 string) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = nil
	fake.identityReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenVerifier) IdentityReturnsOnCall(i int, result1 string) {
	fake.identityMutex.Lock()
	defer fake.identityMutex.Unlock()
	fake.IdentityStub = nil
	if fake.identityReturnsOnCall == nil {
		fake.identityReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.identityReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenVerifier) Verify(arg1 interface{}) (*auth.ClaimGrants, error) {
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.VerifyStub
	fakeReturns := fake.verifyReturns
	fake.recordInvocation("Verify", []interface{}{arg1})
	fake.verifyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTokenVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeTokenVerifier) VerifyCalls(stub func(interface{}) (*auth.ClaimGrants, error)) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = stub
}

func (fake *FakeTokenVerifier) VerifyArgsForCall(i int) interface{} {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	argsForCall := fake.verifyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTokenVerifier) VerifyReturns(result1 *auth.ClaimGrants, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 *auth.ClaimGrants
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenVerifier) VerifyReturnsOnCall(i int, result1 *auth.ClaimGrants, result2 error) {
	fake.verifyMutex.Lock()
	defer fake.verifyMutex.Unlock()
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 *auth.ClaimGrants
			result2 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 *auth.ClaimGrants
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.identityMutex.RLock()
	defer fake.identityMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenVerifier) recordInvocation(key string, args []interface{}) {
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

var _ auth.TokenVerifier = new(FakeTokenVerifier)
