package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/model"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	users := []*model.User{
		{
			Username: "TestUser",
			Password: "TestUser",
		},
		{
			Username: "TestUser2",
			Password: "TestUser2",
		},
		{
			Username: "TestUser3",
			Password: "TestUser3",
		},
	}

	for i, user := range users {
		userCopy := *user
		userCopy.ID = uint(i + 1)
		mockDAL.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&userCopy, nil)
	}

	httpHandler, router, _ := getTestRouter(t, mockDAL)
	router.POST("/register", httpHandler.RegisterHandler)

	for _, user := range users {
		roomJSON, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(roomJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockDAL.EXPECT().CheckUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)

	httpHandler, router, _ := getTestRouter(t, mockDAL)
	router.POST("/login", httpHandler.LoginHandler)

	user := &model.User{
		Username: "TestUser",
		Password: "TestUser",
	}
	userJson, err := json.Marshal(user)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
