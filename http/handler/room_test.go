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
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func TestAdminGetAllRoomsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	expectedRooms := []*model.Room{
		{
			ID:          1,
			Name:        "test room 1",
			Capacity:    10,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			ID:          2,
			Name:        "test room 2",
			Capacity:    20,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			ID:          3,
			Name:        "test room 3",
			Capacity:    30,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
	}

	mockDAL.EXPECT().GetAllRooms().Return(expectedRooms, nil).Times(1)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminGetAllRoomsHandler)

	req, err := http.NewRequest("GET", "/admin/rooms", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Room
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedRooms, response)
}

func TestCreateRoomHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	// Create test rooms
	rooms := []*model.Room{

		{
			Name:        "test room 1",
			Capacity:    10,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			Name:        "test room 2",
			Capacity:    20,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			Name:        "test room 3",
			Capacity:    30,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
	}

	// Set up the expected DAL behavior for creating rooms
	for i, room := range rooms {
		roomCopy := *room
		roomCopy.ID = int64(i + 1) // Set the expected ID based on the index
		mockDAL.EXPECT().CreateRoom(room).Return(&roomCopy, nil)
	}

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.POST("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminCreateRoomHandler)

	// Perform the requests to create rooms
	for _, room := range rooms {
		roomJSON, _ := json.Marshal(room)
		req, _ := http.NewRequest("POST", "/admin/rooms", bytes.NewBuffer(roomJSON))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestAdminDeleteRoomHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	var roomID int64 = 1

	mockDAL.EXPECT().DeleteRoom(roomID).Return(nil)
	mockDAL.EXPECT().DeleteSeatsOfRoom(roomID).Return(nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.DELETE("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteRoomHandler)

	req, err := http.NewRequest("DELETE", "/admin/rooms/1", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAdminGetRoomByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	mockRoom := &model.Room{ID: 1, Name: "Test Room"}
	mockDAL.EXPECT().GetRoomByID(gomock.Any()).Return(mockRoom, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminGetRoomByIDHandler)

	req, err := http.NewRequest("GET", "/admin/rooms/1", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseRoom model.Room
	err = json.Unmarshal(recorder.Body.Bytes(), &responseRoom)
	assert.NoError(t, err)

	assert.Equal(t, mockRoom.ID, responseRoom.ID)
	assert.Equal(t, mockRoom.Name, responseRoom.Name)
}

func TestAdminUpdateRoomHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	updatedRoom := &model.Room{
		ID:          1,
		Name:        "Updated Room",
		Capacity:    30,
		OpeningTime: "08:00:00",
		ClosingTime: "22:30:00",
		Location:    "test location",
		Description: "test description",
		IsAvaliable: true,
	}

	mockDAL.EXPECT().UpdateRoom(gomock.Any()).Return(updatedRoom, nil)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateRoomHandler)

	reqBody, err := json.Marshal(updatedRoom)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/admin/rooms/1", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response *model.Room
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, updatedRoom, response)
}

func TestGetAllRoomsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	expectedRooms := []*model.Room{
		{
			ID:          1,
			Name:        "test room 1",
			Capacity:    10,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			ID:          2,
			Name:        "test room 2",
			Capacity:    20,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
		{
			ID:          3,
			Name:        "test room 3",
			Capacity:    30,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		},
	}

	mockDAL.EXPECT().GetAllRooms().Return(expectedRooms, nil).Times(1)

	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/rooms", middleware.AuthMiddleware(), httpHandler.GetAllRoomsHandler)

	req, err := http.NewRequest("GET", "/rooms", nil)
	assert.NoError(t, err)

	req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Room
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedRooms, response)
}
