package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api-mock-test/api/controller"
	"go-api-mock-test/bootstrap"
	"go-api-mock-test/domain"
	mocks "go-api-mock-test/domain/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginUsecase := mocks.NewMockLoginUsecase(ctrl)

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockUser := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: string(encryptedPassword),
	}

	mockLoginUsecase.EXPECT().GetUserByEmail(gomock.Any(), mockUser.Email).Return(mockUser, nil)
	mockLoginUsecase.EXPECT().CreateAccessToken(&mockUser, gomock.Any(), gomock.Any()).Return("mockAccessToken", nil)
	mockLoginUsecase.EXPECT().CreateRefreshToken(&mockUser, gomock.Any(), gomock.Any()).Return("mockRefreshToken", nil)

	ginRouter := gin.Default()
	rec := httptest.NewRecorder()
	lc := &controller.LoginController{
		LoginUsecase: mockLoginUsecase,
		Env:          &bootstrap.Env{AccessTokenSecret: "secret", AccessTokenExpiryHour: 1, RefreshTokenSecret: "secret", RefreshTokenExpiryHour: 24},
	}
	ginRouter.POST("/login", lc.Login)

	loginRequest := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(loginRequest)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	ginRouter.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse, _ := json.Marshal(domain.LoginResponse{
		AccessToken:  "mockAccessToken",
		RefreshToken: "mockRefreshToken",
	})
	assert.JSONEq(t, string(expectedResponse), rec.Body.String())
}

func TestEmailNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginUsecase := mocks.NewMockLoginUsecase(ctrl)

	mockLoginUsecase.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(domain.User{}, errors.New("User not found with the given email"))

	ginRouter := gin.Default()
	rec := httptest.NewRecorder()
	lc := &controller.LoginController{
		LoginUsecase: mockLoginUsecase,
		Env:          &bootstrap.Env{AccessTokenSecret: "secret", AccessTokenExpiryHour: 1, RefreshTokenSecret: "secret", RefreshTokenExpiryHour: 24},
	}
	ginRouter.POST("/login", lc.Login)

	loginRequest := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(loginRequest)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	ginRouter.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestInvalidPassord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoginUsecase := mocks.NewMockLoginUsecase(ctrl)

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockUser := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: string(encryptedPassword),
	}

	mockLoginUsecase.EXPECT().GetUserByEmail(gomock.Any(), mockUser.Email).Return(mockUser, nil)

	ginRouter := gin.Default()
	rec := httptest.NewRecorder()
	lc := &controller.LoginController{
		LoginUsecase: mockLoginUsecase,
		Env:          &bootstrap.Env{AccessTokenSecret: "secret", AccessTokenExpiryHour: 1, RefreshTokenSecret: "secret", RefreshTokenExpiryHour: 24},
	}
	ginRouter.POST("/login", lc.Login)

	loginRequest := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}
	body, err := json.Marshal(loginRequest)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	ginRouter.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
