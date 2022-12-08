package videodto

type VideoRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `gorm:"type: varchar(255)" json:"thumbnail"`
	Description string `gorm:"type: varchar(255)" json:"description"`
	Video       string `gorm:"type: varchar(255)" json:"video"`
}

type EditVideoRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Thumbnail   string `gorm:"type: varchar(255)" json:"thumbnail"`
	Description string `gorm:"type: varchar(255)" json:"description"`
	Video       string `gorm:"type: varchar(255)" json:"video"`
}
