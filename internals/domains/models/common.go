package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonColumnModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
