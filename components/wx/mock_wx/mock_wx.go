// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/csiabb/donation-service/components/wx (interfaces: IWXClient)

// Package mock_wx is a generated GoMock package.
package mock_wx

import (
	wx "github.com/csiabb/donation-service/components/wx"
	structs "github.com/csiabb/donation-service/structs"
	gomock "github.com/golang/mock/gomock"
	image "image"
	reflect "reflect"
)

// MockIWXClient is a mock of IWXClient interface
type MockIWXClient struct {
	ctrl     *gomock.Controller
	recorder *MockIWXClientMockRecorder
}

// MockIWXClientMockRecorder is the mock recorder for MockIWXClient
type MockIWXClientMockRecorder struct {
	mock *MockIWXClient
}

// NewMockIWXClient creates a new mock instance
func NewMockIWXClient(ctrl *gomock.Controller) *MockIWXClient {
	mock := &MockIWXClient{ctrl: ctrl}
	mock.recorder = &MockIWXClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIWXClient) EXPECT() *MockIWXClientMockRecorder {
	return m.recorder
}

// CheckFinger mocks base method
func (m *MockIWXClient) CheckFinger(arg0 structs.FingerRequest, arg1 string) (*structs.FingerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFinger", arg0, arg1)
	ret0, _ := ret[0].(*structs.FingerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFinger indicates an expected call of CheckFinger
func (mr *MockIWXClientMockRecorder) CheckFinger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFinger", reflect.TypeOf((*MockIWXClient)(nil).CheckFinger), arg0, arg1)
}

// DecryptPhoneNumber mocks base method
func (m *MockIWXClient) DecryptPhoneNumber(arg0, arg1, arg2 string) (wx.PhoneNumber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptPhoneNumber", arg0, arg1, arg2)
	ret0, _ := ret[0].(wx.PhoneNumber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptPhoneNumber indicates an expected call of DecryptPhoneNumber
func (mr *MockIWXClientMockRecorder) DecryptPhoneNumber(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptPhoneNumber", reflect.TypeOf((*MockIWXClient)(nil).DecryptPhoneNumber), arg0, arg1, arg2)
}

// DecryptUserInfo mocks base method
func (m *MockIWXClient) DecryptUserInfo(arg0, arg1, arg2, arg3, arg4 string) (wx.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptUserInfo", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(wx.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptUserInfo indicates an expected call of DecryptUserInfo
func (mr *MockIWXClientMockRecorder) DecryptUserInfo(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptUserInfo", reflect.TypeOf((*MockIWXClient)(nil).DecryptUserInfo), arg0, arg1, arg2, arg3, arg4)
}

// GetAccessToken mocks base method
func (m *MockIWXClient) GetAccessToken(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken
func (mr *MockIWXClientMockRecorder) GetAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockIWXClient)(nil).GetAccessToken), arg0, arg1)
}

// GetWXQrCode mocks base method
func (m *MockIWXClient) GetWXQrCode(arg0, arg1 string) (image.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWXQrCode", arg0, arg1)
	ret0, _ := ret[0].(image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWXQrCode indicates an expected call of GetWXQrCode
func (mr *MockIWXClientMockRecorder) GetWXQrCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWXQrCode", reflect.TypeOf((*MockIWXClient)(nil).GetWXQrCode), arg0, arg1)
}

// WXLogin mocks base method
func (m *MockIWXClient) WXLogin(arg0, arg1, arg2 string) (wx.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WXLogin", arg0, arg1, arg2)
	ret0, _ := ret[0].(wx.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WXLogin indicates an expected call of WXLogin
func (mr *MockIWXClientMockRecorder) WXLogin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WXLogin", reflect.TypeOf((*MockIWXClient)(nil).WXLogin), arg0, arg1, arg2)
}
