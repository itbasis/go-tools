// Code generated by MockGen. DO NOT EDIT.
// Source: token.go
//
// Generated by this command:
//
//	mockgen -source=token.go -package=token -destination=token.mock.go
//

// Package token is a generated GoMock package.
package token

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/itbasis/go-tools/middleware/auth/model"
	gomock "go.uber.org/mock/gomock"
)

// MockJwtToken is a mock of JwtToken interface.
type MockJwtToken struct {
	ctrl     *gomock.Controller
	recorder *MockJwtTokenMockRecorder
	isgomock struct{}
}

// MockJwtTokenMockRecorder is the mock recorder for MockJwtToken.
type MockJwtTokenMockRecorder struct {
	mock *MockJwtToken
}

// NewMockJwtToken creates a new mock instance.
func NewMockJwtToken(ctrl *gomock.Controller) *MockJwtToken {
	mock := &MockJwtToken{ctrl: ctrl}
	mock.recorder = &MockJwtTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtToken) EXPECT() *MockJwtTokenMockRecorder {
	return m.recorder
}

// CreateAccessToken mocks base method.
func (m *MockJwtToken) CreateAccessToken(arg0 context.Context, arg1 model.SessionUser) (string, *time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockJwtTokenMockRecorder) CreateAccessToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockJwtToken)(nil).CreateAccessToken), arg0, arg1)
}

// CreateRefreshToken mocks base method.
func (m *MockJwtToken) CreateRefreshToken(arg0 context.Context, arg1 model.SessionUser) (string, *time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockJwtTokenMockRecorder) CreateRefreshToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockJwtToken)(nil).CreateRefreshToken), arg0, arg1)
}

// CreateTokenCustomDuration mocks base method.
func (m *MockJwtToken) CreateTokenCustomDuration(arg0 context.Context, arg1 model.SessionUser, arg2 time.Duration) (string, *time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTokenCustomDuration", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTokenCustomDuration indicates an expected call of CreateTokenCustomDuration.
func (mr *MockJwtTokenMockRecorder) CreateTokenCustomDuration(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTokenCustomDuration", reflect.TypeOf((*MockJwtToken)(nil).CreateTokenCustomDuration), arg0, arg1, arg2)
}

// Parse mocks base method.
func (m *MockJwtToken) Parse(ctx context.Context, tokenString string) (*model.SessionUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", ctx, tokenString)
	ret0, _ := ret[0].(*model.SessionUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockJwtTokenMockRecorder) Parse(ctx, tokenString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockJwtToken)(nil).Parse), ctx, tokenString)
}
