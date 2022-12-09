package subscriptiondto

type Subscriber struct {
	UserChannelId int `json:"" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ChannelID     int `json:"" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
