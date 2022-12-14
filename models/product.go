package models

type Product struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	Name       string       `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc       string       `json:"desc" gorm:"type:text" form:"desc"`
	Price      int          `json:"price" form:"price" gorm:"type: int"`
	Image      string       `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty        int          `json:"qty" form:"qty"`
	UserID     int          `json:"user_id" form:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User       UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category   []Category   `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int        `json:"category_id" form:"category_id" gorm:"-"`
}

type ProductResponse struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Desc       string       `json:"desc"`
	Price      int          `json:"price"`
	Image      string       `json:"image"`
	Qty        int          `json:"qty"`
	UserID     int          `json:"-"`
	User       UserResponse `json:"user"`
	Category   []Category   `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int        `json:"category_id" form:"category_id" gorm:"-"`
}

type ProductResponseWithCategory struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Desc       string       `json:"desc"`
	Price      int          `json:"price"`
	Image      string       `json:"image"`
	Qty        int          `json:"qty"`
	UserID     int          `json:"-"`
	User       UserResponse `json:"user"`
	Category   []Category   `json:"category" gorm:"many2many:product_categories;ForeignKey:ID;joinForeignKey:ProductID;References:ID;joinReferences:CategoryID"`
	CategoryID []int        `json:"-" form:"category_id" gorm:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductResponseWithCategory) TableName() string {
	return "products"
}
