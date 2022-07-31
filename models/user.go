package models

import "time"

type User struct {
	ID        int               `json:"id"`
	Name      string            `gorm:"type: varchar(255)" json:"name"`
	Email     string            `gorm:"type: varchar(255)" json:"email"`
	Password  string            `gorm:"type: varchar(255)" json:"password"`
	Profile   ProfileResponse   `json:"profile"`
	Products  []ProductResponse `json:"products"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
