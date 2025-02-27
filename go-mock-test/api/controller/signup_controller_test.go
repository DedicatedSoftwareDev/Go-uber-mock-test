package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-api-mock-test/api/controller"
	"go-api-mock-test/bootstrap"
	"go-api-mock-test/domain"
	mocks "go-api-mock-test/domain/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestSignupSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSignupUsecase := mocks.NewMockSignupUsecase(ctrl)

	mockSignupUsecase.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("User not found with the give email"))
	mockSignupUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockSignupUsecase.EXPECT().CreateAccessToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("mockAccessToken", nil)
	mockSignupUsecase.EXPECT().CreateRefreshToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("mockRefreshToken", nil)

	ginRouter := gin.Default()
	rec := httptest.NewRecorder()
	sc := &controller.SignupController{
		SignupUsecase: mockSignupUsecase,
		Env:           &bootstrap.Env{AccessTokenSecret: "secret", AccessTokenExpiryHour: 1, RefreshTokenSecret: "secret", RefreshTokenExpiryHour: 24},
	}
	ginRouter.POST("/signup", sc.Signup)

	signupRequest := domain.SignupRequest{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "signuppassward",
	}
	body, err := json.Marshal(signupRequest)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	ginRouter.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse, _ := json.Marshal(domain.LoginResponse{
		AccessToken:  "mockAccessToken",
		RefreshToken: "mockRefreshToken",
	})
	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
}

func TestSignupUserExisted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSignupUsecase := mocks.NewMockSignupUsecase(ctrl)

	mockSignupUsecase.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, nil)

	ginRouter := gin.Default()
	rec := httptest.NewRecorder()
	sc := &controller.SignupController{
		SignupUsecase: mockSignupUsecase,
		Env:           &bootstrap.Env{AccessTokenSecret: "secret", AccessTokenExpiryHour: 1, RefreshTokenSecret: "secret", RefreshTokenExpiryHour: 24},
	}
	ginRouter.POST("/signup", sc.Signup)

	signupRequest := domain.SignupRequest{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "signuppassward",
	}
	body, err := json.Marshal(signupRequest)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	ginRouter.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusConflict, rec.Code)

	expectedResponse, _ := json.Marshal(domain.ErrorResponse{
		Message: "User already exists with the given email",
	})
	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
}
