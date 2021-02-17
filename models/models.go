package models

import "github.com/jinzhu/gorm"

type TradeHistory struct {
	gorm.Model
	ID     string `json:"id" gorm:"primary_key"`
	Symbol string
	Date   string
	Buy    float32
	Sell   float32
	Lot    uint32
}
