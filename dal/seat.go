package dal

import (
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func GetAllSeatsOfRoom(db *gorm.DB, roomID int64) ([]*model.Seat, error) {
	var seats []*model.Seat
	err := db.Where("room_id = ?", roomID).Find(&seats).Error
	if err != nil {
		return nil, err
	}
	return seats, nil
}

func CreateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	err := db.Create(&seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = GetSeatByID(db, seat.ID)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func CreateSeats(db *gorm.DB, seats []*model.Seat) error {
	err := db.Create(&seats).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteSeat(db *gorm.DB, seatID int64) error {
	return db.Delete(&model.Seat{}, seatID).Error
}

func DeleteSeatsOfRoom(db *gorm.DB, roomID int64) error {
	return db.Where("room_id = ?", roomID).Delete(&model.Seat{}).Error
}

func GetSeatByID(db *gorm.DB, seatID int64) (*model.Seat, error) {
	var seat *model.Seat
	err := db.First(&seat, seatID).Error
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func UpdateSeat(db *gorm.DB, seat *model.Seat) (*model.Seat, error) {
	id := seat.ID
	err := db.Model(&seat).Updates(seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = GetSeatByID(db, id)
	if err != nil {
		return nil, err
	}
	return seat, nil
}
