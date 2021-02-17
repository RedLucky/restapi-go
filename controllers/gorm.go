package controllers

import "github.com/jinzhu/gorm"

type Connection struct {
	Query *gorm.DB
}
