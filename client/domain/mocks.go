// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTimeConverter is a mock of TimeConverter interface.
type MockTimeConverter struct {
	ctrl     *gomock.Controller
	recorder *MockTimeConverterMockRecorder
}

// MockTimeConverterMockRecorder is the mock recorder for MockTimeConverter.
type MockTimeConverterMockRecorder struct {
	mock *MockTimeConverter
}

// NewMockTimeConverter creates a new mock instance.
func NewMockTimeConverter(ctrl *gomock.Controller) *MockTimeConverter {
	mock := &MockTimeConverter{ctrl: ctrl}
	mock.recorder = &MockTimeConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimeConverter) EXPECT() *MockTimeConverterMockRecorder {
	return m.recorder
}

// ConvertTime mocks base method.
func (m *MockTimeConverter) ConvertTime(hours, minutes string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertTime", hours, minutes)
	ret0, _ := ret[0].(string)
	return ret0
}

// ConvertTime indicates an expected call of ConvertTime.
func (mr *MockTimeConverterMockRecorder) ConvertTime(hours, minutes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTime", reflect.TypeOf((*MockTimeConverter)(nil).ConvertTime), hours, minutes)
}

// MockUserInterface is a mock of UserInterface interface.
type MockUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserInterfaceMockRecorder
}

// MockUserInterfaceMockRecorder is the mock recorder for MockUserInterface.
type MockUserInterfaceMockRecorder struct {
	mock *MockUserInterface
}

// NewMockUserInterface creates a new mock instance.
func NewMockUserInterface(ctrl *gomock.Controller) *MockUserInterface {
	mock := &MockUserInterface{ctrl: ctrl}
	mock.recorder = &MockUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInterface) EXPECT() *MockUserInterfaceMockRecorder {
	return m.recorder
}

// GetTime mocks base method.
func (m *MockUserInterface) GetTime() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTime indicates an expected call of GetTime.
func (mr *MockUserInterfaceMockRecorder) GetTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockUserInterface)(nil).GetTime))
}

// ShowTime mocks base method.
func (m *MockUserInterface) ShowTime(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowTime", arg0)
}

// ShowTime indicates an expected call of ShowTime.
func (mr *MockUserInterfaceMockRecorder) ShowTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowTime", reflect.TypeOf((*MockUserInterface)(nil).ShowTime), arg0)
}