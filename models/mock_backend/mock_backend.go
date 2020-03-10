// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/csiabb/donation-service/models (interfaces: IDBBackend)

// Package mock_backend is a generated GoMock package.
package mock_backend

import (
	models "github.com/csiabb/donation-service/models"
	structs "github.com/csiabb/donation-service/structs"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	reflect "reflect"
)

// MockIDBBackend is a mock of IDBBackend interface
type MockIDBBackend struct {
	ctrl     *gomock.Controller
	recorder *MockIDBBackendMockRecorder
}

// MockIDBBackendMockRecorder is the mock recorder for MockIDBBackend
type MockIDBBackendMockRecorder struct {
	mock *MockIDBBackend
}

// NewMockIDBBackend creates a new mock instance
func NewMockIDBBackend(ctrl *gomock.Controller) *MockIDBBackend {
	mock := &MockIDBBackend{ctrl: ctrl}
	mock.recorder = &MockIDBBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDBBackend) EXPECT() *MockIDBBackendMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockIDBBackend) CreateAccount(arg0 *models.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockIDBBackendMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIDBBackend)(nil).CreateAccount), arg0)
}

// CreateAddresses mocks base method
func (m *MockIDBBackend) CreateAddresses(arg0 *gorm.DB, arg1 []*models.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddresses", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAddresses indicates an expected call of CreateAddresses
func (mr *MockIDBBackendMockRecorder) CreateAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddresses", reflect.TypeOf((*MockIDBBackend)(nil).CreateAddresses), arg0, arg1)
}

// CreateFunds mocks base method
func (m *MockIDBBackend) CreateFunds(arg0 *gorm.DB, arg1 *models.PubFunds) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFunds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFunds indicates an expected call of CreateFunds
func (mr *MockIDBBackendMockRecorder) CreateFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFunds", reflect.TypeOf((*MockIDBBackend)(nil).CreateFunds), arg0, arg1)
}

// CreateImages mocks base method
func (m *MockIDBBackend) CreateImages(arg0 *gorm.DB, arg1 []*models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateImages indicates an expected call of CreateImages
func (mr *MockIDBBackendMockRecorder) CreateImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImages", reflect.TypeOf((*MockIDBBackend)(nil).CreateImages), arg0, arg1)
}

// CreateOrganization mocks base method
func (m *MockIDBBackend) CreateOrganization(arg0 *models.DonationStat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *MockIDBBackendMockRecorder) CreateOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockIDBBackend)(nil).CreateOrganization), arg0)
}

// CreateSupplies mocks base method
func (m *MockIDBBackend) CreateSupplies(arg0 *gorm.DB, arg1 []*models.PubSupplies) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSupplies", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSupplies indicates an expected call of CreateSupplies
func (mr *MockIDBBackendMockRecorder) CreateSupplies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSupplies", reflect.TypeOf((*MockIDBBackend)(nil).CreateSupplies), arg0, arg1)
}

// DBTransactionCommit mocks base method
func (m *MockIDBBackend) DBTransactionCommit(arg0 *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DBTransactionCommit", arg0)
}

// DBTransactionCommit indicates an expected call of DBTransactionCommit
func (mr *MockIDBBackendMockRecorder) DBTransactionCommit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBTransactionCommit", reflect.TypeOf((*MockIDBBackend)(nil).DBTransactionCommit), arg0)
}

// DBTransactionRollback mocks base method
func (m *MockIDBBackend) DBTransactionRollback(arg0 *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DBTransactionRollback", arg0)
}

// DBTransactionRollback indicates an expected call of DBTransactionRollback
func (mr *MockIDBBackendMockRecorder) DBTransactionRollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBTransactionRollback", reflect.TypeOf((*MockIDBBackend)(nil).DBTransactionRollback), arg0)
}

// GetDBTransaction mocks base method
func (m *MockIDBBackend) GetDBTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDBTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// GetDBTransaction indicates an expected call of GetDBTransaction
func (mr *MockIDBBackendMockRecorder) GetDBTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDBTransaction", reflect.TypeOf((*MockIDBBackend)(nil).GetDBTransaction))
}

// QueryFunds mocks base method
func (m *MockIDBBackend) QueryFunds(arg0, arg1, arg2, arg3 string, arg4 *structs.QueryParams) ([]*models.PubFunds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFunds", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*models.PubFunds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFunds indicates an expected call of QueryFunds
func (mr *MockIDBBackendMockRecorder) QueryFunds(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFunds", reflect.TypeOf((*MockIDBBackend)(nil).QueryFunds), arg0, arg1, arg2, arg3, arg4)
}

// QueryFundsDetail mocks base method
func (m *MockIDBBackend) QueryFundsDetail(arg0 string) (*models.FundsDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFundsDetail", arg0)
	ret0, _ := ret[0].(*models.FundsDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFundsDetail indicates an expected call of QueryFundsDetail
func (mr *MockIDBBackendMockRecorder) QueryFundsDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFundsDetail", reflect.TypeOf((*MockIDBBackend)(nil).QueryFundsDetail), arg0)
}

// QueryOrgCharities mocks base method
func (m *MockIDBBackend) QueryOrgCharities(arg0 *structs.QueryParams) ([]*structs.OrgCharitiesItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgCharities", arg0)
	ret0, _ := ret[0].([]*structs.OrgCharitiesItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgCharities indicates an expected call of QueryOrgCharities
func (mr *MockIDBBackendMockRecorder) QueryOrgCharities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgCharities", reflect.TypeOf((*MockIDBBackend)(nil).QueryOrgCharities), arg0)
}

// QueryOrgCharitiesDetail mocks base method
func (m *MockIDBBackend) QueryOrgCharitiesDetail(arg0 string) (*structs.OrgCharitiesDetailItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryOrgCharitiesDetail", arg0)
	ret0, _ := ret[0].(*structs.OrgCharitiesDetailItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOrgCharitiesDetail indicates an expected call of QueryOrgCharitiesDetail
func (mr *MockIDBBackendMockRecorder) QueryOrgCharitiesDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOrgCharitiesDetail", reflect.TypeOf((*MockIDBBackend)(nil).QueryOrgCharitiesDetail), arg0)
}

// QueryPubByUserType mocks base method
func (m *MockIDBBackend) QueryPubByUserType(arg0, arg1, arg2 string, arg3 *structs.QueryParams) ([]*structs.PubUserItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryPubByUserType", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*structs.PubUserItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryPubByUserType indicates an expected call of QueryPubByUserType
func (mr *MockIDBBackendMockRecorder) QueryPubByUserType(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryPubByUserType", reflect.TypeOf((*MockIDBBackend)(nil).QueryPubByUserType), arg0, arg1, arg2, arg3)
}

// QuerySupplies mocks base method
func (m *MockIDBBackend) QuerySupplies(arg0, arg1, arg2, arg3 string, arg4 *structs.QueryParams) ([]*models.PubSupplies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySupplies", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*models.PubSupplies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySupplies indicates an expected call of QuerySupplies
func (mr *MockIDBBackendMockRecorder) QuerySupplies(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySupplies", reflect.TypeOf((*MockIDBBackend)(nil).QuerySupplies), arg0, arg1, arg2, arg3, arg4)
}

// QuerySuppliesDetail mocks base method
func (m *MockIDBBackend) QuerySuppliesDetail(arg0 string) (*models.SuppliesDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySuppliesDetail", arg0)
	ret0, _ := ret[0].(*models.SuppliesDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySuppliesDetail indicates an expected call of QuerySuppliesDetail
func (mr *MockIDBBackendMockRecorder) QuerySuppliesDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySuppliesDetail", reflect.TypeOf((*MockIDBBackend)(nil).QuerySuppliesDetail), arg0)
}
