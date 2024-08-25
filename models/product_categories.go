package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	ID   uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Name string    `json:"name"`
	gorm.Model
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

func (pc *ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	pc.ID = uuid.New()
	return
}
