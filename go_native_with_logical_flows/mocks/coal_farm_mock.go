// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artiomgiza/test_me/go_native_with_logical_flows (interfaces: CoalFarm)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCoalFarm is a mock of CoalFarm interface
type MockCoalFarm struct {
	ctrl     *gomock.Controller
	recorder *MockCoalFarmMockRecorder
}

// MockCoalFarmMockRecorder is the mock recorder for MockCoalFarm
type MockCoalFarmMockRecorder struct {
	mock *MockCoalFarm
}

// NewMockCoalFarm creates a new mock instance
func NewMockCoalFarm(ctrl *gomock.Controller) *MockCoalFarm {
	mock := &MockCoalFarm{ctrl: ctrl}
	mock.recorder = &MockCoalFarmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoalFarm) EXPECT() *MockCoalFarmMockRecorder {
	return m.recorder
}

// GetCoalPrice mocks base method
func (m *MockCoalFarm) GetCoalPrice(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoalPrice", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoalPrice indicates an expected call of GetCoalPrice
func (mr *MockCoalFarmMockRecorder) GetCoalPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoalPrice", reflect.TypeOf((*MockCoalFarm)(nil).GetCoalPrice), arg0)
}
