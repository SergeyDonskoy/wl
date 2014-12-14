// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/robdimsdale/wundergo"
)

type FakeLogger struct {
	LogLineStub        func(message string)
	logLineMutex       sync.RWMutex
	logLineArgsForCall []struct {
		message string
	}
}

func (fake *FakeLogger) LogLine(message string) {
	fake.logLineMutex.Lock()
	fake.logLineArgsForCall = append(fake.logLineArgsForCall, struct {
		message string
	}{message})
	fake.logLineMutex.Unlock()
	if fake.LogLineStub != nil {
		fake.LogLineStub(message)
	}
}

func (fake *FakeLogger) LogLineCallCount() int {
	fake.logLineMutex.RLock()
	defer fake.logLineMutex.RUnlock()
	return len(fake.logLineArgsForCall)
}

func (fake *FakeLogger) LogLineArgsForCall(i int) string {
	fake.logLineMutex.RLock()
	defer fake.logLineMutex.RUnlock()
	return fake.logLineArgsForCall[i].message
}

var _ wundergo.Logger = new(FakeLogger)