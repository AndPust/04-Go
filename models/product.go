package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price int `json:"price"`
	Category_Id int `json:"category_id"`
}


