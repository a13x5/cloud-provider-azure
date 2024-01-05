// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/azureclients/privateendpointclient/interface.go
//
// Generated by this command:
//
//	mockgen -copyright_file=/home/shafan/go/src/sigs.k8s.io/cloud-provider-azure/hack/boilerplate/boilerplate.generatego.txt -source=pkg/azureclients/privateendpointclient/interface.go -package=mockprivateendpointclient Interface
//

// Package mockprivateendpointclient is a generated GoMock package.
package mockprivateendpointclient

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2022-07-01/network"
	gomock "go.uber.org/mock/gomock"
	retry "sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, endpointName string, privateEndpoint network.PrivateEndpoint, etag string, waitForCompletion bool) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, endpointName, privateEndpoint, etag, waitForCompletion)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, endpointName, privateEndpoint, etag, waitForCompletion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, endpointName, privateEndpoint, etag, waitForCompletion)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, privateEndpointName, expand string) (network.PrivateEndpoint, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, privateEndpointName, expand)
	ret0, _ := ret[0].(network.PrivateEndpoint)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, privateEndpointName, expand any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, privateEndpointName, expand)
}
