package config

import (
	"fast-trade/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/stocks?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(models.TradeHistory{})
	return db
}
