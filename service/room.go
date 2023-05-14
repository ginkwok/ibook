package service

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func GetAllRooms(ctx context.Context) ([]*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	rooms, err := dal.GetAllRooms(db)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	room, err := dal.CreateRoom(db, room)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func DeleteRoom(ctx context.Context, id int64) error {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	err := DeleteSeatsOfRoom(ctx, id)
	if err != nil {
		logger.Errorln(err)
		return err
	}

	err = dal.DeleteRoom(db, id)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}

func GetRoomByID(ctx context.Context, id int64) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	room, err := dal.GetRoomByID(db, id)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	logger := ctx.Value(util.LOGGER_KEY).(*zap.SugaredLogger)
	db := ctx.Value(util.MYSQL_KEY).(*gorm.DB)

	room, err := dal.UpdateRoom(db, room)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return room, nil
}
