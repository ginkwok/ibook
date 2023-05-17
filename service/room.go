package service

import (
	"context"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
)

func (s *svc) GetAllRooms(ctx context.Context) ([]*model.Room, error) {
	rooms, err := s.dal.GetAllRooms(dal.GetDB())
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func (s *svc) GetAvailableRooms(ctx context.Context) ([]*model.Room, error) {
	rooms, err := s.dal.GetAvailableRooms(dal.GetDB())
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return rooms, nil
}

func (s *svc) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	room, err := s.dal.CreateRoom(dal.GetDB(), room)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) DeleteRoom(ctx context.Context, id int64) error {
	err := s.DeleteSeatsOfRoom(ctx, id)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	err = s.dal.DeleteRoom(dal.GetDB(), id)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) GetRoomByID(ctx context.Context, id int64) (*model.Room, error) {
	room, err := s.dal.GetRoomByID(dal.GetDB(), id)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return room, nil
}

func (s *svc) UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	room, err := s.dal.UpdateRoom(dal.GetDB(), room)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return room, nil
}
