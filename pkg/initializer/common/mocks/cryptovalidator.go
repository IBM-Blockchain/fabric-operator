// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/IBM-Blockchain/fabric-operator/pkg/initializer/common"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CryptoValidator struct {
	CheckClientAuthCryptoStub        func(v1.Object, string) error
	checkClientAuthCryptoMutex       sync.RWMutex
	checkClientAuthCryptoArgsForCall []struct {
		arg1 v1.Object
		arg2 string
	}
	checkClientAuthCryptoReturns struct {
		result1 error
	}
	checkClientAuthCryptoReturnsOnCall map[int]struct {
		result1 error
	}
	CheckEcertCryptoStub        func(v1.Object, string) error
	checkEcertCryptoMutex       sync.RWMutex
	checkEcertCryptoArgsForCall []struct {
		arg1 v1.Object
		arg2 string
	}
	checkEcertCryptoReturns struct {
		result1 error
	}
	checkEcertCryptoReturnsOnCall map[int]struct {
		result1 error
	}
	CheckTLSCryptoStub        func(v1.Object, string) error
	checkTLSCryptoMutex       sync.RWMutex
	checkTLSCryptoArgsForCall []struct {
		arg1 v1.Object
		arg2 string
	}
	checkTLSCryptoReturns struct {
		result1 error
	}
	checkTLSCryptoReturnsOnCall map[int]struct {
		result1 error
	}
	SetHSMEnabledStub        func(bool)
	setHSMEnabledMutex       sync.RWMutex
	setHSMEnabledArgsForCall []struct {
		arg1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CryptoValidator) CheckClientAuthCrypto(arg1 v1.Object, arg2 string) error {
	fake.checkClientAuthCryptoMutex.Lock()
	ret, specificReturn := fake.checkClientAuthCryptoReturnsOnCall[len(fake.checkClientAuthCryptoArgsForCall)]
	fake.checkClientAuthCryptoArgsForCall = append(fake.checkClientAuthCryptoArgsForCall, struct {
		arg1 v1.Object
		arg2 string
	}{arg1, arg2})
	stub := fake.CheckClientAuthCryptoStub
	fakeReturns := fake.checkClientAuthCryptoReturns
	fake.recordInvocation("CheckClientAuthCrypto", []interface{}{arg1, arg2})
	fake.checkClientAuthCryptoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CryptoValidator) CheckClientAuthCryptoCallCount() int {
	fake.checkClientAuthCryptoMutex.RLock()
	defer fake.checkClientAuthCryptoMutex.RUnlock()
	return len(fake.checkClientAuthCryptoArgsForCall)
}

func (fake *CryptoValidator) CheckClientAuthCryptoCalls(stub func(v1.Object, string) error) {
	fake.checkClientAuthCryptoMutex.Lock()
	defer fake.checkClientAuthCryptoMutex.Unlock()
	fake.CheckClientAuthCryptoStub = stub
}

func (fake *CryptoValidator) CheckClientAuthCryptoArgsForCall(i int) (v1.Object, string) {
	fake.checkClientAuthCryptoMutex.RLock()
	defer fake.checkClientAuthCryptoMutex.RUnlock()
	argsForCall := fake.checkClientAuthCryptoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CryptoValidator) CheckClientAuthCryptoReturns(result1 error) {
	fake.checkClientAuthCryptoMutex.Lock()
	defer fake.checkClientAuthCryptoMutex.Unlock()
	fake.CheckClientAuthCryptoStub = nil
	fake.checkClientAuthCryptoReturns = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) CheckClientAuthCryptoReturnsOnCall(i int, result1 error) {
	fake.checkClientAuthCryptoMutex.Lock()
	defer fake.checkClientAuthCryptoMutex.Unlock()
	fake.CheckClientAuthCryptoStub = nil
	if fake.checkClientAuthCryptoReturnsOnCall == nil {
		fake.checkClientAuthCryptoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkClientAuthCryptoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) CheckEcertCrypto(arg1 v1.Object, arg2 string) error {
	fake.checkEcertCryptoMutex.Lock()
	ret, specificReturn := fake.checkEcertCryptoReturnsOnCall[len(fake.checkEcertCryptoArgsForCall)]
	fake.checkEcertCryptoArgsForCall = append(fake.checkEcertCryptoArgsForCall, struct {
		arg1 v1.Object
		arg2 string
	}{arg1, arg2})
	stub := fake.CheckEcertCryptoStub
	fakeReturns := fake.checkEcertCryptoReturns
	fake.recordInvocation("CheckEcertCrypto", []interface{}{arg1, arg2})
	fake.checkEcertCryptoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CryptoValidator) CheckEcertCryptoCallCount() int {
	fake.checkEcertCryptoMutex.RLock()
	defer fake.checkEcertCryptoMutex.RUnlock()
	return len(fake.checkEcertCryptoArgsForCall)
}

func (fake *CryptoValidator) CheckEcertCryptoCalls(stub func(v1.Object, string) error) {
	fake.checkEcertCryptoMutex.Lock()
	defer fake.checkEcertCryptoMutex.Unlock()
	fake.CheckEcertCryptoStub = stub
}

func (fake *CryptoValidator) CheckEcertCryptoArgsForCall(i int) (v1.Object, string) {
	fake.checkEcertCryptoMutex.RLock()
	defer fake.checkEcertCryptoMutex.RUnlock()
	argsForCall := fake.checkEcertCryptoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CryptoValidator) CheckEcertCryptoReturns(result1 error) {
	fake.checkEcertCryptoMutex.Lock()
	defer fake.checkEcertCryptoMutex.Unlock()
	fake.CheckEcertCryptoStub = nil
	fake.checkEcertCryptoReturns = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) CheckEcertCryptoReturnsOnCall(i int, result1 error) {
	fake.checkEcertCryptoMutex.Lock()
	defer fake.checkEcertCryptoMutex.Unlock()
	fake.CheckEcertCryptoStub = nil
	if fake.checkEcertCryptoReturnsOnCall == nil {
		fake.checkEcertCryptoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkEcertCryptoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) CheckTLSCrypto(arg1 v1.Object, arg2 string) error {
	fake.checkTLSCryptoMutex.Lock()
	ret, specificReturn := fake.checkTLSCryptoReturnsOnCall[len(fake.checkTLSCryptoArgsForCall)]
	fake.checkTLSCryptoArgsForCall = append(fake.checkTLSCryptoArgsForCall, struct {
		arg1 v1.Object
		arg2 string
	}{arg1, arg2})
	stub := fake.CheckTLSCryptoStub
	fakeReturns := fake.checkTLSCryptoReturns
	fake.recordInvocation("CheckTLSCrypto", []interface{}{arg1, arg2})
	fake.checkTLSCryptoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CryptoValidator) CheckTLSCryptoCallCount() int {
	fake.checkTLSCryptoMutex.RLock()
	defer fake.checkTLSCryptoMutex.RUnlock()
	return len(fake.checkTLSCryptoArgsForCall)
}

func (fake *CryptoValidator) CheckTLSCryptoCalls(stub func(v1.Object, string) error) {
	fake.checkTLSCryptoMutex.Lock()
	defer fake.checkTLSCryptoMutex.Unlock()
	fake.CheckTLSCryptoStub = stub
}

func (fake *CryptoValidator) CheckTLSCryptoArgsForCall(i int) (v1.Object, string) {
	fake.checkTLSCryptoMutex.RLock()
	defer fake.checkTLSCryptoMutex.RUnlock()
	argsForCall := fake.checkTLSCryptoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CryptoValidator) CheckTLSCryptoReturns(result1 error) {
	fake.checkTLSCryptoMutex.Lock()
	defer fake.checkTLSCryptoMutex.Unlock()
	fake.CheckTLSCryptoStub = nil
	fake.checkTLSCryptoReturns = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) CheckTLSCryptoReturnsOnCall(i int, result1 error) {
	fake.checkTLSCryptoMutex.Lock()
	defer fake.checkTLSCryptoMutex.Unlock()
	fake.CheckTLSCryptoStub = nil
	if fake.checkTLSCryptoReturnsOnCall == nil {
		fake.checkTLSCryptoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkTLSCryptoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CryptoValidator) SetHSMEnabled(arg1 bool) {
	fake.setHSMEnabledMutex.Lock()
	fake.setHSMEnabledArgsForCall = append(fake.setHSMEnabledArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetHSMEnabledStub
	fake.recordInvocation("SetHSMEnabled", []interface{}{arg1})
	fake.setHSMEnabledMutex.Unlock()
	if stub != nil {
		fake.SetHSMEnabledStub(arg1)
	}
}

func (fake *CryptoValidator) SetHSMEnabledCallCount() int {
	fake.setHSMEnabledMutex.RLock()
	defer fake.setHSMEnabledMutex.RUnlock()
	return len(fake.setHSMEnabledArgsForCall)
}

func (fake *CryptoValidator) SetHSMEnabledCalls(stub func(bool)) {
	fake.setHSMEnabledMutex.Lock()
	defer fake.setHSMEnabledMutex.Unlock()
	fake.SetHSMEnabledStub = stub
}

func (fake *CryptoValidator) SetHSMEnabledArgsForCall(i int) bool {
	fake.setHSMEnabledMutex.RLock()
	defer fake.setHSMEnabledMutex.RUnlock()
	argsForCall := fake.setHSMEnabledArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CryptoValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkClientAuthCryptoMutex.RLock()
	defer fake.checkClientAuthCryptoMutex.RUnlock()
	fake.checkEcertCryptoMutex.RLock()
	defer fake.checkEcertCryptoMutex.RUnlock()
	fake.checkTLSCryptoMutex.RLock()
	defer fake.checkTLSCryptoMutex.RUnlock()
	fake.setHSMEnabledMutex.RLock()
	defer fake.setHSMEnabledMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CryptoValidator) recordInvocation(key string, args []interface{}) {
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

var _ common.CryptoValidator = new(CryptoValidator)
