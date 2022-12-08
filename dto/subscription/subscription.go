package subscriptiondto

type Subscriber struct {
	Subscribe int `json:"subscribe"`
	ChannelID int `json:"" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
