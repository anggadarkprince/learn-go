package models

import "time"

type Wallet struct {
	ID        uint64          `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    uint64          `gorm:"column:user_id;unique"`
	Balance   float64         `gorm:"column:balance;default:0.00"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User *User `gorm:"foreignKey:user_id;references:id"` // use reference *User to prevent cyclic deps error
}

func (w *Wallet) TableName() string {
	return "wallets"
}