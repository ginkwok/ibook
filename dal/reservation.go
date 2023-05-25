package dal

import (
	"time"

	"github.com/ginkwok/ibook/model"
	"github.com/ginkwok/ibook/util"
)

func (d *dal) CreateResv(resv *model.Reservation) (*model.Reservation, error) {
	resv.ID = 0
	err := d.db.Create(&resv).Error
	if err != nil {
		return nil, err
	}
	resv, err = d.GetResvByID(resv.ID)
	if err != nil {
		return nil, err
	}
	return resv, nil
}

func (d *dal) GetResvByID(resvID int64) (*model.Reservation, error) {
	var resv *model.Reservation
	err := d.db.First(&resv, resvID).Error
	if err != nil {
		return nil, err
	}
	return resv, nil
}

func (d *dal) GetResvsBySeat(seatID int64) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := d.db.Where("seat_id = ?", seatID).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}

func (d *dal) GetResvsByUser(username string) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := d.db.Where("username = ?", username).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}

func (d *dal) UpdateResv(resv *model.Reservation) (*model.Reservation, error) {
	id := resv.ID
	err := d.db.Model(&resv).Updates(resv).Error
	if err != nil {
		return nil, err
	}
	resv, err = d.GetResvByID(id)
	if err != nil {
		return nil, err
	}
	return resv, nil
}

func (d *dal) GetUnsignedResvsBeforeStart(now time.Time, dur time.Duration) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := d.db.Where("status = ? AND resv_start_time > ? AND resv_start_time - ? <= ?", util.ResvStatusUnsignin, now, dur, now).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}

func (d *dal) GetUnsignedResvsAfterStart(now time.Time, dur time.Duration) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := d.db.Where("status = ? AND ? > resv_start_time AND ? - resv_start_time >= ?", util.ResvStatusUnsignin, now, now, dur).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}
