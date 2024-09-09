package models

import "gorm.io/gorm"

// Cart items model
type Cart_Item struct {
	gorm.Model
	Item_Id int `json:"item_id"`
	Amount int `json:"amount"`
}
