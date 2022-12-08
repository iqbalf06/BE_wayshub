package videodto

import "time"

type VideoResponse struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `gorm:"type: varchar(255)" json:"thumbnail"`
	Description string `gorm:"type: varchar(255)" json:"description"`
	Video       string `gorm:"type: varchar(255)" json:"video"`

	ViewCount int       `json:"viewcount" form:"viewcount" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
