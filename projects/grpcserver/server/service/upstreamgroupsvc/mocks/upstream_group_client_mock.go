// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/api/v1 (interfaces: UpstreamGroupClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockUpstreamGroupClient is a mock of UpstreamGroupClient interface
type MockUpstreamGroupClient struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamGroupClientMockRecorder
}

// MockUpstreamGroupClientMockRecorder is the mock recorder for MockUpstreamGroupClient
type MockUpstreamGroupClientMockRecorder struct {
	mock *MockUpstreamGroupClient
}

// NewMockUpstreamGroupClient creates a new mock instance
func NewMockUpstreamGroupClient(ctrl *gomock.Controller) *MockUpstreamGroupClient {
	mock := &MockUpstreamGroupClient{ctrl: ctrl}
	mock.recorder = &MockUpstreamGroupClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamGroupClient) EXPECT() *MockUpstreamGroupClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockUpstreamGroupClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockUpstreamGroupClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockUpstreamGroupClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockUpstreamGroupClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUpstreamGroupClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUpstreamGroupClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockUpstreamGroupClient) List(arg0 string, arg1 clients.ListOpts) (v1.UpstreamGroupList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.UpstreamGroupList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUpstreamGroupClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUpstreamGroupClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockUpstreamGroupClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.UpstreamGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.UpstreamGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockUpstreamGroupClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockUpstreamGroupClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockUpstreamGroupClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUpstreamGroupClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUpstreamGroupClient)(nil).Register))
}

// Watch mocks base method
func (m *MockUpstreamGroupClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.UpstreamGroupList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.UpstreamGroupList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockUpstreamGroupClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockUpstreamGroupClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockUpstreamGroupClient) Write(arg0 *v1.UpstreamGroup, arg1 clients.WriteOpts) (*v1.UpstreamGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.UpstreamGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockUpstreamGroupClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockUpstreamGroupClient)(nil).Write), arg0, arg1)
}
