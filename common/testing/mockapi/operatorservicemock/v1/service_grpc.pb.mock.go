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
// Source: go.temporal.io/api/operatorservice/v1 (interfaces: OperatorServiceClient)
//
// Generated by this command:
//
//	mockgen -copyright_file ../../../LICENSE -package operatorservicemock -destination operatorservicemock/v1/service_grpc.pb.mock.go go.temporal.io/api/operatorservice/v1 OperatorServiceClient
//

// Package operatorservicemock is a generated GoMock package.
package operatorservicemock

import (
	context "context"
	reflect "reflect"

	operatorservice "go.temporal.io/api/operatorservice/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockOperatorServiceClient is a mock of OperatorServiceClient interface.
type MockOperatorServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorServiceClientMockRecorder
}

// MockOperatorServiceClientMockRecorder is the mock recorder for MockOperatorServiceClient.
type MockOperatorServiceClientMockRecorder struct {
	mock *MockOperatorServiceClient
}

// NewMockOperatorServiceClient creates a new mock instance.
func NewMockOperatorServiceClient(ctrl *gomock.Controller) *MockOperatorServiceClient {
	mock := &MockOperatorServiceClient{ctrl: ctrl}
	mock.recorder = &MockOperatorServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorServiceClient) EXPECT() *MockOperatorServiceClientMockRecorder {
	return m.recorder
}

// AddOrUpdateRemoteCluster mocks base method.
func (m *MockOperatorServiceClient) AddOrUpdateRemoteCluster(arg0 context.Context, arg1 *operatorservice.AddOrUpdateRemoteClusterRequest, arg2 ...grpc.CallOption) (*operatorservice.AddOrUpdateRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOrUpdateRemoteCluster", varargs...)
	ret0, _ := ret[0].(*operatorservice.AddOrUpdateRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateRemoteCluster indicates an expected call of AddOrUpdateRemoteCluster.
func (mr *MockOperatorServiceClientMockRecorder) AddOrUpdateRemoteCluster(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateRemoteCluster", reflect.TypeOf((*MockOperatorServiceClient)(nil).AddOrUpdateRemoteCluster), varargs...)
}

// AddSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) AddSearchAttributes(arg0 context.Context, arg1 *operatorservice.AddSearchAttributesRequest, arg2 ...grpc.CallOption) (*operatorservice.AddSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.AddSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSearchAttributes indicates an expected call of AddSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) AddSearchAttributes(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).AddSearchAttributes), varargs...)
}

// CreateNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) CreateNexusEndpoint(arg0 context.Context, arg1 *operatorservice.CreateNexusEndpointRequest, arg2 ...grpc.CallOption) (*operatorservice.CreateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.CreateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNexusEndpoint indicates an expected call of CreateNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) CreateNexusEndpoint(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).CreateNexusEndpoint), varargs...)
}

// DeleteNamespace mocks base method.
func (m *MockOperatorServiceClient) DeleteNamespace(arg0 context.Context, arg1 *operatorservice.DeleteNamespaceRequest, arg2 ...grpc.CallOption) (*operatorservice.DeleteNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNamespace", varargs...)
	ret0, _ := ret[0].(*operatorservice.DeleteNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockOperatorServiceClientMockRecorder) DeleteNamespace(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockOperatorServiceClient)(nil).DeleteNamespace), varargs...)
}

// DeleteNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) DeleteNexusEndpoint(arg0 context.Context, arg1 *operatorservice.DeleteNexusEndpointRequest, arg2 ...grpc.CallOption) (*operatorservice.DeleteNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.DeleteNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNexusEndpoint indicates an expected call of DeleteNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) DeleteNexusEndpoint(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).DeleteNexusEndpoint), varargs...)
}

// GetNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) GetNexusEndpoint(arg0 context.Context, arg1 *operatorservice.GetNexusEndpointRequest, arg2 ...grpc.CallOption) (*operatorservice.GetNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.GetNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNexusEndpoint indicates an expected call of GetNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) GetNexusEndpoint(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).GetNexusEndpoint), varargs...)
}

// ListClusters mocks base method.
func (m *MockOperatorServiceClient) ListClusters(arg0 context.Context, arg1 *operatorservice.ListClustersRequest, arg2 ...grpc.CallOption) (*operatorservice.ListClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockOperatorServiceClientMockRecorder) ListClusters(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListClusters), varargs...)
}

// ListNexusEndpoints mocks base method.
func (m *MockOperatorServiceClient) ListNexusEndpoints(arg0 context.Context, arg1 *operatorservice.ListNexusEndpointsRequest, arg2 ...grpc.CallOption) (*operatorservice.ListNexusEndpointsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNexusEndpoints", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListNexusEndpointsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNexusEndpoints indicates an expected call of ListNexusEndpoints.
func (mr *MockOperatorServiceClientMockRecorder) ListNexusEndpoints(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNexusEndpoints", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListNexusEndpoints), varargs...)
}

// ListSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) ListSearchAttributes(arg0 context.Context, arg1 *operatorservice.ListSearchAttributesRequest, arg2 ...grpc.CallOption) (*operatorservice.ListSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearchAttributes indicates an expected call of ListSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) ListSearchAttributes(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListSearchAttributes), varargs...)
}

// RemoveRemoteCluster mocks base method.
func (m *MockOperatorServiceClient) RemoveRemoteCluster(arg0 context.Context, arg1 *operatorservice.RemoveRemoteClusterRequest, arg2 ...grpc.CallOption) (*operatorservice.RemoveRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveRemoteCluster", varargs...)
	ret0, _ := ret[0].(*operatorservice.RemoveRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRemoteCluster indicates an expected call of RemoveRemoteCluster.
func (mr *MockOperatorServiceClientMockRecorder) RemoveRemoteCluster(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteCluster", reflect.TypeOf((*MockOperatorServiceClient)(nil).RemoveRemoteCluster), varargs...)
}

// RemoveSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) RemoveSearchAttributes(arg0 context.Context, arg1 *operatorservice.RemoveSearchAttributesRequest, arg2 ...grpc.CallOption) (*operatorservice.RemoveSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.RemoveSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSearchAttributes indicates an expected call of RemoveSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) RemoveSearchAttributes(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).RemoveSearchAttributes), varargs...)
}

// UpdateNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) UpdateNexusEndpoint(arg0 context.Context, arg1 *operatorservice.UpdateNexusEndpointRequest, arg2 ...grpc.CallOption) (*operatorservice.UpdateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.UpdateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNexusEndpoint indicates an expected call of UpdateNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) UpdateNexusEndpoint(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).UpdateNexusEndpoint), varargs...)
}
