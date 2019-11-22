// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artiomgiza/test_me/go_native_with_logical_flows (interfaces: BeefFarm)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBeefFarm is a mock of BeefFarm interface
type MockBeefFarm struct {
	ctrl     *gomock.Controller
	recorder *MockBeefFarmMockRecorder
}

// MockBeefFarmMockRecorder is the mock recorder for MockBeefFarm
type MockBeefFarmMockRecorder struct {
	mock *MockBeefFarm
}

// NewMockBeefFarm creates a new mock instance
func NewMockBeefFarm(ctrl *gomock.Controller) *MockBeefFarm {
	mock := &MockBeefFarm{ctrl: ctrl}
	mock.recorder = &MockBeefFarmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeefFarm) EXPECT() *MockBeefFarmMockRecorder {
	return m.recorder
}

// GetEntrecotePrice mocks base method
func (m *MockBeefFarm) GetEntrecotePrice(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntrecotePrice", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntrecotePrice indicates an expected call of GetEntrecotePrice
func (mr *MockBeefFarmMockRecorder) GetEntrecotePrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntrecotePrice", reflect.TypeOf((*MockBeefFarm)(nil).GetEntrecotePrice), arg0)
}