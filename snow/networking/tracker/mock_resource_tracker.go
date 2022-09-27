// Code generated by MockGen. DO NOT EDIT.
// Source: snow/networking/tracker/resource_tracker.go

// Package tracker is a generated GoMock package.
package tracker

import (
	reflect "reflect"
	time "time"

	ids "github.com/SavaLabs/savannahnode/ids"
	gomock "github.com/golang/mock/gomock"
)

// MockTracker is a mock of Tracker interface.
type MockTracker struct {
	ctrl     *gomock.Controller
	recorder *MockTrackerMockRecorder
}

// MockTrackerMockRecorder is the mock recorder for MockTracker.
type MockTrackerMockRecorder struct {
	mock *MockTracker
}

// NewMockTracker creates a new mock instance.
func NewMockTracker(ctrl *gomock.Controller) *MockTracker {
	mock := &MockTracker{ctrl: ctrl}
	mock.recorder = &MockTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTracker) EXPECT() *MockTrackerMockRecorder {
	return m.recorder
}

// TimeUntilUsage mocks base method.
func (m *MockTracker) TimeUntilUsage(nodeID ids.NodeID, now time.Time, value float64) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeUntilUsage", nodeID, now, value)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeUntilUsage indicates an expected call of TimeUntilUsage.
func (mr *MockTrackerMockRecorder) TimeUntilUsage(nodeID, now, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilUsage", reflect.TypeOf((*MockTracker)(nil).TimeUntilUsage), nodeID, now, value)
}

// TotalUsage mocks base method.
func (m *MockTracker) TotalUsage() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalUsage")
	ret0, _ := ret[0].(float64)
	return ret0
}

// TotalUsage indicates an expected call of TotalUsage.
func (mr *MockTrackerMockRecorder) TotalUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalUsage", reflect.TypeOf((*MockTracker)(nil).TotalUsage))
}

// Usage mocks base method.
func (m *MockTracker) Usage(nodeID ids.NodeID, now time.Time) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Usage", nodeID, now)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Usage indicates an expected call of Usage.
func (mr *MockTrackerMockRecorder) Usage(nodeID, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Usage", reflect.TypeOf((*MockTracker)(nil).Usage), nodeID, now)
}

// MockResourceTracker is a mock of ResourceTracker interface.
type MockResourceTracker struct {
	ctrl     *gomock.Controller
	recorder *MockResourceTrackerMockRecorder
}

// MockResourceTrackerMockRecorder is the mock recorder for MockResourceTracker.
type MockResourceTrackerMockRecorder struct {
	mock *MockResourceTracker
}

// NewMockResourceTracker creates a new mock instance.
func NewMockResourceTracker(ctrl *gomock.Controller) *MockResourceTracker {
	mock := &MockResourceTracker{ctrl: ctrl}
	mock.recorder = &MockResourceTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceTracker) EXPECT() *MockResourceTrackerMockRecorder {
	return m.recorder
}

// CPUTracker mocks base method.
func (m *MockResourceTracker) CPUTracker() Tracker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CPUTracker")
	ret0, _ := ret[0].(Tracker)
	return ret0
}

// CPUTracker indicates an expected call of CPUTracker.
func (mr *MockResourceTrackerMockRecorder) CPUTracker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CPUTracker", reflect.TypeOf((*MockResourceTracker)(nil).CPUTracker))
}

// DiskTracker mocks base method.
func (m *MockResourceTracker) DiskTracker() Tracker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskTracker")
	ret0, _ := ret[0].(Tracker)
	return ret0
}

// DiskTracker indicates an expected call of DiskTracker.
func (mr *MockResourceTrackerMockRecorder) DiskTracker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskTracker", reflect.TypeOf((*MockResourceTracker)(nil).DiskTracker))
}

// StartProcessing mocks base method.
func (m *MockResourceTracker) StartProcessing(arg0 ids.NodeID, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartProcessing", arg0, arg1)
}

// StartProcessing indicates an expected call of StartProcessing.
func (mr *MockResourceTrackerMockRecorder) StartProcessing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartProcessing", reflect.TypeOf((*MockResourceTracker)(nil).StartProcessing), arg0, arg1)
}

// StopProcessing mocks base method.
func (m *MockResourceTracker) StopProcessing(arg0 ids.NodeID, arg1 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopProcessing", arg0, arg1)
}

// StopProcessing indicates an expected call of StopProcessing.
func (mr *MockResourceTrackerMockRecorder) StopProcessing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopProcessing", reflect.TypeOf((*MockResourceTracker)(nil).StopProcessing), arg0, arg1)
}
