package model

import "time"

type Reservation struct {
	ID            int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string     `json:"username"`
	RoomID        int64      `json:"room_id"`
	SeatID        int64      `json:"seat_id"`
	CreateTime    *time.Time `json:"create_time"`
	ResvStartTime *time.Time `json:"reservasion_start_time"`
	ResvEndTime   *time.Time `json:"reservasion_end_time"`
	SigninTime    *time.Time `json:"signin_time"`
	SignoutTime   *time.Time `json:"signout_time"`
	Status        string     `json:"status"`
}

func (Reservation) TableName() string {
	return "reservation_table"
}
