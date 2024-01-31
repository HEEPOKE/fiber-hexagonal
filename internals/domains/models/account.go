package models

type AccountModel struct {
	CommonColumnModel
	Email    string      `gorm:"column:email;type:VARCHAR(255);unique;not null" json:"email" validate:"required"`
	UserName string      `gorm:"column:username;type:VARCHAR(255);unique;not null" json:"username" validate:"required"`
	Password *string     `gorm:"column:password;type:VARCHAR(255);" json:"password"`
	Age      int         `gorm:"column:age;type:INT;not null" json:"age" validate:"required"`
	IsActive *bool       `gorm:"column:is_active;default:true" json:"is_active"`
	Blogs    []BlogModel `gorm:"foreignKey:AccountID" json:"blogs"`
}

func (AccountModel) TableName() string {
	return "accounts"
}
