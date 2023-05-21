package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func (h *HandlerStruct) RegisterHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.svc.CreateUser(ctx, &user); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.GenerateToken(ctx, user.Username)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created", "token": token})
}

func (h *HandlerStruct) LoginHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" {
		err := errors.New("username is null")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Password == "" {
		err := errors.New("password is null")
		logger.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.svc.CheckUser(ctx, user.Username, user.Password)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		err := errors.New("invalid credentials")
		logger.Errorln(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.GenerateToken(ctx, user.Username)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *HandlerStruct) GenerateToken(ctx context.Context, username string) (string, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.key")))
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	return tokenString, nil
}
