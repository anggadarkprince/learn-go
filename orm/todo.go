package orm

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int    `gorm:"column:user_id"`
	Title     string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (t *Todo) TableName() string {
	return "todos"
}

type TodoLog struct {
	gorm.Model // Model includes ID, CreatedAt, UpdatedAt, DeletedAt
	UserID    int    `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
}