package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func TestAdminGetResvsBySeatHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	mockResvs := []*model.Reservation{
		{
			ID:       1,
			Username: "TestUser",
			RoomID:   1,
			SeatID:   1,
			Status:   util.ResvStatusUnsignin,
		},
	}
	mockDAL.EXPECT().GetResvsBySeat(gomock.Any()).Return(mockResvs, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)

	router.GET("/admin/rooms/:room_id/seats/:seat_id/reservations", middleware.AuthMiddleware(), httpHandler.AdminGetResvsBySeatHandler)

	req, err := http.NewRequest("GET", "/admin/rooms/1/seats/1/reservations", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResvs []*model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResvs)
	assert.NoError(t, err)

	assert.Equal(t, mockResvs, responseResvs)
}

func TestAdminCancelResvHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockOldResv := &model.Reservation{
		ID:       1,
		Username: "TestUser",
		RoomID:   1,
		SeatID:   1,
		Status:   util.ResvStatusUnsignin,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil)

	mockResv := &model.Reservation{
		ID:       1,
		Username: "TestUser",
		RoomID:   1,
		SeatID:   1,
		Status:   util.ResvStatusCancelled,
	}
	mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(mockResv, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/admin/rooms/:room_id/seats/:seat_id/reservations/:resv_id", middleware.AuthMiddleware(), httpHandler.AdminCancelResvHandler)

	req, err := http.NewRequest("PATCH", "/admin/rooms/1/seats/1/reservations/1", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResv *model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
	assert.NoError(t, err)

	assert.Equal(t, mockResv, responseResv)
}

func TestGetResvsByUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockResvs := []*model.Reservation{
		{
			ID:       1,
			Username: "TestUser",
			RoomID:   1,
			SeatID:   1,
			Status:   util.ResvStatusCancelled,
		},
	}
	mockDAL.EXPECT().GetResvsByUser(gomock.Any()).Return(mockResvs, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/reservations", middleware.AuthMiddleware(), httpHandler.GetResvsByUserHandler)

	req, err := http.NewRequest("GET", "/reservations", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResvs []*model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResvs)
	assert.NoError(t, err)

	assert.Equal(t, mockResvs, responseResvs)
}

func TestCreateResvHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	time1, err := time.Parse(time.RFC3339, "2022-05-30T09:00:00Z")
	assert.NoError(t, err)
	time2, err := time.Parse(time.RFC3339, "2022-05-30T12:00:00Z")
	assert.NoError(t, err)
	// Create test rooms
	resvs := []*model.Reservation{
		{
			Username:      "TestUser",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		},
		{
			Username:      "TestUser",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		},
		{
			Username:      "TestUser",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		},
	}

	// Set up the expected DAL behavior for creating rooms
	for i, resv := range resvs {
		resvCopy := *resv
		resvCopy.ID = int64(i + 1) // Set the expected ID based on the index
		mockDAL.EXPECT().CreateResv(gomock.Any()).Return(&resvCopy, nil)
	}

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.POST("/reservations", middleware.AuthMiddleware(), httpHandler.CreateResvHandler)

	// Perform the requests to create rooms
	for i, resv := range resvs {
		resvJSON, err := json.Marshal(resv)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(resvJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var responseResv *model.Reservation
		err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
		assert.NoError(t, err)

		resvCopy := resv
		resvCopy.ID = int64(i + 1)

		assert.Equal(t, resvCopy, responseResv)
	}
}

func TestCancelResvHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockOldResv := &model.Reservation{
		ID:       1,
		Username: "TestUser",
		RoomID:   1,
		SeatID:   1,
		Status:   util.ResvStatusUnsignin,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)

	mockResv := &model.Reservation{
		ID:       1,
		Username: "TestUser",
		RoomID:   1,
		SeatID:   1,
		Status:   util.ResvStatusCancelled,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockResv, nil).AnyTimes()
	mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(mockResv, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/cancel", middleware.AuthMiddleware(), httpHandler.CancelResvHandler)

	req, err := http.NewRequest("PATCH", "/reservations/1/cancel", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResv *model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
	assert.NoError(t, err)

	assert.Equal(t, mockResv, responseResv)
}

func TestSigninResvHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	time1, err := time.Parse(time.RFC3339, "2022-05-01T09:00:00Z")
	assert.NoError(t, err)
	time2 := time.Now().Add(time.Hour)

	mockOldResv := &model.Reservation{
		ID:            1,
		Username:      "TestUser",
		RoomID:        1,
		SeatID:        1,
		ResvStartTime: &time1,
		ResvEndTime:   &time2,
		Status:        util.ResvStatusUnsignin,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)

	mockResv := &model.Reservation{
		ID:            1,
		Username:      "TestUser",
		RoomID:        1,
		SeatID:        1,
		ResvStartTime: &time1,
		ResvEndTime:   &time2,
		Status:        util.ResvStatusSignined,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockResv, nil).AnyTimes()

	mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(mockResv, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/signin", middleware.AuthMiddleware(), httpHandler.SigninResvHandler)

	req, err := http.NewRequest("PATCH", "/reservations/1/signin", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResv *model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
	assert.NoError(t, err)

	assert.Equal(t, mockResv.ID, responseResv.ID)
	assert.Equal(t, mockResv.Status, responseResv.Status)
}

func TestSignoutResvHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	time1, err := time.Parse(time.RFC3339, "2022-05-01T09:00:00Z")
	assert.NoError(t, err)
	time2 := time.Now().Add(time.Hour)

	mockOldResv := &model.Reservation{
		ID:            1,
		Username:      "TestUser",
		RoomID:        1,
		SeatID:        1,
		ResvStartTime: &time1,
		ResvEndTime:   &time2,
		Status:        util.ResvStatusSignined,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)

	mockResv := &model.Reservation{
		ID:            1,
		Username:      "TestUser",
		RoomID:        1,
		SeatID:        1,
		ResvStartTime: &time1,
		ResvEndTime:   &time2,
		Status:        util.ResvStatusSignouted,
	}
	mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockResv, nil).AnyTimes()

	mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(mockResv, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/signout", middleware.AuthMiddleware(), httpHandler.SignoutResvHandler)

	req, err := http.NewRequest("PATCH", "/reservations/1/signout", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseResv *model.Reservation
	err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
	assert.NoError(t, err)

	assert.Equal(t, mockResv.ID, responseResv.ID)
	assert.Equal(t, mockResv.Status, responseResv.Status)
}
