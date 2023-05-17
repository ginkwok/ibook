package dal

import (
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func (d *dal) CreateResv(db *gorm.DB, resv *model.Reservation) (*model.Reservation, error) {
	resv.ID = 0
	err := db.Create(&resv).Error
	if err != nil {
		return nil, err
	}
	resv, err = d.GetResvByID(db, resv.ID)
	if err != nil {
		return nil, err
	}
	return resv, nil
}

func (d *dal) GetResvByID(db *gorm.DB, resvID int64) (*model.Reservation, error) {
	var resv *model.Reservation
	err := db.First(&resv, resvID).Error
	if err != nil {
		return nil, err
	}
	return resv, nil
}

func (d *dal) GetResvsBySeat(db *gorm.DB, seatID int64) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := db.Where("seat_id = ?", seatID).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}

func (d *dal) GetResvsByUser(db *gorm.DB, username string) ([]*model.Reservation, error) {
	var resvs []*model.Reservation
	err := db.Where("username = ?", username).Find(&resvs).Error
	if err != nil {
		return nil, err
	}
	return resvs, nil
}

func (d *dal) UpdateResv(db *gorm.DB, resv *model.Reservation) (*model.Reservation, error) {
	id := resv.ID
	err := db.Model(&resv).Updates(resv).Error
	if err != nil {
		return nil, err
	}
	resv, err = d.GetResvByID(db, id)
	if err != nil {
		return nil, err
	}
	return resv, nil
}
