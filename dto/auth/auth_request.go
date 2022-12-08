package authdto

type RegisterRequest struct {
	Channelname string `gorm:"type: varchar(255)" json:"channelName" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Description string `gorm:"type: varchar(255)" json:"description"`
	Image    string `gorm:"type: varchar(255)" json:"image"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
