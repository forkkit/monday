// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/build/builder.go

// Package build is a generated GoMock package.
package build

import (
	reflect "reflect"

	config "github.com/eko/monday/pkg/config"
	gomock "github.com/golang/mock/gomock"
)

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// BuildAll mocks base method.
func (m *MockBuilder) BuildAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BuildAll")
}

// BuildAll indicates an expected call of BuildAll.
func (mr *MockBuilderMockRecorder) BuildAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAll", reflect.TypeOf((*MockBuilder)(nil).BuildAll))
}

// Build mocks base method.
func (m *MockBuilder) Build(application *config.Application) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Build", application)
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build(application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), application)
}
