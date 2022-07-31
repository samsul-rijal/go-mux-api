package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `gorm:"type: varchar(255)" json:"name"`
	Email     string    `gorm:"type: varchar(255)" json:"email"`
	Password  string    `gorm:"type: varchar(255)" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
