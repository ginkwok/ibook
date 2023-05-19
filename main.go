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
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func main() {
	logger := util.NewLogger()
	defer logger.Sync()

	db := dal.GetDB()
	dalClient := dal.GetDal(db)

	svc := service.NewService(dalClient, logger)

	httpHandler := handler.NewHandler(svc)

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(middleware.MySQLMiddleware(db))

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", httpHandler.RegisterHandler)
		v1.POST("/login", httpHandler.LoginHandler)

		v1.GET("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminGetAllRoomsHandler)
		v1.POST("/admin/rooms", middleware.AuthMiddleware(), httpHandler.AdminCreateRoomHandler)
		v1.DELETE("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteRoomHandler)
		v1.GET("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminGetRoomByIDHandler)
		v1.PATCH("/admin/rooms/:room_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateRoomHandler)

		v1.GET("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminGetAllSeatsOfRoomHandler)
		v1.POST("/admin/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.AdminCreateSeatsHandler)
		v1.DELETE("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminDeleteSeatHandler)
		v1.GET("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminGetSeatByIDHandler)
		v1.PATCH("/admin/rooms/:room_id/seats/:seat_id", middleware.AuthMiddleware(), httpHandler.AdminUpdateSeatHandler)

		v1.GET("/admin/rooms/:room_id/seats/:seat_id/reservations", middleware.AuthMiddleware(), httpHandler.AdminGetResvsBySeatHandler)
		v1.PATCH("/admin/rooms/:room_id/seats/:seat_id/reservations/:resv_id", middleware.AuthMiddleware(), httpHandler.AdminCancelResvHandler)

		v1.GET("/rooms", middleware.AuthMiddleware(), httpHandler.GetAllRoomsHandler)
		v1.GET("/rooms/:room_id/seats", middleware.AuthMiddleware(), httpHandler.GetAllSeatsOfRoomHandler)

		v1.GET("/reservations", middleware.AuthMiddleware(), httpHandler.GetResvsByUserHandler)
		v1.POST("/reservations", middleware.AuthMiddleware(), httpHandler.CreateResvHandler)
		v1.PATCH("/reservations/:resv_id/cancel", middleware.AuthMiddleware(), httpHandler.CancelResvHandler)
		v1.PATCH("/reservations/:resv_id/signin", middleware.AuthMiddleware(), httpHandler.SigninResvHandler)
		v1.PATCH("/reservations/:resv_id/signout", middleware.AuthMiddleware(), httpHandler.SignoutResvHandler)

		v1.GET("/search", middleware.AuthMiddleware(), httpHandler.SearchSeatsHandler)

		v1.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Hello, %v", username)})
		})
	}
	router.Run(":" + viper.GetString("server.port"))
}
