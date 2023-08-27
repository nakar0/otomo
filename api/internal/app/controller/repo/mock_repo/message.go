// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	model "otomo/internal/app/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMessageRepository) Add(ctx context.Context, msg *model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMessageRepositoryMockRecorder) Add(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMessageRepository)(nil).Add), ctx, msg)
}

// DeleteByIDAndUserID mocks base method.
func (m *MockMessageRepository) DeleteByIDAndUserID(ctx context.Context, id, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByIDAndUserID", ctx, id, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByIDAndUserID indicates an expected call of DeleteByIDAndUserID.
func (mr *MockMessageRepositoryMockRecorder) DeleteByIDAndUserID(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByIDAndUserID", reflect.TypeOf((*MockMessageRepository)(nil).DeleteByIDAndUserID), ctx, id, userID)
}