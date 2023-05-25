package service

import (
	"context"
	"errors"
	"time"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func (s *svc) GetResvsBySeat(ctx context.Context, seatID int64) ([]*model.Reservation, error) {
	resvs, err := s.dal.GetResvsBySeat(seatID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func (s *svc) GetResvsByUser(ctx context.Context, username string) ([]*model.Reservation, error) {

	resvs, err := s.dal.GetResvsByUser(username)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resvs, nil
}

func (s *svc) CreateResv(ctx context.Context, resv *model.Reservation) (*model.Reservation, error) {

	resv.Status = util.ResvStatusUnsignin

	resv, err := s.dal.CreateResv(resv)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) CancelResv(ctx context.Context, resvID int64) (*model.Reservation, error) {

	resv, err := s.dal.GetResvByID(resvID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusCancelled

	resv, err = s.dal.UpdateResv(resv)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) SigninResv(ctx context.Context, signinTime *time.Time, resvID int64) (*model.Reservation, error) {

	resv, err := s.dal.GetResvByID(resvID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	if resv.Status != util.ResvStatusUnsignin {
		err := errors.New("reservation status error")
		s.logger.Errorln(err)
		return nil, err
	}

	if signinTime.Before(*resv.ResvStartTime) {
		err := errors.New("reservation start time has not arrived")
		s.logger.Errorln(err)
		return nil, err
	}
	if signinTime.After(*resv.ResvEndTime) {
		err := errors.New("reservation end time has passed")
		s.logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusSignined
	resv.SigninTime = signinTime

	resv, err = s.dal.UpdateResv(resv)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) SignoutResv(ctx context.Context, signoutTime *time.Time, resvID int64) (*model.Reservation, error) {
	resv, err := s.dal.GetResvByID(resvID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	if resv.Status != util.ResvStatusSignined {
		err := errors.New("reservation status error")
		s.logger.Errorln(err)
		return nil, err
	}

	if signoutTime.After(*resv.ResvEndTime) {
		err := errors.New("reservation end time has passed")
		s.logger.Errorln(err)
		return nil, err
	}

	resv.Status = util.ResvStatusSignouted
	resv.SignoutTime = signoutTime

	resv, err = s.dal.UpdateResv(resv)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}

func (s *svc) GetResvByID(ctx context.Context, resvID int64) (*model.Reservation, error) {

	resv, err := s.dal.GetResvByID(resvID)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return resv, nil
}
