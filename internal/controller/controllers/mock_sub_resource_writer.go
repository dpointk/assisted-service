// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/controller-runtime/pkg/client (interfaces: SubResourceWriter)

// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSubResourceWriter is a mock of SubResourceWriter interface.
type MockSubResourceWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSubResourceWriterMockRecorder
}

// MockSubResourceWriterMockRecorder is the mock recorder for MockSubResourceWriter.
type MockSubResourceWriterMockRecorder struct {
	mock *MockSubResourceWriter
}

// NewMockSubResourceWriter creates a new mock instance.
func NewMockSubResourceWriter(ctrl *gomock.Controller) *MockSubResourceWriter {
	mock := &MockSubResourceWriter{ctrl: ctrl}
	mock.recorder = &MockSubResourceWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubResourceWriter) EXPECT() *MockSubResourceWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSubResourceWriter) Create(arg0 context.Context, arg1, arg2 client.Object, arg3 ...client.SubResourceCreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSubResourceWriterMockRecorder) Create(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubResourceWriter)(nil).Create), varargs...)
}

// Patch mocks base method.
func (m *MockSubResourceWriter) Patch(arg0 context.Context, arg1 client.Object, arg2 client.Patch, arg3 ...client.SubResourcePatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockSubResourceWriterMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockSubResourceWriter)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockSubResourceWriter) Update(arg0 context.Context, arg1 client.Object, arg2 ...client.SubResourceUpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSubResourceWriterMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubResourceWriter)(nil).Update), varargs...)
}
