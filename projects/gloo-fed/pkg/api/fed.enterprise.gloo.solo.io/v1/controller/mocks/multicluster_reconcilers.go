// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.enterprise.gloo.solo.io/v1"
	controller "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.enterprise.gloo.solo.io/v1/controller"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterFederatedAuthConfigReconciler is a mock of MulticlusterFederatedAuthConfigReconciler interface
type MockMulticlusterFederatedAuthConfigReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedAuthConfigReconcilerMockRecorder
}

// MockMulticlusterFederatedAuthConfigReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedAuthConfigReconciler
type MockMulticlusterFederatedAuthConfigReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedAuthConfigReconciler
}

// NewMockMulticlusterFederatedAuthConfigReconciler creates a new mock instance
func NewMockMulticlusterFederatedAuthConfigReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedAuthConfigReconciler {
	mock := &MockMulticlusterFederatedAuthConfigReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedAuthConfigReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFederatedAuthConfigReconciler) EXPECT() *MockMulticlusterFederatedAuthConfigReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedAuthConfig mocks base method
func (m *MockMulticlusterFederatedAuthConfigReconciler) ReconcileFederatedAuthConfig(clusterName string, obj *v1.FederatedAuthConfig) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedAuthConfig", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFederatedAuthConfig indicates an expected call of ReconcileFederatedAuthConfig
func (mr *MockMulticlusterFederatedAuthConfigReconcilerMockRecorder) ReconcileFederatedAuthConfig(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedAuthConfig", reflect.TypeOf((*MockMulticlusterFederatedAuthConfigReconciler)(nil).ReconcileFederatedAuthConfig), clusterName, obj)
}

// MockMulticlusterFederatedAuthConfigDeletionReconciler is a mock of MulticlusterFederatedAuthConfigDeletionReconciler interface
type MockMulticlusterFederatedAuthConfigDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder
}

// MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterFederatedAuthConfigDeletionReconciler
type MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterFederatedAuthConfigDeletionReconciler
}

// NewMockMulticlusterFederatedAuthConfigDeletionReconciler creates a new mock instance
func NewMockMulticlusterFederatedAuthConfigDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterFederatedAuthConfigDeletionReconciler {
	mock := &MockMulticlusterFederatedAuthConfigDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFederatedAuthConfigDeletionReconciler) EXPECT() *MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileFederatedAuthConfigDeletion mocks base method
func (m *MockMulticlusterFederatedAuthConfigDeletionReconciler) ReconcileFederatedAuthConfigDeletion(clusterName string, req reconcile.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFederatedAuthConfigDeletion", clusterName, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileFederatedAuthConfigDeletion indicates an expected call of ReconcileFederatedAuthConfigDeletion
func (mr *MockMulticlusterFederatedAuthConfigDeletionReconcilerMockRecorder) ReconcileFederatedAuthConfigDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFederatedAuthConfigDeletion", reflect.TypeOf((*MockMulticlusterFederatedAuthConfigDeletionReconciler)(nil).ReconcileFederatedAuthConfigDeletion), clusterName, req)
}

// MockMulticlusterFederatedAuthConfigReconcileLoop is a mock of MulticlusterFederatedAuthConfigReconcileLoop interface
type MockMulticlusterFederatedAuthConfigReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder
}

// MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder is the mock recorder for MockMulticlusterFederatedAuthConfigReconcileLoop
type MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder struct {
	mock *MockMulticlusterFederatedAuthConfigReconcileLoop
}

// NewMockMulticlusterFederatedAuthConfigReconcileLoop creates a new mock instance
func NewMockMulticlusterFederatedAuthConfigReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterFederatedAuthConfigReconcileLoop {
	mock := &MockMulticlusterFederatedAuthConfigReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMulticlusterFederatedAuthConfigReconcileLoop) EXPECT() *MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterFederatedAuthConfigReconciler mocks base method
func (m *MockMulticlusterFederatedAuthConfigReconcileLoop) AddMulticlusterFederatedAuthConfigReconciler(ctx context.Context, rec controller.MulticlusterFederatedAuthConfigReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterFederatedAuthConfigReconciler", varargs...)
}

// AddMulticlusterFederatedAuthConfigReconciler indicates an expected call of AddMulticlusterFederatedAuthConfigReconciler
func (mr *MockMulticlusterFederatedAuthConfigReconcileLoopMockRecorder) AddMulticlusterFederatedAuthConfigReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterFederatedAuthConfigReconciler", reflect.TypeOf((*MockMulticlusterFederatedAuthConfigReconcileLoop)(nil).AddMulticlusterFederatedAuthConfigReconciler), varargs...)
}
