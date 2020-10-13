// Code generated by MockGen. DO NOT EDIT.
// Source: pkgs/coal-farm/coal_farm.go

// Package mock_coalfarm is a generated GoMock package.
package mock_coalfarm

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetCoal mocks base method
func (m *MockProvider) GetCoal(mangalsCounter int) (int, error) {
	ret := m.ctrl.Call(m, "GetCoal", mangalsCounter)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoal indicates an expected call of GetCoal
func (mr *MockProviderMockRecorder) GetCoal(mangalsCounter interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoal", reflect.TypeOf((*MockProvider)(nil).GetCoal), mangalsCounter)
}
