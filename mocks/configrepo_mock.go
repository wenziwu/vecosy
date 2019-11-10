// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/configrepo/init.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	go_version "github.com/hashicorp/go-version"
	configrepo "github.com/vecosy/vecosy/v2/pkg/configrepo"
	reflect "reflect"
	time "time"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockRepo) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRepoMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRepo)(nil).Init))
}

// GetAppsVersions mocks base method
func (m *MockRepo) GetAppsVersions() map[string][]*go_version.Version {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppsVersions")
	ret0, _ := ret[0].(map[string][]*go_version.Version)
	return ret0
}

// GetAppsVersions indicates an expected call of GetAppsVersions
func (mr *MockRepoMockRecorder) GetAppsVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppsVersions", reflect.TypeOf((*MockRepo)(nil).GetAppsVersions))
}

// GetFile mocks base method
func (m *MockRepo) GetFile(targetApp, targetVersion, path string) (*configrepo.RepoFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", targetApp, targetVersion, path)
	ret0, _ := ret[0].(*configrepo.RepoFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockRepoMockRecorder) GetFile(targetApp, targetVersion, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockRepo)(nil).GetFile), targetApp, targetVersion, path)
}

// Fetch mocks base method
func (m *MockRepo) Fetch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch")
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch
func (mr *MockRepoMockRecorder) Fetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockRepo)(nil).Fetch))
}

// GetLastFetch mocks base method
func (m *MockRepo) GetLastFetch() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFetch")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// GetLastFetch indicates an expected call of GetLastFetch
func (mr *MockRepoMockRecorder) GetLastFetch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFetch", reflect.TypeOf((*MockRepo)(nil).GetLastFetch))
}

// StartFetchingEvery mocks base method
func (m *MockRepo) StartFetchingEvery(period time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFetchingEvery", period)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartFetchingEvery indicates an expected call of StartFetchingEvery
func (mr *MockRepoMockRecorder) StartFetchingEvery(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFetchingEvery", reflect.TypeOf((*MockRepo)(nil).StartFetchingEvery), period)
}

// StopFetching mocks base method
func (m *MockRepo) StopFetching() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopFetching")
}

// StopFetching indicates an expected call of StopFetching
func (mr *MockRepoMockRecorder) StopFetching() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopFetching", reflect.TypeOf((*MockRepo)(nil).StopFetching))
}

// AddOnChangeHandler mocks base method
func (m *MockRepo) AddOnChangeHandler(handler configrepo.OnChangeHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddOnChangeHandler", handler)
}

// AddOnChangeHandler indicates an expected call of AddOnChangeHandler
func (mr *MockRepoMockRecorder) AddOnChangeHandler(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOnChangeHandler", reflect.TypeOf((*MockRepo)(nil).AddOnChangeHandler), handler)
}
