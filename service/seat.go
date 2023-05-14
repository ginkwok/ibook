package service

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func GetAllSeatsOfRoom(ctx context.Context, roomID int64) ([]*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	seats, err := dal.GetAllSeatsOfRoom(db, roomID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	seat, err := dal.CreateSeat(db, seat)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func CreateSeats(ctx context.Context, seats []*model.Seat) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	err := dal.CreateSeats(db, seats)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func DeleteSeat(ctx context.Context, seatID int64) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	err := dal.DeleteSeat(db, seatID)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func DeleteSeatsOfRoom(ctx context.Context, roomID int64) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	err := dal.DeleteSeatsOfRoom(db, roomID)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func GetSeatByID(ctx context.Context, seatID int64) (*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	seat, err := dal.GetSeatByID(db, seatID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	seat, err := dal.UpdateSeat(db, seat)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}
