package service

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func GetResvsBySeat(ctx context.Context, seatID int64) ([]*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resvs, err := dal.GetResvsBySeat(db, seatID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func GetResvsOfUser(ctx context.Context, username string) ([]*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resvs, err := dal.GetResvsByUser(db, username)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func CreateResv(ctx context.Context, resv *model.Reservation) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resv.Status = util.ResvStatusUnsignin

	resv, err := dal.CreateResv(db, resv)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func CancelResv(ctx context.Context, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resv, err := dal.GetResvByID(db, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusCancelled

	resv, err = dal.UpdateResv(db, resv)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func SigninResv(ctx context.Context, signinTime *time.Time, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resv, err := dal.GetResvByID(db, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	if resv.Status != util.ResvStatusUnsignin {
		err := errors.New("reservation status error")
		logger.Errorln(err)
		return nil, err
	}

	if signinTime.Before(*resv.ResvStartTime) {
		err := errors.New("reservation start time has not arrived")
		logger.Errorln(err)
		return nil, err
	}
	if signinTime.After(*resv.ResvEndTime) {
		err := errors.New("reservation end time has passed")
		logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusSignined
	resv.SigninTime = signinTime

	resv, err = dal.UpdateResv(db, resv)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func SignoutResv(ctx context.Context, signoutTime *time.Time, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resv, err := dal.GetResvByID(db, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	if resv.Status != util.ResvStatusSignined {
		err := errors.New("reservation status error")
		logger.Errorln(err)
		return nil, err
	}

	if signoutTime.After(*resv.ResvEndTime) {
		err := errors.New("reservation end time has passed")
		logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusSignouted
	resv.SignoutTime = signoutTime

	resv, err = dal.UpdateResv(db, resv)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func GetResvByID(ctx context.Context, resvID int64) (*model.Reservation, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	resv, err := dal.GetResvByID(db, resvID)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}
