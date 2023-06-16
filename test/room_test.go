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
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.GET("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminGetAllRoomsHandler)
	router.GET("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminGetRoomByIDHandler)

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

	t.Run("Test get all rooms", func(t *testing.T) {
		mockDAL.EXPECT().GetAllRooms().Return(expectedRooms, nil).Times(1)

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
	})

	t.Run("Test get a room", func(t *testing.T) {
		mockDAL.EXPECT().GetRoomByID(gomock.Any()).Return(expectedRooms[0], nil)
		req, err := http.NewRequest("GET", "/admin/rooms/1", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var responseRoom *model.Room
		err = json.Unmarshal(recorder.Body.Bytes(), &responseRoom)
		assert.NoError(t, err)

		assert.Equal(t, expectedRooms[0], responseRoom)
	})

	t.Run("Test get a room that don't exist", func(t *testing.T) {
		mockDAL.EXPECT().GetRoomByID(gomock.Any()).Return(nil, errors.New(""))
		req, err := http.NewRequest("GET", "/admin/rooms/99", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var responseRoom *model.Room
		err = json.Unmarshal(recorder.Body.Bytes(), &responseRoom)
		assert.NoError(t, err)

		var want model.Room
		assert.Equal(t, &want, responseRoom)
	})

}

func TestCreateRoom(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.POST("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminCreateRoomHandler)

	room := model.Room{
		ID:          1,
		Name:        "test room 1",
		Capacity:    10,
		OpeningTime: "08:00:00",
		ClosingTime: "22:30:00",
		Location:    "test location",
		Description: "test description",
		IsAvaliable: true,
	}

	t.Run("Test create room", func(t *testing.T) {
		mockDAL.EXPECT().CreateRoom(gomock.Any()).Return(&room, nil)
		roomJSON, err := json.Marshal(&room)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/admin/rooms", bytes.NewBuffer(roomJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusOK, w.Code)

		var got model.Room
		err = json.Unmarshal(w.Body.Bytes(), &got)
		assert.NoError(t, err)

		assert.Equal(t, &room, &got)
	})

	t.Run("Test create existed room", func(t *testing.T) {
		mockDAL.EXPECT().CreateRoom(gomock.Any()).Return(nil, errors.New(""))
		roomJSON, err := json.Marshal(&room)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/admin/rooms", bytes.NewBuffer(roomJSON))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var got model.Room
		err = json.Unmarshal(w.Body.Bytes(), &got)
		assert.NoError(t, err)

		var want model.Room
		assert.Equal(t, &want, &got)
	})

}

func TestDeleteRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.DELETE("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteRoomHandler)

	t.Run("Test delete room", func(t *testing.T) {
		var roomID int64 = 1

		mockDAL.EXPECT().DeleteRoom(roomID).Return(nil)
		mockDAL.EXPECT().DeleteSeatsOfRoom(roomID).Return(nil)

		req, err := http.NewRequest("DELETE", "/admin/rooms/1", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Test delete room that don't exist", func(t *testing.T) {
		var roomID int64 = 99

		mockDAL.EXPECT().DeleteRoom(roomID).Return(nil)
		mockDAL.EXPECT().DeleteSeatsOfRoom(roomID).Return(nil)

		req, err := http.NewRequest("DELETE", "/admin/rooms/99", nil)
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

}

func TestUpdateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)
	httpHandler, router, token := getTestRouter(t, mockDAL)
	router.PATCH("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateRoomHandler)

	t.Run("Test update room", func(t *testing.T) {
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

		reqBody, err := json.Marshal(updatedRoom)
		assert.NoError(t, err)

		req, err := http.NewRequest("PATCH", "/admin/rooms/1", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var got *model.Room
		err = json.Unmarshal(recorder.Body.Bytes(), &got)
		assert.NoError(t, err)

		assert.Equal(t, updatedRoom, got)
	})

	t.Run("Test update room that don't exist", func(t *testing.T) {
		mockDAL.EXPECT().UpdateRoom(gomock.Any()).Return(nil, errors.New(""))

		updatedRoom := &model.Room{
			ID:          99,
			Name:        "Updated Room",
			Capacity:    30,
			OpeningTime: "08:00:00",
			ClosingTime: "22:30:00",
			Location:    "test location",
			Description: "test description",
			IsAvaliable: true,
		}
		reqBody, err := json.Marshal(updatedRoom)
		assert.NoError(t, err)
		req, err := http.NewRequest("PATCH", "/admin/rooms/99", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		req.Header.Set(util.HTTP_HAED_AUTH, util.HTTP_HAED_AUTH_BEAR+token)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var got *model.Room
		err = json.Unmarshal(recorder.Body.Bytes(), &got)
		assert.NoError(t, err)

		var want model.Room
		assert.Equal(t, &want, got)
	})
}
