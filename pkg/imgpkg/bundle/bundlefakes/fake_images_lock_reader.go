// Code generated by counterfeiter. DO NOT EDIT.
package bundlefakes

import (
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/k14s/imgpkg/pkg/imgpkg/bundle"
	"github.com/k14s/imgpkg/pkg/imgpkg/lockconfig"
)

type FakeImagesLockReader struct {
	ReadStub        func(v1.Image) (lockconfig.ImagesLock, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 v1.Image
	}
	readReturns struct {
		result1 lockconfig.ImagesLock
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 lockconfig.ImagesLock
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImagesLockReader) Read(arg1 v1.Image) (lockconfig.ImagesLock, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 v1.Image
	}{arg1})
	stub := fake.ReadStub
	fakeReturns := fake.readReturns
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImagesLockReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeImagesLockReader) ReadCalls(stub func(v1.Image) (lockconfig.ImagesLock, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeImagesLockReader) ReadArgsForCall(i int) v1.Image {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImagesLockReader) ReadReturns(result1 lockconfig.ImagesLock, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 lockconfig.ImagesLock
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesLockReader) ReadReturnsOnCall(i int, result1 lockconfig.ImagesLock, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 lockconfig.ImagesLock
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 lockconfig.ImagesLock
		result2 error
	}{result1, result2}
}

func (fake *FakeImagesLockReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImagesLockReader) recordInvocation(key string, args []interface{}) {
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

var _ bundle.ImagesLockReader = new(FakeImagesLockReader)
