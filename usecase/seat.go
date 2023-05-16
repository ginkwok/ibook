package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func GetAllSeatsOfRoom(ctx context.Context, roomID int64) ([]*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	seats, err := service.GetAllSeatsOfRoom(ctx, roomID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func CreateSeats(ctx context.Context, seats []*model.Seat) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	err := service.CreateSeats(ctx, seats)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func DeleteSeat(ctx context.Context, seatID int64) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	err := service.DeleteSeat(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func GetSeatByID(ctx context.Context, seatID int64) (*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	seat, err := service.GetSeatByID(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	seat, err := service.UpdateSeat(ctx, seat)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func SearchSeats(ctx context.Context) ([]*model.Seat, error) {
	return nil, nil
}
