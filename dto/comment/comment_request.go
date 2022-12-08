package commentdto

type CommentRequest struct {
	Title     string `json:"title" gorm:"type: varchar(255)"`
	Comment   string `gorm:"type: varchar(255)" json:"comment"`
	ChannelID int    `json:"" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EditCommentRequest struct {
	Title   string `json:"title" gorm:"type: varchar(255)"`
	Comment string `gorm:"type: varchar(255)" json:"comment"`
}
