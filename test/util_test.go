package test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	_ "github.com/ginkwok/ibook/config"
	"github.com/ginkwok/ibook/dal/mocks"
	"github.com/ginkwok/ibook/http/handler"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func getTestRouter(t *testing.T, mockDAL *mocks.MockDal) (*handler.HandlerStruct, *gin.Engine, string) {
	logger := zap.NewNop().Sugar()
	defer logger.Sync()

	testService := service.NewService(mockDAL, logger)

	httpHandler := handler.NewHandler(testService)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.Use(loggerMiddleware(logger))

	token, err := httpHandler.GenerateToken(
		context.WithValue(context.Background(), util.LOGGER_KEY, logger),
		"TestUser1",
	)
	assert.NoError(t, err)
	return httpHandler, router, token
}

func loggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.LOGGER_KEY, logger)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
