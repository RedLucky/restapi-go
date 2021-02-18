package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TradeHistory struct {
	gorm.Model
	// Id     string     `json:"id" gorm:"primary_key"`
	Symbol string     `json:"symbol"`
	Date   *time.Time `json:"date"`
	Buy    float32    `json:"buy"`
	Sell   float32    `json:"sell"`
	Lot    uint32     `json:"lot"`
	Type   string     `json:"type"`
}
