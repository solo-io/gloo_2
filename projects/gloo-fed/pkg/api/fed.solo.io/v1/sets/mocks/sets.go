// Code generated by MockGen. DO NOT EDIT.
// Source: ./sets.go

// Package mock_v1sets is a generated GoMock package.
package mock_v1sets

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1"
	v1sets "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1/sets"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockGlooInstanceSet is a mock of GlooInstanceSet interface
type MockGlooInstanceSet struct {
	ctrl     *gomock.Controller
	recorder *MockGlooInstanceSetMockRecorder
}

// MockGlooInstanceSetMockRecorder is the mock recorder for MockGlooInstanceSet
type MockGlooInstanceSetMockRecorder struct {
	mock *MockGlooInstanceSet
}

// NewMockGlooInstanceSet creates a new mock instance
func NewMockGlooInstanceSet(ctrl *gomock.Controller) *MockGlooInstanceSet {
	mock := &MockGlooInstanceSet{ctrl: ctrl}
	mock.recorder = &MockGlooInstanceSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGlooInstanceSet) EXPECT() *MockGlooInstanceSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockGlooInstanceSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockGlooInstanceSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockGlooInstanceSet)(nil).Keys))
}

// List mocks base method
func (m *MockGlooInstanceSet) List(filterResource ...func(*v1.GlooInstance) bool) []*v1.GlooInstance {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.GlooInstance)
	return ret0
}

// List indicates an expected call of List
func (mr *MockGlooInstanceSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGlooInstanceSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockGlooInstanceSet) Map() map[string]*v1.GlooInstance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.GlooInstance)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockGlooInstanceSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockGlooInstanceSet)(nil).Map))
}

// Insert mocks base method
func (m *MockGlooInstanceSet) Insert(glooInstance ...*v1.GlooInstance) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range glooInstance {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockGlooInstanceSetMockRecorder) Insert(glooInstance ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockGlooInstanceSet)(nil).Insert), glooInstance...)
}

// Equal mocks base method
func (m *MockGlooInstanceSet) Equal(glooInstanceSet v1sets.GlooInstanceSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", glooInstanceSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockGlooInstanceSetMockRecorder) Equal(glooInstanceSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockGlooInstanceSet)(nil).Equal), glooInstanceSet)
}

// Has mocks base method
func (m *MockGlooInstanceSet) Has(glooInstance ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", glooInstance)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockGlooInstanceSetMockRecorder) Has(glooInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockGlooInstanceSet)(nil).Has), glooInstance)
}

// Delete mocks base method
func (m *MockGlooInstanceSet) Delete(glooInstance ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", glooInstance)
}

// Delete indicates an expected call of Delete
func (mr *MockGlooInstanceSetMockRecorder) Delete(glooInstance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGlooInstanceSet)(nil).Delete), glooInstance)
}

// Union mocks base method
func (m *MockGlooInstanceSet) Union(set v1sets.GlooInstanceSet) v1sets.GlooInstanceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.GlooInstanceSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockGlooInstanceSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockGlooInstanceSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockGlooInstanceSet) Difference(set v1sets.GlooInstanceSet) v1sets.GlooInstanceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.GlooInstanceSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockGlooInstanceSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockGlooInstanceSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockGlooInstanceSet) Intersection(set v1sets.GlooInstanceSet) v1sets.GlooInstanceSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.GlooInstanceSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockGlooInstanceSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockGlooInstanceSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockGlooInstanceSet) Find(id ezkube.ResourceId) (*v1.GlooInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.GlooInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockGlooInstanceSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGlooInstanceSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockGlooInstanceSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockGlooInstanceSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockGlooInstanceSet)(nil).Length))
}

// MockFailoverSchemeSet is a mock of FailoverSchemeSet interface
type MockFailoverSchemeSet struct {
	ctrl     *gomock.Controller
	recorder *MockFailoverSchemeSetMockRecorder
}

// MockFailoverSchemeSetMockRecorder is the mock recorder for MockFailoverSchemeSet
type MockFailoverSchemeSetMockRecorder struct {
	mock *MockFailoverSchemeSet
}

// NewMockFailoverSchemeSet creates a new mock instance
func NewMockFailoverSchemeSet(ctrl *gomock.Controller) *MockFailoverSchemeSet {
	mock := &MockFailoverSchemeSet{ctrl: ctrl}
	mock.recorder = &MockFailoverSchemeSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFailoverSchemeSet) EXPECT() *MockFailoverSchemeSetMockRecorder {
	return m.recorder
}

// Keys mocks base method
func (m *MockFailoverSchemeSet) Keys() sets.String {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].(sets.String)
	return ret0
}

// Keys indicates an expected call of Keys
func (mr *MockFailoverSchemeSetMockRecorder) Keys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Keys))
}

// List mocks base method
func (m *MockFailoverSchemeSet) List(filterResource ...func(*v1.FailoverScheme) bool) []*v1.FailoverScheme {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range filterResource {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*v1.FailoverScheme)
	return ret0
}

// List indicates an expected call of List
func (mr *MockFailoverSchemeSetMockRecorder) List(filterResource ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFailoverSchemeSet)(nil).List), filterResource...)
}

// Map mocks base method
func (m *MockFailoverSchemeSet) Map() map[string]*v1.FailoverScheme {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]*v1.FailoverScheme)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockFailoverSchemeSetMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Map))
}

// Insert mocks base method
func (m *MockFailoverSchemeSet) Insert(failoverScheme ...*v1.FailoverScheme) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range failoverScheme {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Insert", varargs...)
}

// Insert indicates an expected call of Insert
func (mr *MockFailoverSchemeSetMockRecorder) Insert(failoverScheme ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Insert), failoverScheme...)
}

// Equal mocks base method
func (m *MockFailoverSchemeSet) Equal(failoverSchemeSet v1sets.FailoverSchemeSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", failoverSchemeSet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal
func (mr *MockFailoverSchemeSetMockRecorder) Equal(failoverSchemeSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Equal), failoverSchemeSet)
}

// Has mocks base method
func (m *MockFailoverSchemeSet) Has(failoverScheme ezkube.ResourceId) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", failoverScheme)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockFailoverSchemeSetMockRecorder) Has(failoverScheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Has), failoverScheme)
}

// Delete mocks base method
func (m *MockFailoverSchemeSet) Delete(failoverScheme ezkube.ResourceId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", failoverScheme)
}

// Delete indicates an expected call of Delete
func (mr *MockFailoverSchemeSetMockRecorder) Delete(failoverScheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Delete), failoverScheme)
}

// Union mocks base method
func (m *MockFailoverSchemeSet) Union(set v1sets.FailoverSchemeSet) v1sets.FailoverSchemeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Union", set)
	ret0, _ := ret[0].(v1sets.FailoverSchemeSet)
	return ret0
}

// Union indicates an expected call of Union
func (mr *MockFailoverSchemeSetMockRecorder) Union(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Union", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Union), set)
}

// Difference mocks base method
func (m *MockFailoverSchemeSet) Difference(set v1sets.FailoverSchemeSet) v1sets.FailoverSchemeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", set)
	ret0, _ := ret[0].(v1sets.FailoverSchemeSet)
	return ret0
}

// Difference indicates an expected call of Difference
func (mr *MockFailoverSchemeSetMockRecorder) Difference(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Difference), set)
}

// Intersection mocks base method
func (m *MockFailoverSchemeSet) Intersection(set v1sets.FailoverSchemeSet) v1sets.FailoverSchemeSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersection", set)
	ret0, _ := ret[0].(v1sets.FailoverSchemeSet)
	return ret0
}

// Intersection indicates an expected call of Intersection
func (mr *MockFailoverSchemeSetMockRecorder) Intersection(set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersection", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Intersection), set)
}

// Find mocks base method
func (m *MockFailoverSchemeSet) Find(id ezkube.ResourceId) (*v1.FailoverScheme, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", id)
	ret0, _ := ret[0].(*v1.FailoverScheme)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockFailoverSchemeSetMockRecorder) Find(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Find), id)
}

// Length mocks base method
func (m *MockFailoverSchemeSet) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockFailoverSchemeSetMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFailoverSchemeSet)(nil).Length))
}
