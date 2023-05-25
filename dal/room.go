package dal

import (
	"github.com/ginkwok/ibook/model"
)

func (d *dal) GetAllRooms() ([]*model.Room, error) {
	var rooms []*model.Room
	err := d.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *dal) GetAvailableRooms() ([]*model.Room, error) {
	var rooms []*model.Room
	err := d.db.Where("is_avaliable = ?", true).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *dal) CreateRoom(room *model.Room) (*model.Room, error) {
	room.ID = 0
	err := d.db.Create(&room).Error
	if err != nil {
		return nil, err
	}
	room, err = d.GetRoomByID(room.ID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (d *dal) DeleteRoom(id int64) error {
	return d.db.Delete(&model.Room{}, id).Error
}

func (d *dal) GetRoomByID(id int64) (*model.Room, error) {
	var room *model.Room
	err := d.db.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (d *dal) UpdateRoom(room *model.Room) (*model.Room, error) {
	id := room.ID
	err := d.db.Model(&room).Save(&room).Error
	if err != nil {
		return nil, err
	}
	room, err = d.GetRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}
