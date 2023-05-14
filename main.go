package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/ginkwok/ibook/config"
	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/http/handler"
	"github.com/ginkwok/ibook/http/middleware"
	"github.com/ginkwok/ibook/util"
)

func main() {
	logger := util.NewLogger()
	defer logger.Sync()

	db := dal.GetDB(
		viper.GetString("db.mysql.host"),
		viper.GetString("db.mysql.port"),
		viper.GetString("db.mysql.username"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.database"),
	)

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(middleware.MySQLMiddleware(db))

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", handler.RegisterHandler)
		v1.POST("/login", handler.LoginHandler)

		v1.GET("/admin/rooms", middleware.AuthMiddleware(), handler.AdminGetAllRoomsHandler)
		v1.POST("/admin/rooms", middleware.AuthMiddleware(), handler.AdminCreateRoomHandler)
		v1.DELETE("/admin/rooms/:room_id", middleware.AuthMiddleware(), handler.AdminDeleteRoomHandler)
		v1.GET("/admin/rooms/:room_id", middleware.AuthMiddleware(), handler.AdminGetRoomByIDHandler)
		v1.PATCH("/admin/rooms/:room_id", middleware.AuthMiddleware(), handler.AdminUpdateRoomHandler)

		v1.GET("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), handler.AdminGetAllSeatsOfRoomHandler)
		v1.POST("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), handler.AdminCreateSeatsHandler)
		v1.DELETE("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), handler.AdminDeleteSeatHandler)
		v1.GET("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), handler.AdminGetSeatByIDHandler)
		v1.PATCH("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), handler.AdminUpdateSeatHandler)

		v1.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Hello, %v", username)})
		})
	}
	router.Run(":" + viper.GetString("server.port"))
}
