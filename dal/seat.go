package dal

import (
	"strings"

	"github.com/ginkwok/ibook/model"
)

func (d *dal) GetAllSeatsOfRoom(roomID int64) ([]*model.Seat, error) {
	var seats []*model.Seat
	err := d.db.Where("room_id = ?", roomID).Find(&seats).Error
	if err != nil {
		return nil, err
	}
	return seats, nil
}

func (d *dal) CreateSeat(seat *model.Seat) (*model.Seat, error) {
	seat.ID = 0
	err := d.db.Create(&seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = d.GetSeatByID(seat.ID)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) CreateSeats(seats []*model.Seat) error {
	err := d.db.Create(&seats).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *dal) DeleteSeat(seatID int64) error {
	return d.db.Delete(&model.Seat{}, seatID).Error
}

func (d *dal) DeleteSeatsOfRoom(roomID int64) error {
	return d.db.Where("room_id = ?", roomID).Delete(&model.Seat{}).Error
}

func (d *dal) GetSeatByID(seatID int64) (*model.Seat, error) {
	var seat *model.Seat
	err := d.db.First(&seat, seatID).Error
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) UpdateSeat(seat *model.Seat) (*model.Seat, error) {
	id := seat.ID
	err := d.db.Model(&seat).Save(seat).Error
	if err != nil {
		return nil, err
	}
	seat, err = d.GetSeatByID(id)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (d *dal) SearchSeats(conditions []string, args []interface{}) ([]*model.Seat, error) {
	queryStr := "SELECT * FROM seats"
	if len(conditions) > 0 {
		queryStr += " WHERE " + strings.Join(conditions, " AND ")
	}

	var seats []*model.Seat
	err := d.db.Raw(queryStr, args...).Scan(&seats).Error
	if err != nil {
		return nil, err
	}
	return seats, nil
}
