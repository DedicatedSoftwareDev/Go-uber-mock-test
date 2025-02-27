// Code generated by MockGen. DO NOT EDIT.
// Source: ./ (interfaces: RefreshTokenUsecase)
//
// Generated by this command:
//
//	mockgen -destination mock/refresh_token_mock.go ./ RefreshTokenUsecase
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "go-api-mock-test/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRefreshTokenUsecase is a mock of RefreshTokenUsecase interface.
type MockRefreshTokenUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenUsecaseMockRecorder
	isgomock struct{}
}

// MockRefreshTokenUsecaseMockRecorder is the mock recorder for MockRefreshTokenUsecase.
type MockRefreshTokenUsecaseMockRecorder struct {
	mock *MockRefreshTokenUsecase
}

// NewMockRefreshTokenUsecase creates a new mock instance.
func NewMockRefreshTokenUsecase(ctrl *gomock.Controller) *MockRefreshTokenUsecase {
	mock := &MockRefreshTokenUsecase{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshTokenUsecase) EXPECT() *MockRefreshTokenUsecaseMockRecorder {
	return m.recorder
}

// CreateAccessToken mocks base method.
func (m *MockRefreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", user, secret, expiry)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockRefreshTokenUsecaseMockRecorder) CreateAccessToken(user, secret, expiry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockRefreshTokenUsecase)(nil).CreateAccessToken), user, secret, expiry)
}

// CreateRefreshToken mocks base method.
func (m *MockRefreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken", user, secret, expiry)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockRefreshTokenUsecaseMockRecorder) CreateRefreshToken(user, secret, expiry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockRefreshTokenUsecase)(nil).CreateRefreshToken), user, secret, expiry)
}

// ExtractIDFromToken mocks base method.
func (m *MockRefreshTokenUsecase) ExtractIDFromToken(requestToken, secret string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractIDFromToken", requestToken, secret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractIDFromToken indicates an expected call of ExtractIDFromToken.
func (mr *MockRefreshTokenUsecaseMockRecorder) ExtractIDFromToken(requestToken, secret any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractIDFromToken", reflect.TypeOf((*MockRefreshTokenUsecase)(nil).ExtractIDFromToken), requestToken, secret)
}

// GetUserByID mocks base method.
func (m *MockRefreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", c, id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockRefreshTokenUsecaseMockRecorder) GetUserByID(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRefreshTokenUsecase)(nil).GetUserByID), c, id)
}
