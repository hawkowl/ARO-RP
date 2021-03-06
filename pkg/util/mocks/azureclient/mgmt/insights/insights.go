// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/insights (interfaces: MetricAlertsClient)

// Package mock_insights is a generated GoMock package.
package mock_insights

import (
	context "context"
	reflect "reflect"

	insights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2018-03-01/insights"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockMetricAlertsClient is a mock of MetricAlertsClient interface
type MockMetricAlertsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMetricAlertsClientMockRecorder
}

// MockMetricAlertsClientMockRecorder is the mock recorder for MockMetricAlertsClient
type MockMetricAlertsClientMockRecorder struct {
	mock *MockMetricAlertsClient
}

// NewMockMetricAlertsClient creates a new mock instance
func NewMockMetricAlertsClient(ctrl *gomock.Controller) *MockMetricAlertsClient {
	mock := &MockMetricAlertsClient{ctrl: ctrl}
	mock.recorder = &MockMetricAlertsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricAlertsClient) EXPECT() *MockMetricAlertsClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockMetricAlertsClient) Delete(arg0 context.Context, arg1, arg2 string) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockMetricAlertsClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMetricAlertsClient)(nil).Delete), arg0, arg1, arg2)
}

// ListByResourceGroup mocks base method
func (m *MockMetricAlertsClient) ListByResourceGroup(arg0 context.Context, arg1 string) (insights.MetricAlertResourceCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByResourceGroup", arg0, arg1)
	ret0, _ := ret[0].(insights.MetricAlertResourceCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByResourceGroup indicates an expected call of ListByResourceGroup
func (mr *MockMetricAlertsClientMockRecorder) ListByResourceGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByResourceGroup", reflect.TypeOf((*MockMetricAlertsClient)(nil).ListByResourceGroup), arg0, arg1)
}
