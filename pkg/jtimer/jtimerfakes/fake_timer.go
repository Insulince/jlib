// Code generated by counterfeiter. DO NOT EDIT.
package jtimerfakes

import (
	"sync"
	"time"

	"github.com/Insulince/jlib/pkg/jtimer"
)

type FakeTimer struct {
	StopStub        func(*time.Duration)
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		arg1 *time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTimer) Stop(arg1 *time.Duration) {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		arg1 *time.Duration
	}{arg1})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{arg1})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub(arg1)
	}
}

func (fake *FakeTimer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeTimer) StopCalls(stub func(*time.Duration)) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeTimer) StopArgsForCall(i int) *time.Duration {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	argsForCall := fake.stopArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTimer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTimer) recordInvocation(key string, args []interface{}) {
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

var _ jtimer.Timer = new(FakeTimer)
