/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services (interfaces: ObjectStoreInterface)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scope "sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/scope"
)

// MockObjectStoreInterface is a mock of ObjectStoreInterface interface.
type MockObjectStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreInterfaceMockRecorder
}

// MockObjectStoreInterfaceMockRecorder is the mock recorder for MockObjectStoreInterface.
type MockObjectStoreInterfaceMockRecorder struct {
	mock *MockObjectStoreInterface
}

// NewMockObjectStoreInterface creates a new mock instance.
func NewMockObjectStoreInterface(ctrl *gomock.Controller) *MockObjectStoreInterface {
	mock := &MockObjectStoreInterface{ctrl: ctrl}
	mock.recorder = &MockObjectStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStoreInterface) EXPECT() *MockObjectStoreInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockObjectStoreInterface) Create(arg0 *scope.MachineScope, arg1 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockObjectStoreInterfaceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockObjectStoreInterface)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockObjectStoreInterface) Delete(arg0 *scope.MachineScope) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockObjectStoreInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockObjectStoreInterface)(nil).Delete), arg0)
}

// DeleteBucket mocks base method.
func (m *MockObjectStoreInterface) DeleteBucket() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockObjectStoreInterfaceMockRecorder) DeleteBucket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockObjectStoreInterface)(nil).DeleteBucket))
}

// ReconcileBucket mocks base method.
func (m *MockObjectStoreInterface) ReconcileBucket() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileBucket")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileBucket indicates an expected call of ReconcileBucket.
func (mr *MockObjectStoreInterfaceMockRecorder) ReconcileBucket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileBucket", reflect.TypeOf((*MockObjectStoreInterface)(nil).ReconcileBucket))
}
