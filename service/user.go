package service

import (
	"context"
	"errors"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	user, err := dal.CreateUser(db, user)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return user, nil
}

func CheckUser(ctx context.Context, username string, password string) (bool, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	ok, err := dal.CheckUser(db, username, password)
	if err != nil {
		logger.Errorln(err)
		return false, err
	}
	if !ok {
		err = errors.New("password error")
		logger.Errorln(err)
		return false, err
	}
	return true, nil
}
