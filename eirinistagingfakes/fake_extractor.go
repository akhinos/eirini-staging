// Code generated by counterfeiter. DO NOT EDIT.
package eirinistagingfakes

import (
	"sync"

	"code.cloudfoundry.org/eirinistaging"
)

type FakeExtractor struct {
	ExtractStub        func(string, string) error
	extractMutex       sync.RWMutex
	extractArgsForCall []struct {
		arg1 string
		arg2 string
	}
	extractReturns struct {
		result1 error
	}
	extractReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExtractor) Extract(arg1 string, arg2 string) error {
	fake.extractMutex.Lock()
	ret, specificReturn := fake.extractReturnsOnCall[len(fake.extractArgsForCall)]
	fake.extractArgsForCall = append(fake.extractArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Extract", []interface{}{arg1, arg2})
	fake.extractMutex.Unlock()
	if fake.ExtractStub != nil {
		return fake.ExtractStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.extractReturns
	return fakeReturns.result1
}

func (fake *FakeExtractor) ExtractCallCount() int {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return len(fake.extractArgsForCall)
}

func (fake *FakeExtractor) ExtractCalls(stub func(string, string) error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = stub
}

func (fake *FakeExtractor) ExtractArgsForCall(i int) (string, string) {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	argsForCall := fake.extractArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeExtractor) ExtractReturns(result1 error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = nil
	fake.extractReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractor) ExtractReturnsOnCall(i int, result1 error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = nil
	if fake.extractReturnsOnCall == nil {
		fake.extractReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.extractReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExtractor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExtractor) recordInvocation(key string, args []interface{}) {
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

var _ eirinistaging.Extractor = new(FakeExtractor)
