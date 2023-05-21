package dal

import (
	"time"

	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

type Dal interface {
	CreateUser(db *gorm.DB, user *model.User) (*model.User, error)
	GetUserByName(db *gorm.DB, username string) (*model.User, error)
	CheckUser(db *gorm.DB, username string, password string) (bool, error)

	GetAllRooms(db *gorm.DB) ([]*model.Room, error)
	GetAvailableRooms(db *gorm.DB) ([]*model.Room, error)
	CreateRoom(db *gorm.DB, room *model.Room) (*model.Room, error)
	DeleteRoom(db *gorm.DB, id int64) error
	GetRoomByID(db *gorm.DB, id int64) (*model.Room, error)
	UpdateRoom(db *gorm.DB, room *model.Room) (*model.Room, error)

	GetAllSeatsOfRoom(db *gorm.DB, roomID int64) ([]*model.Seat, error)
	CreateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error)
	CreateSeats(db *gorm.DB, seats []*model.Seat) error
	DeleteSeat(db *gorm.DB, seatID int64) error
	DeleteSeatsOfRoom(db *gorm.DB, roomID int64) error
	GetSeatByID(db *gorm.DB, seatID int64) (*model.Seat, error)
	UpdateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error)
	SearchSeats(db *gorm.DB, conditions []string, args []interface{}) ([]*model.Seat, error)

	CreateResv(db *gorm.DB, resv *model.Reservation) (*model.Reservation, error)
	GetResvByID(db *gorm.DB, resvID int64) (*model.Reservation, error)
	GetResvsBySeat(db *gorm.DB, seatID int64) ([]*model.Reservation, error)
	GetResvsByUser(db *gorm.DB, username string) ([]*model.Reservation, error)
	UpdateResv(db *gorm.DB, resv *model.Reservation) (*model.Reservation, error)
	GetUnsignedResvsBeforeStart(db *gorm.DB, now time.Time, dur time.Duration) ([]*model.Reservation, error)
	GetUnsignedResvsAfterStart(db *gorm.DB, now time.Time, dur time.Duration) ([]*model.Reservation, error)
}

func GetDal(db *gorm.DB) Dal {
	return &dal{db: db}
}

type dal struct {
	db *gorm.DB
}
