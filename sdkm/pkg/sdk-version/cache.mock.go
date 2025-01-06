// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -source=cache.go -package=sdkversion -destination=cache.mock.go
//

// Package sdkversion is a generated GoMock package.
package sdkversion

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
	isgomock struct{}
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockCache) Load(ctx context.Context, versionType VersionType) []SDKVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", ctx, versionType)
	ret0, _ := ret[0].([]SDKVersion)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockCacheMockRecorder) Load(ctx, versionType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockCache)(nil).Load), ctx, versionType)
}

// Store mocks base method.
func (m *MockCache) Store(ctx context.Context, versionType VersionType, sdkVersions []SDKVersion) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Store", ctx, versionType, sdkVersions)
}

// Store indicates an expected call of Store.
func (mr *MockCacheMockRecorder) Store(ctx, versionType, sdkVersions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockCache)(nil).Store), ctx, versionType, sdkVersions)
}

// String mocks base method.
func (m *MockCache) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockCacheMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCache)(nil).String))
}

// Valid mocks base method.
func (m *MockCache) Valid(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Valid", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Valid indicates an expected call of Valid.
func (mr *MockCacheMockRecorder) Valid(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Valid", reflect.TypeOf((*MockCache)(nil).Valid), ctx)
}

// WithExternalStore mocks base method.
func (m *MockCache) WithExternalStore(cacheStorage CacheStorage) Cache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithExternalStore", cacheStorage)
	ret0, _ := ret[0].(Cache)
	return ret0
}

// WithExternalStore indicates an expected call of WithExternalStore.
func (mr *MockCacheMockRecorder) WithExternalStore(cacheStorage any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithExternalStore", reflect.TypeOf((*MockCache)(nil).WithExternalStore), cacheStorage)
}
