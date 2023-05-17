package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func (h *handlerStruct) AdminGetAllRoomsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	rooms, err := h.svc.GetAllRooms(ctx)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *handlerStruct) AdminCreateRoomHandler(c *gin.Context) {
	ctx := c.Request.Context()
	// logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	// _, ok := c.Get("username")
	// if !ok {
	// 	err := errors.New("invalid credentials")
	// 	logger.Errorln(err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }

	var room *model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		// logger.Errorln(err)
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.ID = 0

	room, err := h.svc.CreateRoom(ctx, room)
	if err != nil {
		// logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *handlerStruct) AdminDeleteRoomHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("room_id")
	if idStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.DeleteRoom(ctx, id)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *handlerStruct) AdminGetRoomByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("room_id")
	if idStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.svc.GetRoomByID(ctx, id)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *handlerStruct) AdminUpdateRoomHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("room_id")
	if idStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var room *model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.ID = id

	room, err = h.svc.UpdateRoom(ctx, room)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *handlerStruct) GetAllRoomsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	rooms, err := h.svc.GetAllRooms(ctx)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}
