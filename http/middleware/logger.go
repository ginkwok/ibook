package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ginkwok/ibook/util"
)

func LoggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			logger.Infoln("received request",
				"method: "+c.Request.Method+";",
				"url: "+c.Request.URL.String()+";",
				"user-agent: "+c.Request.UserAgent()+";",
				"ip: "+c.ClientIP()+";",
				"sent response status:", c.Writer.Status(),
			)
			logger.Sync()
		}()

		ctx := context.WithValue(c.Request.Context(), util.LOGGER_KEY, logger)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
