package models

import "time"

type Profile struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string       `json:"phone" gorm:"type: varchar(255)"`
	Gender    string       `json:"gender" gorm:"type: varchar(255)"`
	Address   string       `json:"address" gorm:"type: text"`
	UserID    int          `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type ProfileResponse struct {
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
