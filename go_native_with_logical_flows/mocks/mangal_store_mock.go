// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artiomgiza/test_me/go_native_with_logical_flows (interfaces: MangalStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMangalStore is a mock of MangalStore interface
type MockMangalStore struct {
	ctrl     *gomock.Controller
	recorder *MockMangalStoreMockRecorder
}

// MockMangalStoreMockRecorder is the mock recorder for MockMangalStore
type MockMangalStoreMockRecorder struct {
	mock *MockMangalStore
}

// NewMockMangalStore creates a new mock instance
func NewMockMangalStore(ctrl *gomock.Controller) *MockMangalStore {
	mock := &MockMangalStore{ctrl: ctrl}
	mock.recorder = &MockMangalStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMangalStore) EXPECT() *MockMangalStoreMockRecorder {
	return m.recorder
}

// GetMangalPrice mocks base method
func (m *MockMangalStore) GetMangalPrice() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMangalPrice")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMangalPrice indicates an expected call of GetMangalPrice
func (mr *MockMangalStoreMockRecorder) GetMangalPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMangalPrice", reflect.TypeOf((*MockMangalStore)(nil).GetMangalPrice))
}