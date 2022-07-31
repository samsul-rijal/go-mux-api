package models

type Product struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text" form:"desc"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int    `json:"qty" form:"qty"`
	UserID int    `json:"user_id" form:"user_id"`
	User   User   `json:"user"`
}
