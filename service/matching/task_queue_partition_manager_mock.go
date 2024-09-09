// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: task_queue_partition_manager_interface.go

// Package matching is a generated GoMock package.
package matching

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	taskqueue "go.temporal.io/api/taskqueue/v1"
	matchingservice "go.temporal.io/server/api/matchingservice/v1"
	namespace "go.temporal.io/server/common/namespace"
	tqid "go.temporal.io/server/common/tqid"
)

// MocktaskQueuePartitionManager is a mock of taskQueuePartitionManager interface.
type MocktaskQueuePartitionManager struct {
	ctrl     *gomock.Controller
	recorder *MocktaskQueuePartitionManagerMockRecorder
}

// MocktaskQueuePartitionManagerMockRecorder is the mock recorder for MocktaskQueuePartitionManager.
type MocktaskQueuePartitionManagerMockRecorder struct {
	mock *MocktaskQueuePartitionManager
}

// NewMocktaskQueuePartitionManager creates a new mock instance.
func NewMocktaskQueuePartitionManager(ctrl *gomock.Controller) *MocktaskQueuePartitionManager {
	mock := &MocktaskQueuePartitionManager{ctrl: ctrl}
	mock.recorder = &MocktaskQueuePartitionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktaskQueuePartitionManager) EXPECT() *MocktaskQueuePartitionManagerMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MocktaskQueuePartitionManager) AddTask(ctx context.Context, params addTaskParams) (string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", ctx, params)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddTask indicates an expected call of AddTask.
func (mr *MocktaskQueuePartitionManagerMockRecorder) AddTask(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).AddTask), ctx, params)
}

// Describe mocks base method.
func (m *MocktaskQueuePartitionManager) Describe(ctx context.Context, buildIds map[string]bool, includeAllActive, reportStats, reportPollers, internalTaskQueueStatus bool) (*matchingservice.DescribeTaskQueuePartitionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe", ctx, buildIds, includeAllActive, reportStats, reportPollers, internalTaskQueueStatus)
	ret0, _ := ret[0].(*matchingservice.DescribeTaskQueuePartitionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Describe indicates an expected call of Describe.
func (mr *MocktaskQueuePartitionManagerMockRecorder) Describe(ctx, buildIds, includeAllActive, reportStats, reportPollers, internalTaskQueueStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).Describe), ctx, buildIds, includeAllActive, reportStats, reportPollers, internalTaskQueueStatus)
}

// DispatchNexusTask mocks base method.
func (m *MocktaskQueuePartitionManager) DispatchNexusTask(ctx context.Context, taskId string, request *matchingservice.DispatchNexusTaskRequest) (*matchingservice.DispatchNexusTaskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchNexusTask", ctx, taskId, request)
	ret0, _ := ret[0].(*matchingservice.DispatchNexusTaskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispatchNexusTask indicates an expected call of DispatchNexusTask.
func (mr *MocktaskQueuePartitionManagerMockRecorder) DispatchNexusTask(ctx, taskId, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchNexusTask", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).DispatchNexusTask), ctx, taskId, request)
}

// DispatchQueryTask mocks base method.
func (m *MocktaskQueuePartitionManager) DispatchQueryTask(ctx context.Context, taskId string, request *matchingservice.QueryWorkflowRequest) (*matchingservice.QueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchQueryTask", ctx, taskId, request)
	ret0, _ := ret[0].(*matchingservice.QueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DispatchQueryTask indicates an expected call of DispatchQueryTask.
func (mr *MocktaskQueuePartitionManagerMockRecorder) DispatchQueryTask(ctx, taskId, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchQueryTask", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).DispatchQueryTask), ctx, taskId, request)
}

// GetAllPollerInfo mocks base method.
func (m *MocktaskQueuePartitionManager) GetAllPollerInfo() []*taskqueue.PollerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPollerInfo")
	ret0, _ := ret[0].([]*taskqueue.PollerInfo)
	return ret0
}

// GetAllPollerInfo indicates an expected call of GetAllPollerInfo.
func (mr *MocktaskQueuePartitionManagerMockRecorder) GetAllPollerInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPollerInfo", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).GetAllPollerInfo))
}

// GetUserDataManager mocks base method.
func (m *MocktaskQueuePartitionManager) GetUserDataManager() userDataManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDataManager")
	ret0, _ := ret[0].(userDataManager)
	return ret0
}

// GetUserDataManager indicates an expected call of GetUserDataManager.
func (mr *MocktaskQueuePartitionManagerMockRecorder) GetUserDataManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDataManager", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).GetUserDataManager))
}

// HasAnyPollerAfter mocks base method.
func (m *MocktaskQueuePartitionManager) HasAnyPollerAfter(accessTime time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasAnyPollerAfter", accessTime)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasAnyPollerAfter indicates an expected call of HasAnyPollerAfter.
func (mr *MocktaskQueuePartitionManagerMockRecorder) HasAnyPollerAfter(accessTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasAnyPollerAfter", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).HasAnyPollerAfter), accessTime)
}

// HasPollerAfter mocks base method.
func (m *MocktaskQueuePartitionManager) HasPollerAfter(buildId string, accessTime time.Time) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPollerAfter", buildId, accessTime)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPollerAfter indicates an expected call of HasPollerAfter.
func (mr *MocktaskQueuePartitionManagerMockRecorder) HasPollerAfter(buildId, accessTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPollerAfter", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).HasPollerAfter), buildId, accessTime)
}

// LegacyDescribeTaskQueue mocks base method.
func (m *MocktaskQueuePartitionManager) LegacyDescribeTaskQueue(includeTaskQueueStatus bool) *matchingservice.DescribeTaskQueueResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyDescribeTaskQueue", includeTaskQueueStatus)
	ret0, _ := ret[0].(*matchingservice.DescribeTaskQueueResponse)
	return ret0
}

// LegacyDescribeTaskQueue indicates an expected call of LegacyDescribeTaskQueue.
func (mr *MocktaskQueuePartitionManagerMockRecorder) LegacyDescribeTaskQueue(includeTaskQueueStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyDescribeTaskQueue", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).LegacyDescribeTaskQueue), includeTaskQueueStatus)
}

// LongPollExpirationInterval mocks base method.
func (m *MocktaskQueuePartitionManager) LongPollExpirationInterval() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LongPollExpirationInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// LongPollExpirationInterval indicates an expected call of LongPollExpirationInterval.
func (mr *MocktaskQueuePartitionManagerMockRecorder) LongPollExpirationInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LongPollExpirationInterval", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).LongPollExpirationInterval))
}

// MarkAlive mocks base method.
func (m *MocktaskQueuePartitionManager) MarkAlive() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkAlive")
}

// MarkAlive indicates an expected call of MarkAlive.
func (mr *MocktaskQueuePartitionManagerMockRecorder) MarkAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlive", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).MarkAlive))
}

// Namespace mocks base method.
func (m *MocktaskQueuePartitionManager) Namespace() *namespace.Namespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(*namespace.Namespace)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MocktaskQueuePartitionManagerMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).Namespace))
}

// Partition mocks base method.
func (m *MocktaskQueuePartitionManager) Partition() tqid.Partition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partition")
	ret0, _ := ret[0].(tqid.Partition)
	return ret0
}

// Partition indicates an expected call of Partition.
func (mr *MocktaskQueuePartitionManagerMockRecorder) Partition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partition", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).Partition))
}

// PollTask mocks base method.
func (m *MocktaskQueuePartitionManager) PollTask(ctx context.Context, pollMetadata *pollMetadata) (*internalTask, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollTask", ctx, pollMetadata)
	ret0, _ := ret[0].(*internalTask)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PollTask indicates an expected call of PollTask.
func (mr *MocktaskQueuePartitionManagerMockRecorder) PollTask(ctx, pollMetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollTask", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).PollTask), ctx, pollMetadata)
}

// ProcessSpooledTask mocks base method.
func (m *MocktaskQueuePartitionManager) ProcessSpooledTask(ctx context.Context, task *internalTask, assignedBuildId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessSpooledTask", ctx, task, assignedBuildId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSpooledTask indicates an expected call of ProcessSpooledTask.
func (mr *MocktaskQueuePartitionManagerMockRecorder) ProcessSpooledTask(ctx, task, assignedBuildId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSpooledTask", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).ProcessSpooledTask), ctx, task, assignedBuildId)
}

// Start mocks base method.
func (m *MocktaskQueuePartitionManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MocktaskQueuePartitionManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).Start))
}

// Stop mocks base method.
func (m *MocktaskQueuePartitionManager) Stop(arg0 unloadCause) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", arg0)
}

// Stop indicates an expected call of Stop.
func (mr *MocktaskQueuePartitionManagerMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).Stop), arg0)
}

// String mocks base method.
func (m *MocktaskQueuePartitionManager) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MocktaskQueuePartitionManagerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).String))
}

// WaitUntilInitialized mocks base method.
func (m *MocktaskQueuePartitionManager) WaitUntilInitialized(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilInitialized", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilInitialized indicates an expected call of WaitUntilInitialized.
func (mr *MocktaskQueuePartitionManagerMockRecorder) WaitUntilInitialized(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilInitialized", reflect.TypeOf((*MocktaskQueuePartitionManager)(nil).WaitUntilInitialized), arg0)
}
