// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/antrea/pkg/agent/flowexporter/connections (interfaces: ConnTrackPoller,ConnTrack)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	conntrack "github.com/ti-mo/conntrack"
	flowexporter "github.com/vmware-tanzu/antrea/pkg/agent/flowexporter"
	reflect "reflect"
)

// MockConnTrackPoller is a mock of ConnTrackPoller interface
type MockConnTrackPoller struct {
	ctrl     *gomock.Controller
	recorder *MockConnTrackPollerMockRecorder
}

// MockConnTrackPollerMockRecorder is the mock recorder for MockConnTrackPoller
type MockConnTrackPollerMockRecorder struct {
	mock *MockConnTrackPoller
}

// NewMockConnTrackPoller creates a new mock instance
func NewMockConnTrackPoller(ctrl *gomock.Controller) *MockConnTrackPoller {
	mock := &MockConnTrackPoller{ctrl: ctrl}
	mock.recorder = &MockConnTrackPollerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnTrackPoller) EXPECT() *MockConnTrackPollerMockRecorder {
	return m.recorder
}

// DumpFlows mocks base method
func (m *MockConnTrackPoller) DumpFlows(arg0 uint16) ([]*flowexporter.Connection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpFlows", arg0)
	ret0, _ := ret[0].([]*flowexporter.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpFlows indicates an expected call of DumpFlows
func (mr *MockConnTrackPollerMockRecorder) DumpFlows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpFlows", reflect.TypeOf((*MockConnTrackPoller)(nil).DumpFlows), arg0)
}

// MockConnTrack is a mock of ConnTrack interface
type MockConnTrack struct {
	ctrl     *gomock.Controller
	recorder *MockConnTrackMockRecorder
}

// MockConnTrackMockRecorder is the mock recorder for MockConnTrack
type MockConnTrackMockRecorder struct {
	mock *MockConnTrack
}

// NewMockConnTrack creates a new mock instance
func NewMockConnTrack(ctrl *gomock.Controller) *MockConnTrack {
	mock := &MockConnTrack{ctrl: ctrl}
	mock.recorder = &MockConnTrackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnTrack) EXPECT() *MockConnTrackMockRecorder {
	return m.recorder
}

// Dial mocks base method
func (m *MockConnTrack) Dial() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dial")
	ret0, _ := ret[0].(error)
	return ret0
}

// Dial indicates an expected call of Dial
func (mr *MockConnTrackMockRecorder) Dial() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dial", reflect.TypeOf((*MockConnTrack)(nil).Dial))
}

// DumpFilter mocks base method
func (m *MockConnTrack) DumpFilter(arg0 conntrack.Filter) ([]conntrack.Flow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpFilter", arg0)
	ret0, _ := ret[0].([]conntrack.Flow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpFilter indicates an expected call of DumpFilter
func (mr *MockConnTrackMockRecorder) DumpFilter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpFilter", reflect.TypeOf((*MockConnTrack)(nil).DumpFilter), arg0)
}
