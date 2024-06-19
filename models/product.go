package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float32 `json:"price" validate:"required,numeric"`
	Instock  bool    `json:"instock" validate:"required,boolean"`
	Category string  `json:"category" validate:"required"`
}
