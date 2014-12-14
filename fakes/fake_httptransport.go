// This file was generated by counterfeiter
package fakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/robdimsdale/wundergo"
)

type FakeHTTPTransport struct {
	NewRequestStub        func(method, urlStr string, body io.Reader) (*http.Request, error)
	newRequestMutex       sync.RWMutex
	newRequestArgsForCall []struct {
		method string
		urlStr string
		body   io.Reader
	}
	newRequestReturns struct {
		result1 *http.Request
		result2 error
	}
	DoRequestStub        func(req *http.Request) (resp *http.Response, err error)
	doRequestMutex       sync.RWMutex
	doRequestArgsForCall []struct {
		req *http.Request
	}
	doRequestReturns struct {
		result1 *http.Response
		result2 error
	}
}

func (fake *FakeHTTPTransport) NewRequest(method string, urlStr string, body io.Reader) (*http.Request, error) {
	fake.newRequestMutex.Lock()
	fake.newRequestArgsForCall = append(fake.newRequestArgsForCall, struct {
		method string
		urlStr string
		body   io.Reader
	}{method, urlStr, body})
	fake.newRequestMutex.Unlock()
	if fake.NewRequestStub != nil {
		return fake.NewRequestStub(method, urlStr, body)
	} else {
		return fake.newRequestReturns.result1, fake.newRequestReturns.result2
	}
}

func (fake *FakeHTTPTransport) NewRequestCallCount() int {
	fake.newRequestMutex.RLock()
	defer fake.newRequestMutex.RUnlock()
	return len(fake.newRequestArgsForCall)
}

func (fake *FakeHTTPTransport) NewRequestArgsForCall(i int) (string, string, io.Reader) {
	fake.newRequestMutex.RLock()
	defer fake.newRequestMutex.RUnlock()
	return fake.newRequestArgsForCall[i].method, fake.newRequestArgsForCall[i].urlStr, fake.newRequestArgsForCall[i].body
}

func (fake *FakeHTTPTransport) NewRequestReturns(result1 *http.Request, result2 error) {
	fake.NewRequestStub = nil
	fake.newRequestReturns = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPTransport) DoRequest(req *http.Request) (resp *http.Response, err error) {
	fake.doRequestMutex.Lock()
	fake.doRequestArgsForCall = append(fake.doRequestArgsForCall, struct {
		req *http.Request
	}{req})
	fake.doRequestMutex.Unlock()
	if fake.DoRequestStub != nil {
		return fake.DoRequestStub(req)
	} else {
		return fake.doRequestReturns.result1, fake.doRequestReturns.result2
	}
}

func (fake *FakeHTTPTransport) DoRequestCallCount() int {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	return len(fake.doRequestArgsForCall)
}

func (fake *FakeHTTPTransport) DoRequestArgsForCall(i int) *http.Request {
	fake.doRequestMutex.RLock()
	defer fake.doRequestMutex.RUnlock()
	return fake.doRequestArgsForCall[i].req
}

func (fake *FakeHTTPTransport) DoRequestReturns(result1 *http.Response, result2 error) {
	fake.DoRequestStub = nil
	fake.doRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

var _ wundergo.HTTPTransport = new(FakeHTTPTransport)
