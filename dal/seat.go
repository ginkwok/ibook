package dal

import (
	"strings"

	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func (d *dal) GetAllSeatsOfRoom(db *gorm.DB, roomID int64) ([]*model.Seat, error) {
	var seats []*model.Seat
	err := db.Where("room_id = ?", roomID).Find(&seats).Error
	if err != nil {
		return nil, err
	}
	return seats, nil
}

func (d *dal) CreateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	seat.ID = 0
	err := db.Create(&seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = d.GetSeatByID(db, seat.ID)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) CreateSeats(db *gorm.DB, seats []*model.Seat) error {
	err := db.Create(&seats).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *dal) DeleteSeat(db *gorm.DB, seatID int64) error {
	return db.Delete(&model.Seat{}, seatID).Error
}

func (d *dal) DeleteSeatsOfRoom(db *gorm.DB, roomID int64) error {
	return db.Where("room_id = ?", roomID).Delete(&model.Seat{}).Error
}

func (d *dal) GetSeatByID(db *gorm.DB, seatID int64) (*model.Seat, error) {
	var seat *model.Seat
	err := db.First(&seat, seatID).Error
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) UpdateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	id := seat.ID
	err := db.Model(&seat).Save(seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = d.GetSeatByID(db, id)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) SearchSeats(db *gorm.DB, conditions []string, args []interface{}) ([]*model.Seat, error) {
	queryStr := "SELECT * FROM seats"
	if len(conditions) > 0 {
		queryStr += " WHERE " + strings.Join(conditions, " AND ")
	}

	var seats []*model.Seat
	err := db.Raw(queryStr, args...).Scan(&seats).Error
	if err != nil {
		return nil, err
	}
	return seats, nil
}
