package dal

import (
	"time"

	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

type Dal interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByName(username string) (*model.User, error)
	CheckUser(username string, password string) (bool, error)

	GetAllRooms() ([]*model.Room, error)
	GetAvailableRooms() ([]*model.Room, error)
	CreateRoom(room *model.Room) (*model.Room, error)
	DeleteRoom(id int64) error
	GetRoomByID(id int64) (*model.Room, error)
	UpdateRoom(room *model.Room) (*model.Room, error)

	GetAllSeatsOfRoom(roomID int64) ([]*model.Seat, error)
	CreateSeat(seat *model.Seat) (*model.Seat, error)
	CreateSeats(seats []*model.Seat) error
	DeleteSeat(seatID int64) error
	DeleteSeatsOfRoom(roomID int64) error
	GetSeatByID(seatID int64) (*model.Seat, error)
	UpdateSeat(seat *model.Seat) (*model.Seat, error)
	SearchSeats(conditions []string, args []interface{}) ([]*model.Seat, error)

	CreateResv(resv *model.Reservation) (*model.Reservation, error)
	GetResvByID(resvID int64) (*model.Reservation, error)
	GetResvsBySeat(seatID int64) ([]*model.Reservation, error)
	GetResvsByUser(username string) ([]*model.Reservation, error)
	UpdateResv(resv *model.Reservation) (*model.Reservation, error)
	GetUnsignedResvsBeforeStart(now time.Time, dur time.Duration) ([]*model.Reservation, error)
	GetUnsignedResvsAfterStart(now time.Time, dur time.Duration) ([]*model.Reservation, error)
}

func GetDal(db *gorm.DB) Dal {
	return &dal{db: db}
}

type dal struct {
	db *gorm.DB
}
