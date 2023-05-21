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

func (h *HandlerStruct) AdminGetAllSeatsOfRoomHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seats, err := h.svc.GetAllSeatsOfRoom(ctx, roomID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (h *HandlerStruct) AdminCreateSeatHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var seat *model.Seat
	if err := c.ShouldBindJSON(&seat); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat.ID = 0
	seat.RoomID = roomID

	seat, err = h.svc.CreateSeat(ctx, seat)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) AdminCreateSeatsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var seats []*model.Seat
	if err := c.ShouldBindJSON(&seats); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range seats {
		v.ID = 0
		v.RoomID = roomID
	}

	err = h.svc.CreateSeats(ctx, seats)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *HandlerStruct) AdminDeleteSeatHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seatID, err := strconv.ParseInt(seatIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.DeleteSeat(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *HandlerStruct) AdminGetSeatByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seatID, err := strconv.ParseInt(seatIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, err := h.svc.GetSeatByID(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) AdminUpdateSeatHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatIDStr := c.Param("seat_id")
	if seatIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seatID, err := strconv.ParseInt(seatIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var seat *model.Seat
	if err := c.ShouldBindJSON(&seat); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat.ID = seatID
	seat.RoomID = roomID

	seat, err = h.svc.UpdateSeat(ctx, seat)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
}

func (h *HandlerStruct) GetAllSeatsOfRoomHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get("username")
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	roomIDStr := c.Param("room_id")
	if roomIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roomID, err := strconv.ParseInt(roomIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seats, err := h.svc.GetAllSeatsOfRoom(ctx, roomID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (h *HandlerStruct) SearchSeatsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	var conditions []string
	var args []interface{}

	if roomID := c.Query("room_id"); roomID != "" {
		conditions = append(conditions, "room_id = ?")
		args = append(args, roomID)
	}

	if isAvailable := c.Query("is_avaliable"); isAvailable != "" {
		conditions = append(conditions, "is_avaliable = ?")
		args = append(args, isAvailable)
	}

	if tagWindow := c.Query("tag_window"); tagWindow != "" {
		conditions = append(conditions, "tag_window = ?")
		args = append(args, tagWindow)
	}

	if tagOutlet := c.Query("tag_outlet"); tagOutlet != "" {
		conditions = append(conditions, "tag_outlet = ?")
		args = append(args, tagOutlet)
	}

	if len(conditions) <= 0 {
		err := errors.New("query is null")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seats, err := h.svc.SearchSeats(ctx, conditions, args)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}
