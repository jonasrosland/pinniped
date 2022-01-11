// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coreos/go-oidc/v3/oidc (interfaces: KeySet)

// Package mockkeyset is a generated GoMock package.
package mockkeyset

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKeySet is a mock of KeySet interface.
type MockKeySet struct {
	ctrl     *gomock.Controller
	recorder *MockKeySetMockRecorder
}

// MockKeySetMockRecorder is the mock recorder for MockKeySet.
type MockKeySetMockRecorder struct {
	mock *MockKeySet
}

// NewMockKeySet creates a new mock instance.
func NewMockKeySet(ctrl *gomock.Controller) *MockKeySet {
	mock := &MockKeySet{ctrl: ctrl}
	mock.recorder = &MockKeySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeySet) EXPECT() *MockKeySetMockRecorder {
	return m.recorder
}

// VerifySignature mocks base method.
func (m *MockKeySet) VerifySignature(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySignature", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySignature indicates an expected call of VerifySignature.
func (mr *MockKeySetMockRecorder) VerifySignature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySignature", reflect.TypeOf((*MockKeySet)(nil).VerifySignature), arg0, arg1)
}
