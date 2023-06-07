package handler_test

import (
	"bytes"
	"encoding/json"
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

func TestAdminGetAllSeatsOfRoomHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	var roomID int64 = 1

	seats := []*model.Seat{
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

	mockDAL.EXPECT().GetAllSeatsOfRoom(roomID).Return(seats, nil).Times(1)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminGetAllSeatsOfRoomHandler)

	req, err := http.NewRequest("GET", "/admin/rooms/1/seats", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Seat
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, seats, response)
}

func TestAdminCreateSeatsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	// Create test rooms
	seats := []*model.Seat{
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

	mockDAL.EXPECT().CreateSeats(gomock.Eq(seats)).Return(nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.POST("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminCreateSeatsHandler)

	seatsJSON, _ := json.Marshal(seats)
	req, err := http.NewRequest("POST", "/admin/rooms/1/seats", bytes.NewBuffer(seatsJSON))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAdminDeleteSeatHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	var seatID int64 = 1
	mockDAL.EXPECT().DeleteSeat(seatID).Return(nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.DELETE("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteSeatHandler)

	req, err := http.NewRequest("DELETE", "/admin/rooms/1/seats/1", nil)
	assert.NoError(t, err)
	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAdminGetSeatByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockSeat := &model.Seat{
		ID:          1,
		RoomID:      1,
		Number:      "Test-R1-S1",
		Location:    "test location 1",
		IsAvaliable: true,
	}
	mockDAL.EXPECT().GetSeatByID(mockSeat.ID).Return(mockSeat, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminGetSeatByIDHandler)

	req, err := http.NewRequest("GET", "/admin/rooms/1/seats/1", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseSeat *model.Seat
	err = json.Unmarshal(recorder.Body.Bytes(), &responseSeat)
	assert.NoError(t, err)

	assert.Equal(t, mockSeat, responseSeat)
}

func TestAdminUpdateSeatHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	updatedSeat := &model.Seat{
		ID:          1,
		RoomID:      1,
		Number:      "Test-R1-S1",
		Location:    "test new location 1",
		IsAvaliable: true,
	}

	mockDAL.EXPECT().UpdateSeat(gomock.Any()).Return(updatedSeat, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)

	router.PATCH("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateSeatHandler)

	reqBody, err := json.Marshal(updatedSeat)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/admin/rooms/1/seats/1", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response *model.Seat
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, updatedSeat, response)
}

func TestGetAllSeatsOfRoomHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	var roomID int64 = 1

	seats := []*model.Seat{
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

	mockDAL.EXPECT().GetAllSeatsOfRoom(roomID).Return(seats, nil).Times(1)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.GetAllSeatsOfRoomHandler)

	req, err := http.NewRequest("GET", "/rooms/1/seats", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Seat
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, seats, response)
}
