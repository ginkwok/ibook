package service

import (
	"context"
	"errors"

	"github.com/ginkwok/ibook/model"
)

func (s *svc) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := s.dal.CreateUser(user)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}
	return user, nil
}

func (s *svc) CheckUser(ctx context.Context, username string, password string) (bool, error) {
	ok, err := s.dal.CheckUser(username, password)
	if err != nil {
		s.logger.Errorln(err)
		return false, err
	}
	if !ok {
		err = errors.New("password error")
		s.logger.Errorln(err)
		return false, err
	}
	return true, nil
}
