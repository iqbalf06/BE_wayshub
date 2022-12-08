package models

type Subscription struct {
	ID         int     `json:"id" gorm:"primary_key:auto_increment"`
	ChannelID  int     `json:"channel_id"`
	Subscriber int     `json:"subscriber"`
	Channel    Channel `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
