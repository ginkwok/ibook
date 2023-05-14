package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/service"
	"github.com/ginkwok/ibook/util"
)

func GetAllRooms(ctx context.Context) ([]*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	rooms, err := service.GetAllRooms(ctx)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func GetAvailableRooms(ctx context.Context) ([]*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	rooms, err := service.GetAvailableRooms(ctx)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	room, err := service.CreateRoom(ctx, room)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func DeleteRoom(ctx context.Context, id int64) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	err := service.DeleteRoom(ctx, id)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func GetRoomByID(ctx context.Context, id int64) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	room, err := service.GetRoomByID(ctx, id)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)

	room, err := service.UpdateRoom(ctx, room)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}
