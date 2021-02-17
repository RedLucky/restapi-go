package controllers

import (
	"fast-trade/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get one data with {id}
func (idb *InDB) GetTradeHistory(c *gin.Context) {
	var (
		trade  models.TradeHistory
		result gin.H
	)
	id := c.Param("id")
	if err := idb.DB.Where("id = ?", id).First(&trade).Error; err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": trade,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}
