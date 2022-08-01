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

type UserResponseWithProduct struct {
	ID       int               `json:"id"`
	Name     string            `gorm:"type: varchar(255)" json:"name"`
	Email    string            `gorm:"type: varchar(255)" json:"email"`
	Profile  ProfileResponse   `json:"profile" gorm:"-"`
	Products []ProductResponse `json:"products" gorm:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}
func (UserResponseWithProduct) TableName() string {
	return "users"
}
