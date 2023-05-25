package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func (s *svc) NoticeMonitoring(ctx context.Context) {
	for {
		go s.noticeMonitoring(ctx)
		time.Sleep(15 * time.Minute)
	}
}

func (s *svc) noticeMonitoring(ctx context.Context) {
	now := time.Now()
	go s.noticeMonitoringBeforeStart(ctx, now)
	go s.noticeMonitoringAfterStart(ctx, now)
}

func (s *svc) noticeMonitoringBeforeStart(ctx context.Context, now time.Time) {
	resvs, err := s.dal.GetUnsignedResvsBeforeStart(now, time.Minute*15)
	if err != nil {
		s.logger.Errorln(err)
		return
	}
	for _, resv := range resvs {
		err := s.beforeStartNotice(resv)
		if err != nil {
			s.logger.Errorln(err)
			continue
		}
	}
}

func (s *svc) noticeMonitoringAfterStart(ctx context.Context, now time.Time) {
	resvs, err := s.dal.GetUnsignedResvsAfterStart(now, time.Minute*15)
	if err != nil {
		s.logger.Errorln(err)
		return
	}
	for _, resv := range resvs {
		if now.Sub(*resv.ResvStartTime) <= time.Minute*30 {
			err := s.afterStartNotice(resv)
			if err != nil {
				s.logger.Errorln(err)
				continue
			}
		} else {
			err := s.cancelAndNotice(ctx, resv)
			if err != nil {
				s.logger.Errorln(err)
				continue
			}
		}
	}

}

func (s *svc) beforeStartNotice(resv *model.Reservation) error {
	user, err := s.dal.GetUserByName(resv.Username)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	err = notice(user.NoticeURL, util.NoticeBeforeStart)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) afterStartNotice(resv *model.Reservation) error {
	user, err := s.dal.GetUserByName(resv.Username)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	err = notice(user.NoticeURL, util.NoticeAfterStart)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func (s *svc) cancelAndNotice(ctx context.Context, resv *model.Reservation) error {
	resv, err := s.CancelResv(ctx, resv.ID)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	user, err := s.dal.GetUserByName(resv.Username)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}

	err = notice(user.NoticeURL, util.NoticeCancel)
	if err != nil {
		s.logger.Errorln(err)
		return err
	}
	return nil
}

func notice(url string, msg string) error {
	type Msg struct {
		Msg string `json:"msg"`
	}
	data, err := json.Marshal(&Msg{Msg: msg})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
