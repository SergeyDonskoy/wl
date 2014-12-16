// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/robdimsdale/wundergo"
)

type FakeJSONHelper struct {
	MarshalStub        func(v interface{}) ([]byte, error)
	marshalMutex       sync.RWMutex
	marshalArgsForCall []struct {
		v interface{}
	}
	marshalReturns struct {
		result1 []byte
		result2 error
	}
	UnmarshalStub        func(data []byte, v interface{}) (interface{}, error)
	unmarshalMutex       sync.RWMutex
	unmarshalArgsForCall []struct {
		data []byte
		v    interface{}
	}
	unmarshalReturns struct {
		result1 interface{}
		result2 error
	}
}

func (fake *FakeJSONHelper) Marshal(v interface{}) ([]byte, error) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		v interface{}
	}{v})
	if fake.MarshalStub != nil {
		return fake.MarshalStub(v)
	} else {
		return fake.marshalReturns.result1, fake.marshalReturns.result2
	}
}

func (fake *FakeJSONHelper) MarshalCallCount() int {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return len(fake.marshalArgsForCall)
}

func (fake *FakeJSONHelper) MarshalArgsForCall(i int) interface{} {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return fake.marshalArgsForCall[i].v
}

func (fake *FakeJSONHelper) MarshalReturns(result1 []byte, result2 error) {
	fake.MarshalStub = nil
	fake.marshalReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeJSONHelper) Unmarshal(data []byte, v interface{}) (interface{}, error) {
	fake.unmarshalMutex.Lock()
	defer fake.unmarshalMutex.Unlock()
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		data []byte
		v    interface{}
	}{data, v})
	if fake.UnmarshalStub != nil {
		return fake.UnmarshalStub(data, v)
	} else {
		return fake.unmarshalReturns.result1, fake.unmarshalReturns.result2
	}
}

func (fake *FakeJSONHelper) UnmarshalCallCount() int {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return len(fake.unmarshalArgsForCall)
}

func (fake *FakeJSONHelper) UnmarshalArgsForCall(i int) ([]byte, interface{}) {
	fake.unmarshalMutex.RLock()
	defer fake.unmarshalMutex.RUnlock()
	return fake.unmarshalArgsForCall[i].data, fake.unmarshalArgsForCall[i].v
}

func (fake *FakeJSONHelper) UnmarshalReturns(result1 interface{}, result2 error) {
	fake.UnmarshalStub = nil
	fake.unmarshalReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

var _ wundergo.JSONHelper = new(FakeJSONHelper)
