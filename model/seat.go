package model

type Seat struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID      int64  `json:"room_id"`
	Number      string `json:"number"`
	Location    string `json:"location"`
	QRCode      string `json:"qrcode"`
	IsAvaliable bool   `json:"is_avaliable"`
	TagWindow   bool   `json:"tag_window"`
	TagOutlet   bool   `json:"tag_outlet"`
}

func (Seat) TableName() string {
	return "seat_table"
}
