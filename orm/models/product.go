package models

import "time"

type Product struct {
	ID uint64 `gorm:"primaryKey;column:id;autoIncrement"`
	Name string `gorm:"column:name"`
	Price uint64 `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User `gorm:"many2many:user_like_products;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}

func (a *Product) TableName() string {
	return "products"
}