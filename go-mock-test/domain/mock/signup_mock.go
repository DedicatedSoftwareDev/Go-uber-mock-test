// Code generated by MockGen. DO NOT EDIT.
// Source: ./ (interfaces: SignupUsecase)
//
// Generated by this command:
//
//	mockgen -destination mock/signup_mock.go ./ SignupUsecase
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "go-api-mock-test/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSignupUsecase is a mock of SignupUsecase interface.
type MockSignupUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockSignupUsecaseMockRecorder
	isgomock struct{}
}

// MockSignupUsecaseMockRecorder is the mock recorder for MockSignupUsecase.
type MockSignupUsecaseMockRecorder struct {
	mock *MockSignupUsecase
}

// NewMockSignupUsecase creates a new mock instance.
func NewMockSignupUsecase(ctrl *gomock.Controller) *MockSignupUsecase {
	mock := &MockSignupUsecase{ctrl: ctrl}
	mock.recorder = &MockSignupUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignupUsecase) EXPECT() *MockSignupUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSignupUsecase) Create(c context.Context, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSignupUsecaseMockRecorder) Create(c, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSignupUsecase)(nil).Create), c, user)
}

// CreateAccessToken mocks base method.
func (m *MockSignupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", user, secret, expiry)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockSignupUsecaseMockRecorder) CreateAccessToken(user, secret, expiry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockSignupUsecase)(nil).CreateAccessToken), user, secret, expiry)
}

// CreateRefreshToken mocks base method.
func (m *MockSignupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken", user, secret, expiry)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockSignupUsecaseMockRecorder) CreateRefreshToken(user, secret, expiry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockSignupUsecase)(nil).CreateRefreshToken), user, secret, expiry)
}

// GetUserByEmail mocks base method.
func (m *MockSignupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", c, email)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockSignupUsecaseMockRecorder) GetUserByEmail(c, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockSignupUsecase)(nil).GetUserByEmail), c, email)
}
