package models

import (
	"time"
)

// User => users
// OrderDetail => order_details
// https://gorm.io/docs/models.html#Fields-Tags
type User struct {
	ID uint64 `gorm:"primaryKey;column:id;autoIncrement;<-:create"`
	Name string `gorm:"column:name"`
	Username string `gorm:"column:username"`
	Email string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Status string `gorm:"column:status;default:PENDING"` // default value
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string `gorm:"-"` // no need create / update in database (custom attribute that not available in the table)
	ReferalCode ReferalCode `gorm:"embedded"` // add embeded struct into custom field
	Wallet Wallet `gorm:"foreignKey:user_id;references:id"` // one to one relationship with Wallet
	Addresses []Address `gorm:"foreignKey:user_id;references:id"` // one to many relationship with Addresses
	LikeProducts []Product `gorm:"many2many:user_like_products;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

type ReferalCode struct {
	ID string `gorm:"column:id"`
	Username string `gorm:"column:username"`
}

// add custom table name
func (u *User) TableName() string {
	return "users"
}

type UserLog struct {
	ID int `gorm:"primaryKey;column:id;autoIncrement"`
	UserID int `gorm:"column:user_id"`
	Action string `gorm:"column:action"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
 func (ul *UserLog) TableName() string {
	return "user_logs"
 }