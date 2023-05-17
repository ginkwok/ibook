package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
)

type Service interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CheckUser(ctx context.Context, username string, password string) (bool, error)

	GetAllRooms(ctx context.Context) ([]*model.Room, error)
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, id int64) error
	GetRoomByID(ctx context.Context, id int64) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error)

	GetAllSeatsOfRoom(ctx context.Context, roomID int64) ([]*model.Seat, error)
	CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	CreateSeats(ctx context.Context, seats []*model.Seat) error
	DeleteSeat(ctx context.Context, seatID int64) error
	DeleteSeatsOfRoom(ctx context.Context, roomID int64) error
	GetSeatByID(ctx context.Context, seatID int64) (*model.Seat, error)
	UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error)
	SearchSeats(ctx context.Context) ([]*model.Seat, error)

	GetResvsBySeat(ctx context.Context, seatID int64) ([]*model.Reservation, error)
	GetResvsByUser(ctx context.Context, username string) ([]*model.Reservation, error)
	CreateResv(ctx context.Context, resv *model.Reservation) (*model.Reservation, error)
	CancelResv(ctx context.Context, resvID int64) (*model.Reservation, error)
	SigninResv(ctx context.Context, signinTime *time.Time, resvID int64) (*model.Reservation, error)
	SignoutResv(ctx context.Context, signoutTime *time.Time, resvID int64) (*model.Reservation, error)
	GetResvByID(ctx context.Context, resvID int64) (*model.Reservation, error)
}

func NewService(dalClient dal.Dal, logger *zap.SugaredLogger) Service {
	return &svc{
		dal:    dalClient,
		logger: logger,
	}
}

type svc struct {
	dal    dal.Dal
	logger *zap.SugaredLogger
}
