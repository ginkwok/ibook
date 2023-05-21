package handler_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	_ "github.com/ginkwok/ibook/config"
	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/handler"
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func getTestRouter(t *testing.T, mockDAL *mocks.MockDal) (*handler.HandlerStruct, *gin.Engine, string) {
	testService := service.NewService(mockDAL, nil)

	httpHandler := handler.NewHandler(testService)

	logger := util.NewLogger()
	defer logger.Sync()

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))

	token, err := httpHandler.GenerateToken(
		context.WithValue(context.Background(), util.LOGGER_KEY, logger),
		"TestUser",
	)
	assert.NoError(t, err)
	return httpHandler, router, token
}
