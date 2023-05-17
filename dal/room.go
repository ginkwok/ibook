package dal

import (
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func (d *dal) GetAllRooms(db *gorm.DB) ([]*model.Room, error) {
	var rooms []*model.Room
	err := db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *dal) GetAvailableRooms(db *gorm.DB) ([]*model.Room, error) {
	var rooms []*model.Room
	err := db.Where("is_avaliable = ?", true).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *dal) CreateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	room.ID = 0
	err := db.Create(&room).Error
	if err != nil {
		return nil, err
	}
	room, err = d.GetRoomByID(db, room.ID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (d *dal) DeleteRoom(db *gorm.DB, id int64) error {
	return db.Delete(&model.Room{}, id).Error
}

func (d *dal) GetRoomByID(db *gorm.DB, id int64) (*model.Room, error) {
	var room *model.Room
	err := db.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (d *dal) UpdateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	id := room.ID
	err := db.Model(&room).Save(&room).Error
	if err != nil {
		return nil, err
	}
	room, err = d.GetRoomByID(db, id)
	if err != nil {
		return nil, err
	}
	return room, nil
}
