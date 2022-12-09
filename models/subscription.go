package models

type Subscription struct {
	ID            int `json:"id" gorm:"primary_key:auto_increment"`
	UserChannelId int `json:"user_channel_id"`
	ChannelId     int `json:"channel_id"`
}
