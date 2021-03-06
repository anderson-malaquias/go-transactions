// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecase/interface.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	gomock "github.com/golang/mock/gomock"
	account "github.com/rafaelbmateus/go-transactions/internal/domain/entity/account"
	transaction "github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// NewAccount mocks base method
func (m *MockUseCase) NewAccount(arg0 *account.Account) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAccount indicates an expected call of NewAccount
func (mr *MockUseCaseMockRecorder) NewAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockUseCase)(nil).NewAccount), arg0)
}

// GetAccount mocks base method
func (m *MockUseCase) GetAccount(id int) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", id)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockUseCaseMockRecorder) GetAccount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockUseCase)(nil).GetAccount), id)
}

// GetTransactions mocks base method
func (m *MockUseCase) GetTransactions() ([]*transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions")
	ret0, _ := ret[0].([]*transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions
func (mr *MockUseCaseMockRecorder) GetTransactions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockUseCase)(nil).GetTransactions))
}

// RegisterTransaction mocks base method
func (m *MockUseCase) RegisterTransaction(arg0 *transaction.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterTransaction indicates an expected call of RegisterTransaction
func (mr *MockUseCaseMockRecorder) RegisterTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTransaction", reflect.TypeOf((*MockUseCase)(nil).RegisterTransaction), arg0)
}
