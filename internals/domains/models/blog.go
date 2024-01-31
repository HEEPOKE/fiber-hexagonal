package models

type BlogModel struct {
	CommonColumnModel
	AccountID   uint    `gorm:"column:account_id;type:int" json:"account_id"`
	Name        string  `gorm:"column:name;type:VARCHAR(255);unique;not null" json:"name" validate:"required"`
	Description *string `gorm:"column:description;type:TEXT;" json:"description"`
	Active      *bool   `gorm:"column:is_active;default:true" json:"active"`
}

func (BlogModel) TableName() string {
	return "blogs"
}
