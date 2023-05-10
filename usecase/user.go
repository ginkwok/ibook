package usecase

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	user, err := service.CreateUser(ctx, user)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return user, nil
}

func CheckUser(ctx context.Context, username, password string) bool {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	ok, err := service.CheckUser(ctx, username, password)
	if err != nil {
		logger.Errorln(err)
		return false
	}
	if !ok {
		err = errors.New("unexpected error")
		logger.Errorln(err)
		return false
	}
	return true
}
