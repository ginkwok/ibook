package test

import (
	"testing"
)

func TestUserRegister(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockDAL := mocks.NewMockDal(ctrl)

	// user := model.User{

	// 	Username:  "TestUser1",
	// 	Password:  "TestUserPass1",
	// 	Email:     "test1@test.com",
	// 	NoticeURL: "https://test1.test.com/test1",
	// }

	// httpHandler, router, _ := getTestRouter(t, mockDAL)
	// router.POST("/register", httpHandler.RegisterHandler)

	// // t.Run("Register User", func(t *testing.T) {
	// userCopy := user

	// mockDAL.EXPECT().CreateUser(gomock.Any()).Return(&userCopy, nil)

	// userJson, _ := json.Marshal(&user)
	// req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
	// req.Header.Set("Content-Type", "application/json")

	// got := httptest.NewRecorder()
	// router.ServeHTTP(got, req.WithContext(context.Background()))

	// assert.Equal(t, http.StatusOK, got.Code)

	// })

	// t.Run("Register Repeat User", func(t *testing.T) {
	// 	userCopy := user

	// 	mockDAL.EXPECT().CreateUser(gomock.Any()).Return(&userCopy, nil)

	// 	userJson, _ := json.Marshal(&user)
	// 	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
	// 	req.Header.Set("Content-Type", "application/json")

	// 	got := httptest.NewRecorder()
	// 	router.ServeHTTP(got, req.WithContext(context.Background()))

	// 	assert.Equal(t, http.StatusOK, got.Code)
	// userCopy := user

	// mockDAL.EXPECT().CreateUser(gomock.Any()).Return(&userCopy, nil)

	// userJson, _ := json.Marshal(&user)
	// req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
	// req.Header.Set("Content-Type", "application/json")

	// got := httptest.NewRecorder()
	// router.ServeHTTP(got, req.WithContext(context.Background()))

	// mockDAL.EXPECT().CreateUser(gomock.Any()).Return(nil, errors.New(""))

	// req, _ = http.NewRequest("POST", "/register", bytes.NewBuffer(userJson))
	// req.Header.Set("Content-Type", "application/json")

	// got = httptest.NewRecorder()
	// router.ServeHTTP(got, req.WithContext(context.Background()))

	// assert.Equal(t, http.StatusBadRequest, got.Code)
	// })
}

func TestUserLogin(t *testing.T) {
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// mockDAL := mocks.NewMockDal(ctrl)

	// mockDAL.EXPECT().CheckUser(gomock.Any(), gomock.Any()).Return(true, nil)

	// httpHandler, router, _ := getTestRouter(t, mockDAL)
	// router.POST("/login", httpHandler.LoginHandler)

	// user := &model.User{
	// 	Username: "TestUser",
	// 	Password: "TestUser",
	// }
	// userJson, err := json.Marshal(user)
	// assert.NoError(t, err)

	// req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userJson))
	// assert.NoError(t, err)

	// recorder := httptest.NewRecorder()

	// router.ServeHTTP(recorder, req)

	// assert.Equal(t, http.StatusOK, recorder.Code)
}
