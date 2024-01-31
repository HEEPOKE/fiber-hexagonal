package models

import (
	"time"

	"gorm.io/gorm"
)

type CommonColumnModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type FailMessage struct {
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description error
}

type SuccessMessage struct {
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description string
	Payload     interface{}
}
