package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/util"
)

func MySQLMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), util.MYSQL_KEY, db)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
