package model

type Room struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Capacity    int64  `json:"capacity"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
	Location    string `json:"location"`
	Description string `json:"description"`
	IsAvaliable bool   `json:"is_avaliable"`
}

func (Room) TableName() string {
	return "study_room_table"
}
