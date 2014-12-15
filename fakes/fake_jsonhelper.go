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
	UnmarshalInlineStub        func(data []byte, v interface{}) error
	unmarshalInlineMutex       sync.RWMutex
	unmarshalInlineArgsForCall []struct {
		data []byte
		v    interface{}
	}
	unmarshalInlineReturns struct {
		result1 error
	}
}

func (fake *FakeJSONHelper) Marshal(v interface{}) ([]byte, error) {
	fake.marshalMutex.Lock()
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		v interface{}
	}{v})
	fake.marshalMutex.Unlock()
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
	fake.unmarshalArgsForCall = append(fake.unmarshalArgsForCall, struct {
		data []byte
		v    interface{}
	}{data, v})
	fake.unmarshalMutex.Unlock()
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

func (fake *FakeJSONHelper) UnmarshalInline(data []byte, v interface{}) error {
	fake.unmarshalInlineMutex.Lock()
	fake.unmarshalInlineArgsForCall = append(fake.unmarshalInlineArgsForCall, struct {
		data []byte
		v    interface{}
	}{data, v})
	fake.unmarshalInlineMutex.Unlock()
	if fake.UnmarshalInlineStub != nil {
		return fake.UnmarshalInlineStub(data, v)
	} else {
		return fake.unmarshalInlineReturns.result1
	}
}

func (fake *FakeJSONHelper) UnmarshalInlineCallCount() int {
	fake.unmarshalInlineMutex.RLock()
	defer fake.unmarshalInlineMutex.RUnlock()
	return len(fake.unmarshalInlineArgsForCall)
}

func (fake *FakeJSONHelper) UnmarshalInlineArgsForCall(i int) ([]byte, interface{}) {
	fake.unmarshalInlineMutex.RLock()
	defer fake.unmarshalInlineMutex.RUnlock()
	return fake.unmarshalInlineArgsForCall[i].data, fake.unmarshalInlineArgsForCall[i].v
}

func (fake *FakeJSONHelper) UnmarshalInlineReturns(result1 error) {
	fake.UnmarshalInlineStub = nil
	fake.unmarshalInlineReturns = struct {
		result1 error
	}{result1}
}

var _ wundergo.JSONHelper = new(FakeJSONHelper)