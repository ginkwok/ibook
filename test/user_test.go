package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, _ := getTestRouter(t, mockDAL)
	router.POST("/register", httpHandler.RegisterHandler)

	user := model.User{
		Username:  "TestUser1",
		Password:  "TestUserPass1",
		Email:     "test1@test.com",
		NoticeURL: "https://test1.test.com/test1",
	}

	t.Run("Register user", func(t *testing.T) {
		userCopy := user
		mockDAL.EXPECT().CreateUser(gomock.Any()).Return(&userCopy, nil)

		userJson, err := json.Marshal(&user)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		got := httptest.NewRecorder()
		router.ServeHTTP(got, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusOK, got.Code)
	})

	t.Run("Register repeat user", func(t *testing.T) {
		mockDAL.EXPECT().CreateUser(gomock.Any()).Return(nil, errors.New(""))

		userJson, err := json.Marshal(&user)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		got := httptest.NewRecorder()
		router.ServeHTTP(got, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusBadRequest, got.Code)
	})
}

func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, _ := getTestRouter(t, mockDAL)
	router.POST("/login", httpHandler.LoginHandler)

	t.Run("Test user login", func(t *testing.T) {
		mockDAL.EXPECT().CheckUser(gomock.Any(), gomock.Any()).Return(true, nil)

		user := &model.User{
			Username: "TestUser1",
			Password: "TestUserPass1",
		}
		userJson, err := json.Marshal(user)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		assert.NoError(t, err)

		got := httptest.NewRecorder()

		router.ServeHTTP(got, req)

		assert.Equal(t, http.StatusOK, got.Code)
	})

	t.Run("Test unregistered user login", func(t *testing.T) {

		mockDAL.EXPECT().CheckUser(gomock.Any(), gomock.Any()).Return(false, nil)

		user := &model.User{
			Username: "TestUser99",
			Password: "TestUserPass99",
		}
		userJson, err := json.Marshal(user)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
		assert.NoError(t, err)

		got := httptest.NewRecorder()

		router.ServeHTTP(got, req)

		assert.Equal(t, http.StatusUnauthorized, got.Code)
	})
}
