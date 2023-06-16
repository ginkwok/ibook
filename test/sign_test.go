package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/signin", middleware.AuthMiddleware(), httpHandler.SigninResvHandler)

	t.Run("Test signin", func(t *testing.T) {
		time1 := time.Now().Add(-time.Hour)
		time2 := time.Now().Add(time.Hour)
		mockOldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusUnsignin,
		}
		expectResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusSignined,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)
		mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(expectResv, nil).Times(1)
		req, err := http.NewRequest("PATCH", "/reservations/1/signin", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseResv *model.Reservation
		err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
		assert.NoError(t, err)
		assert.Equal(t, expectResv.ID, responseResv.ID)
		assert.Equal(t, expectResv.Status, responseResv.Status)
	})

	t.Run("Test signin, but be late", func(t *testing.T) {
		time1 := time.Now().Add(-time.Hour)
		time2 := time.Now().Add(-time.Hour)
		mockOldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusUnsignin,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)
		req, err := http.NewRequest("PATCH", "/reservations/1/signin", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test signin, but be not yet started", func(t *testing.T) {
		time1 := time.Now().Add(time.Hour)
		time2 := time1.Add(time.Hour)
		mockOldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusUnsignin,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)
		req, err := http.NewRequest("PATCH", "/reservations/2/signin", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test signin, but repeat", func(t *testing.T) {
		time1 := time.Now().Add(-time.Hour)
		time2 := time.Now().Add(time.Hour)
		mockOldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusSignined,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(mockOldResv, nil).Times(2)
		req, err := http.NewRequest("PATCH", "/reservations/3/signin", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestSignout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/reservations/:resv_id/signout", middleware.AuthMiddleware(), httpHandler.SignoutResvHandler)

	t.Run("Test signout", func(t *testing.T) {
		time1 := time.Now().Add(-time.Hour)
		time2 := time.Now().Add(time.Hour)
		oldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusSignined,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(oldResv, nil).Times(2)
		expectResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusSignouted,
		}
		mockDAL.EXPECT().UpdateResv(gomock.Any()).Return(expectResv, nil).Times(1)
		req, err := http.NewRequest("PATCH", "/reservations/1/signout", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var responseResv *model.Reservation
		err = json.Unmarshal(recorder.Body.Bytes(), &responseResv)
		assert.NoError(t, err)
		assert.Equal(t, expectResv.ID, responseResv.ID)
		assert.Equal(t, expectResv.Status, responseResv.Status)
	})

	t.Run("Test signout, but be not yet started", func(t *testing.T) {
		time1 := time.Now().Add(time.Hour)
		time2 := time1.Add(time.Hour)
		oldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusUnsignin,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(oldResv, nil).Times(2)
		req, err := http.NewRequest("PATCH", "/reservations/1/signout", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Test signout, but repeated", func(t *testing.T) {
		time1 := time.Now().Add(-time.Hour)
		time2 := time1.Add(2 * time.Hour)
		oldResv := &model.Reservation{
			ID:            1,
			Username:      "TestUser1",
			RoomID:        1,
			SeatID:        1,
			ResvStartTime: &time1,
			ResvEndTime:   &time2,
			Status:        util.ResvStatusSignouted,
		}
		mockDAL.EXPECT().GetResvByID(gomock.Any()).Return(oldResv, nil).Times(2)
		req, err := http.NewRequest("PATCH", "/reservations/1/signout", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
