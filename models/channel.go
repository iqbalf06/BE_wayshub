package models

import "time"

type Channel struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Channelname string    `json:"channelName" gorm:"type: varchar(255)"`
	Email       string    `json:"email" gorm:"type: varchar(255)"`
	Photo       string    `json:"photo" gorm:"type: varchar(255)"`
	Cover       string    `json:"cover" gorm:"type: varchar(255)"`
	Password    string    `json:"-" gorm:"type: varchar(255)"`
	Description string    `gorm:"type: varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ChannelResponse struct {
	ID          int    `json:"id"`
	Channelname string `json:"channelName"`
	Email       string `json:"email"`
	Photo       string `json:"photo"`
	Cover       string    `json:"cover"`
	Description string    `json:"description"`


}

func (ChannelResponse) TableName() string {
	return "channels"
}
