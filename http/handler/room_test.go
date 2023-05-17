// http/handler/room_test.go
package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	_ "github.com/ginkwok/ibook/config"
	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/handler"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/service"
)

func TestCreateRoomHandler(t *testing.T) {
	// Create a mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock DAL
	mockDAL := mocks.NewMockDal(ctrl)

	// Create a service with the mock DAL
	mockService := service.NewService(mockDAL, nil)

	// Create a handler with the mock service
	httpHandler := handler.NewHandler(mockService)

	// Create a test router
	router := gin.Default()
	router.POST("/admin/rooms", httpHandler.AdminCreateRoomHandler)

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
		mockDAL.EXPECT().CreateRoom(gomock.Any(), room).Return(&roomCopy, nil)
	}

	// Perform the requests to create rooms
	for _, room := range rooms {
		roomJSON, _ := json.Marshal(room)
		req, _ := http.NewRequest("POST", "/admin/rooms", bytes.NewBuffer(roomJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req.WithContext(context.Background()))

		assert.Equal(t, http.StatusOK, w.Code)
	}
}

// func TestCreateRoomHandler1(t *testing.T) {
// 	// Create a mock controller
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Create a mock DAL
// 	mockDAL := mocks.NewMockDal(ctrl)

// 	// Set up the expected DAL behavior
// 	room := &model.Room{
// 		ID:          1,
// 		Name:        "test room",
// 		Capacity:    10,
// 		OpeningTime: "08:00:00",
// 		ClosingTime: "22:30:00",
// 		Location:    "test location",
// 		Description: "test description",
// 		IsAvaliable: true,
// 	}
// 	mockDAL.EXPECT().CreateRoom(gomock.Any(), room).Return(room, nil)

// 	// Create a service with the mock DAL
// 	mockService := service.NewService(mockDAL, nil)

// 	// Create a handler with the mock service
// 	httpHandler := handler.NewHandler(mockService)

// 	// Create a test router
// 	router := gin.Default()
// 	router.POST("/admin/rooms", httpHandler.AdminCreateRoomHandler)

// 	// Create a test request
// 	room.ID = 0
// 	roomData, _ := json.Marshal(&room)
// 	roomJSON := string(roomData)

// 	req, _ := http.NewRequest("POST", "/admin/rooms", strings.NewReader(roomJSON))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Perform the request
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req.WithContext(context.Background()))

// 	// Check the response status code
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Check the response body
// 	room.ID = 0
// 	expectedResponseData, _ := json.Marshal(&room)
// 	expectedResponse := string(expectedResponseData)
// 	assert.Equal(t, expectedResponse, w.Body.String())
// }
