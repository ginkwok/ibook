package usecase

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func AdminGetAllResvsOfSeat(ctx context.Context, seatID int64) ([]*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resvs, err := service.GetResvsBySeat(ctx, seatID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func AdminCancelResv(ctx context.Context, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.CancelResv(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResvOfUser(ctx context.Context, username string) ([]*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.GetResvsOfUser(ctx, username)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func CreateResv(ctx context.Context, resv *model.Reservation) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.CreateResv(ctx, resv)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func CancelResv(ctx context.Context, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.CancelResv(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResvByID(ctx context.Context, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.GetResvByID(ctx, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func SigninResv(ctx context.Context, signinTime *time.Time, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.SigninResv(ctx, signinTime, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func SignoutResv(ctx context.Context, signinTime *time.Time, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	resv, err := service.SignoutResv(ctx, signinTime, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}
