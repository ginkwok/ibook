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

		v1.GET("/admin/rooms/:room_id/seats/:seat_id/reservations", middleware.AuthMiddleware(), handler.AdminGetAllResvsOfSeatHandler)
		v1.PATCH("/admin/rooms/:room_id/seats/:seat_id/reservations/:resv_id", middleware.AuthMiddleware(), handler.AdminCancelResvHandler)

		v1.GET("/rooms", middleware.AuthMiddleware(), handler.GetAllRoomsHandler)
		v1.GET("/rooms/:room_id/seats", middleware.AuthMiddleware(), handler.GetAllSeatsOfRoomHandler)

		v1.GET("/reservations", middleware.AuthMiddleware(), handler.GetResvOfUserHandler)
		v1.POST("/reservations", middleware.AuthMiddleware(), handler.CreateResvHandler)
		v1.PATCH("/reservations/:resv_id/cancel", middleware.AuthMiddleware(), handler.CancelResvHandler)
		v1.PATCH("/reservations/:resv_id/signin", middleware.AuthMiddleware(), handler.SigninResvHandler)
		v1.PATCH("/reservations/:resv_id/signout", middleware.AuthMiddleware(), handler.SignoutResvHandler)

		v1.GET("/search", middleware.AuthMiddleware(), handler.SearchSeatsHandler)

		v1.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Hello, %v", username)})
		})
	}
	router.Run(":" + viper.GetString("server.port"))
}
