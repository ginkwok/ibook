package test

import (
	"bytes"
	"encoding/json"
	"errors"
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

func TestCreateReservation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)

	router.POST("/reservations", middleware.AuthMiddleware(), httpHandler.CreateResvHandler)
	t.Run("Test create reservation", func(t *testing.T) {
		time1, err := time.Parse(time.RFC3339, "2022-05-30T09:00:00Z")
		assert.NoError(t, err)
		time2, err := time.Parse(time.RFC3339, "2022-05-30T12:00:00Z")
		assert.NoError(t, err)
		resv := &model.Reservation{
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		}
		mockDAL.EXPECT().CreateResv(gomock.Any()).Return(resv, nil)
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
		assert.Equal(t, resv, responseResv)
	})

	t.Run("Test create reservation, but seat don't exist", func(t *testing.T) {
		time1, err := time.Parse(time.RFC3339, "2022-05-30T09:00:00Z")
		assert.NoError(t, err)
		time2, err := time.Parse(time.RFC3339, "2022-05-30T12:00:00Z")
		assert.NoError(t, err)
		resv := &model.Reservation{
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        99,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		}
		mockDAL.EXPECT().CreateResv(gomock.Any()).Return(nil, errors.New(""))
		resvJSON, err := json.Marshal(resv)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(resvJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test create reservation, but seat has been reserved", func(t *testing.T) {
		time1, err := time.Parse(time.RFC3339, "2022-05-30T09:00:00Z")
		assert.NoError(t, err)
		time2, err := time.Parse(time.RFC3339, "2022-05-30T12:00:00Z")
		assert.NoError(t, err)
		resv := &model.Reservation{
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		}
		mockDAL.EXPECT().CreateResv(gomock.Any()).Return(nil, errors.New(""))
		resvJSON, err := json.Marshal(resv)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(resvJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test create reservation, but repeated time", func(t *testing.T) {
		time1, err := time.Parse(time.RFC3339, "2022-03-30T09:00:00Z")
		assert.NoError(t, err)
		time2, err := time.Parse(time.RFC3339, "2022-03-30T12:00:00Z")
		assert.NoError(t, err)
		resv := &model.Reservation{
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		}
		mockDAL.EXPECT().CreateResv(gomock.Any()).Return(nil, errors.New(""))
		resvJSON, err := json.Marshal(resv)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(resvJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test create reservation, but illegal time", func(t *testing.T) {
		time1, err := time.Parse(time.RFC3339, "2022-05-30T10:00:00Z")
		assert.NoError(t, err)
		time2, err := time.Parse(time.RFC3339, "2022-05-30T09:00:00Z")
		assert.NoError(t, err)
		resv := &model.Reservation{
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
		}
		resvJSON, err := json.Marshal(resv)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/reservations", bytes.NewBuffer(resvJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestCancelReservation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/cancel", middleware.AuthMiddleware(), httpHandler.CancelResvHandler)

	t.Run("Test cancel reservation", func(t *testing.T) {
		oldResv := &model.Reservation{
			ID:       1,
			Username: "TestUser1",
			RoomID:   1,
			SeatID:   1,
			Status:   util.ResvStatusUnsignin,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(oldResv, nil).Times(2)

		expectResv := &model.Reservation{
			ID:       1,
			Username: "TestUser1",
			RoomID:   1,
			SeatID:   1,
			Status:   util.ResvStatusCancelled,
		}
		mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(expectResv, nil).AnyTimes()

		req, err := http.NewRequest("PATCH", "/reservations/1/cancel", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var responseResv *model.Reservation
		err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
		assert.NoError(t, err)

		assert.Equal(t, expectResv, responseResv)
	})

	t.Run("Test cancel reservation that don't exist", func(t *testing.T) {

		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(nil, errors.New(""))

		req, err := http.NewRequest("PATCH", "/reservations/1/cancel", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test cancel reservation that has already started", func(t *testing.T) {
		resvStartTime := time.Now().Add(-time.Hour)
		oldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			Status:        util.ResvStatusUnsignin,
			ResvStartTime: &resvStartTime,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(oldResv, nil).Times(1)
		mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(nil, errors.New("")).AnyTimes()
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(nil, errors.New("")).Times(1)

		req, err := http.NewRequest("PATCH", "/reservations/1/cancel", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetReservationHistory(t *testing.T) {

}

func TestGetSeatReservation(t *testing.T) {

}

func TestCancelSeatReservation(t *testing.T) {

}
