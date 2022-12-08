package authdto

type RegisterResponse struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Channelname  string `gorm:"type: varchar(255)" json:"channelName" validate:"required"`
}

type LoginResponse struct {
	Email string `gorm:"type: varchar(255)" json:"email"`
	Channelname  string `gorm:"type: varchar(255)" json:"channelName"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	ID int `gorm:"type: int" json:"id"`
	Channelname string `gorm:"type:varchar(255)" json:"channelName"`
	Email string `gorm:"type: varchar(255)" json:"email"`
}