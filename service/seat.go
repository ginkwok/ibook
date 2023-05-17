package service

import (
	"context"

	"github.com/ginkwok/ibook/dal"
	"github.com/ginkwok/ibook/model"
)

func (s *svc) GetAllSeatsOfRoom(ctx context.Context, roomID int64) ([]*model.Seat, error) {
	seats, err := s.dal.GetAllSeatsOfRoom(dal.GetDB(), roomID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return seats, nil
}

func (s *svc) CreateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.CreateSeat(dal.GetDB(), seat)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) CreateSeats(ctx context.Context, seats []*model.Seat) error {
	err := s.dal.CreateSeats(dal.GetDB(), seats)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) DeleteSeat(ctx context.Context, seatID int64) error {
	err := s.dal.DeleteSeat(dal.GetDB(), seatID)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) DeleteSeatsOfRoom(ctx context.Context, roomID int64) error {
	err := s.dal.DeleteSeatsOfRoom(dal.GetDB(), roomID)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) GetSeatByID(ctx context.Context, seatID int64) (*model.Seat, error) {
	seat, err := s.dal.GetSeatByID(dal.GetDB(), seatID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) UpdateSeat(ctx context.Context, seat *model.Seat) (*model.Seat, error) {
	seat, err := s.dal.UpdateSeat(dal.GetDB(), seat)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return seat, nil
}

func (s *svc) SearchSeats(ctx context.Context) ([]*model.Seat, error) {
	return nil, nil
}
