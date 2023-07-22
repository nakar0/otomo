// Code generated by MockGen. DO NOT EDIT.
// Source: message_with_otomo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	message "otomo/internal/app/domain/entity/message"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageWithOtomoRepository is a mock of MessageWithOtomoRepository interface.
type MockMessageWithOtomoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageWithOtomoRepositoryMockRecorder
}

// MockMessageWithOtomoRepositoryMockRecorder is the mock recorder for MockMessageWithOtomoRepository.
type MockMessageWithOtomoRepositoryMockRecorder struct {
	mock *MockMessageWithOtomoRepository
}

// NewMockMessageWithOtomoRepository creates a new mock instance.
func NewMockMessageWithOtomoRepository(ctrl *gomock.Controller) *MockMessageWithOtomoRepository {
	mock := &MockMessageWithOtomoRepository{ctrl: ctrl}
	mock.recorder = &MockMessageWithOtomoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageWithOtomoRepository) EXPECT() *MockMessageWithOtomoRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMessageWithOtomoRepository) Add(ctx context.Context, msg *message.MessageWithOtomo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMessageWithOtomoRepositoryMockRecorder) Add(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMessageWithOtomoRepository)(nil).Add), ctx, msg)
}

// DeleteByID mocks base method.
func (m *MockMessageWithOtomoRepository) DeleteByID(ctx context.Context, id message.MessageWithOtomoID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockMessageWithOtomoRepositoryMockRecorder) DeleteByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockMessageWithOtomoRepository)(nil).DeleteByID), ctx, id)
}
