package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func TestGetSeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminGetAllSeatsOfRoomHandler)
	router.GET("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminGetSeatByIDHandler)

	expectedSeats := []*model.Seat{
		{
			RoomID:      1,
			Number:      "Test-R1-S1",
			Location:    "test location 1",
			IsAvaliable: true,
		},
		{
			RoomID:      1,
			Number:      "Test-R1-S2",
			Location:    "test location 2",
			IsAvaliable: true,
		},
		{
			RoomID:      1,
			Number:      "Test-R1-S3",
			Location:    "test location 3",
			IsAvaliable: true,
		},
	}

	t.Run("Test get seats list", func(t *testing.T) {
		mockDAL.EXPECT().GetAllSeatsOfRoom(int64(1)).Return(expectedSeats, nil).Times(1)

		req, err := http.NewRequest("GET", "/admin/rooms/1/seats", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response []*model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedSeats, response)
	})

	t.Run("Test get a seat", func(t *testing.T) {
		mockDAL.EXPECT().GetSeatByID(gomock.Any()).Return(expectedSeats[0], nil)
		req, err := http.NewRequest("GET", "/admin/rooms/1/seats/1", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response *model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, expectedSeats[0], response)
	})

	t.Run("Test get a seat that don't exist", func(t *testing.T) {
		mockDAL.EXPECT().GetSeatByID(gomock.Any()).Return(nil, errors.New(""))
		req, err := http.NewRequest("GET", "/admin/rooms/1/seats/99", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var response *model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		var want model.Seat
		assert.Equal(t, &want, response)
	})
}

func TestCreateSeat(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.POST("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminCreateSeatsHandler)

	t.Run("Test create seat", func(t *testing.T) {
		expectedSeat := []*model.Seat{{
			RoomID:      1,
			Number:      "Test-R1-S1",
			Location:    "test location 1",
			IsAvaliable: true,
		}}
		mockDAL.EXPECT().CreateSeats(gomock.Eq(expectedSeat)).Return(nil)
		seatsJSON, _ := json.Marshal(expectedSeat)
		req, err := http.NewRequest("POST", "/admin/rooms/1/seats", bytes.NewBuffer(seatsJSON))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Test create existed seat", func(t *testing.T) {
		expectedSeat := []*model.Seat{{
			RoomID:      2,
			Number:      "Test-R1-S1",
			Location:    "test location 1",
			IsAvaliable: true,
		}}
		mockDAL.EXPECT().CreateSeats(gomock.Eq(expectedSeat)).Return(errors.New(""))
		seatsJSON, _ := json.Marshal(expectedSeat)
		req, err := http.NewRequest("POST", "/admin/rooms/2/seats", bytes.NewBuffer(seatsJSON))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestDeleteSeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.DELETE("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteSeatHandler)

	t.Run("Test delete seat", func(t *testing.T) {
		var seatID int64 = 1
		mockDAL.EXPECT().DeleteSeat(seatID).Return(nil)
		req, err := http.NewRequest("DELETE", "/admin/rooms/1/seats/1", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("Test delete seat that don't exist", func(t *testing.T) {
		var seatID int64 = 99
		mockDAL.EXPECT().DeleteSeat(seatID).Return(nil)
		req, err := http.NewRequest("DELETE", "/admin/rooms/1/seats/99", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestUpdateSeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateSeatHandler)

	t.Run("Test update seat", func(t *testing.T) {
		updatedSeat := &model.Seat{
			ID:          1,
			RoomID:      1,
			Number:      "Test-R1-S1",
			Location:    "test new location 1",
			IsAvaliable: true,
		}
		mockDAL.EXPECT().UpdateSeat(gomock.Any()).Return(updatedSeat, nil)
		reqBody, err := json.Marshal(updatedSeat)
		assert.NoError(t, err)
		req, err := http.NewRequest("PATCH", "/admin/rooms/1/seats/1", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var got *model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &got)
		assert.NoError(t, err)
		assert.Equal(t, updatedSeat, got)
	})
	t.Run("Test update seat that don't exist", func(t *testing.T) {
		updatedSeat := &model.Seat{
			ID:          99,
			RoomID:      1,
			Number:      "Test-R1-S1",
			Location:    "test new location 1",
			IsAvaliable: true,
		}
		mockDAL.EXPECT().UpdateSeat(gomock.Any()).Return(nil, errors.New(""))
		reqBody, err := json.Marshal(updatedSeat)
		assert.NoError(t, err)
		req, err := http.NewRequest("PATCH", "/admin/rooms/1/seats/99", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		var response *model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)
		var want model.Seat
		assert.Equal(t, &want, response)
	})

}

func TestSearchSeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/search", middleware.AuthMiddleware(), httpHandler.SearchSeatsHandler)

	t.Run("Test search seat", func(t *testing.T) {
		expectSeats := []*model.Seat{
			{
				RoomID:      1,
				Number:      "Test-R1-S1",
				Location:    "test location 1",
				IsAvaliable: true,
			},
			{
				RoomID:      1,
				Number:      "Test-R1-S2",
				Location:    "test location 2",
				IsAvaliable: true,
			},
			{
				RoomID:      1,
				Number:      "Test-R1-S3",
				Location:    "test location 3",
				IsAvaliable: true,
			},
		}
		mockDAL.EXPECT().SearchSeats(gomock.Any(), gomock.Any()).Return(expectSeats, nil)
		condition := "is_avaliable=true&room_id=1"
		req, err := http.NewRequest("GET", "/search?"+condition, nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var response []*model.Seat
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectSeats, response)
	})

	t.Run("Test search seat with empty query", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/search", nil)
		assert.NoError(t, err)
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
