package dal

import (
	"gorm.io/gorm"

	"github.com/ginkwok/ibook/model"
)

func GetAllRooms(db *gorm.DB) ([]*model.Room, error) {
	var rooms []*model.Room
	err := db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func CreateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	err := db.Create(&room).Error
	if err != nil {
		return nil, err
	}
	room, err = GetRoomByID(db, room.ID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func DeleteRoom(db *gorm.DB, id int64) error {
	return db.Delete(&model.Room{}, id).Error
}

func GetRoomByID(db *gorm.DB, id int64) (*model.Room, error) {
	var room *model.Room
	err := db.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return room, nil
}

func UpdateRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	id := room.ID
	err := db.Model(&room).Updates(room).Error
	if err != nil {
		return nil, err
	}
	room, err = GetRoomByID(db, id)
	if err != nil {
		return nil, err
	}
	return room, nil
}
