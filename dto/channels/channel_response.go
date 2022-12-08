package channelsdto

type ChannelResponse struct {
  ID       int    `json:"id"`
  Channelname     string `json:"channelName" form:"fullname" validate:"required"`
  Email    string `json:"email" form:"email" validate:"required"`
  Photo string `json:"photo" form:"photo"`
  Description string `gorm:"type: varchar(255)" json:"description"`
  Cover string `json:"cover" form:"cover"`

}

type DeleteResponse struct {
	ID int `json:"id"`
}