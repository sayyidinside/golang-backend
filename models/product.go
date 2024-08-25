package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID       `json:"id" gorm:"primaryKey;type:char(36)"`
	CategoryID  uuid.UUID       `json:"category_id" gorm:"index;type:char(36)"`
	Name        string          `json:"name" gorm:"index;type:varchar(255)"`
	Description string          `json:"description" gorm:"type:text"`
	Category    ProductCategory `json:"category" gorm:"foreignKey:CategoryID"`
	gorm.Model
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
