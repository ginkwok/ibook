package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
	"go.uber.org/zap"
)

func (h *HandlerStruct) AdminGetResvsBySeatHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get(util.JWT_USERNAME)
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

	resvs, err := h.svc.GetResvsBySeat(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resvs)
}

func (h *HandlerStruct) AdminCancelResvHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	_, ok := c.Get(util.JWT_USERNAME)
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
	_, err = strconv.ParseInt(seatIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resvID, err := strconv.ParseInt(resvIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := h.svc.CancelResv(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resv)
}

func (h *HandlerStruct) GetResvsByUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	username, ok := c.Get(util.JWT_USERNAME)
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resvs, err := h.svc.GetResvsByUser(ctx, username.(string))
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resvs)
}

func (h *HandlerStruct) CreateResvHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	username, ok := c.Get(util.JWT_USERNAME)
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var resv *model.Reservation
	if err := c.ShouldBindJSON(&resv); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resv.RoomID <= 0 {
		err := errors.New("room id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if resv.SeatID <= 0 {
		err := errors.New("seat id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nowtime := time.Now().In(loc)
	resv.CreateTime = &nowtime

	// if !resv.CreateTime.Before(resv.ResvStartTime) || !resv.ResvStartTime.Before(resv.ResvEndTime) {
	// 	err := errors.New("reservation time is illegal")
	// 	logger.Errorln(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if !resv.ResvStartTime.Before(*resv.ResvEndTime) {
		err := errors.New("reservation time is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv.ID = 0
	resv.Username = username.(string)
	resv.Status = util.ResvStatusUnsignin

	resv, err = h.svc.CreateResv(ctx, resv)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resv)
}

func (h *HandlerStruct) CancelResvHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	username, ok := c.Get(util.JWT_USERNAME)
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resvID, err := strconv.ParseInt(resvIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := h.svc.GetResvByID(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resv.Username != username.(string) {
		err := errors.New("reservation not belong to current user")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resv, err = h.svc.CancelResv(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resv)
}

func (h *HandlerStruct) SigninResvHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	username, ok := c.Get(util.JWT_USERNAME)
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resvID, err := strconv.ParseInt(resvIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := h.svc.GetResvByID(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resv.Username != username.(string) {
		err := errors.New("reservation not belong to current user")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nowtime := time.Now().In(loc)

	resv, err = h.svc.SigninResv(ctx, &nowtime, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resv)
}

func (h *HandlerStruct) SignoutResvHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	username, ok := c.Get(util.JWT_USERNAME)
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resvIDStr := c.Param("resv_id")
	if resvIDStr == "" {
		err := errors.New("id is illegal")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resvID, err := strconv.ParseInt(resvIDStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resv, err := h.svc.GetResvByID(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resv.Username != username.(string) {
		err := errors.New("reservation not belong to current user")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nowtime := time.Now().In(loc)

	resv, err = h.svc.SignoutResv(ctx, &nowtime, resvID)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resv)
}
