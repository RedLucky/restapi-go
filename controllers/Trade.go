package controllers

import (
	"fast-trade/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// to get one data with {id}
func (m *Connection) GetTradeById(c *gin.Context) {
	var (
		trade  models.TradeHistory
		result gin.H
	)
	id := c.Param("id")
	if err := m.Query.Where("id = ?", id).First(&trade).Error; err != nil {
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

// to get all data
func (m *Connection) GetTradeHistory(c *gin.Context) {
	var (
		history []models.TradeHistory
		result  gin.H
	)

	m.Query.Find(&history)
	if len(history) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": history,
			"count":  len(history),
		}
	}

	c.JSON(http.StatusOK, result)
}

// insert data
func (m *Connection) AddTrade(c *gin.Context) {
	var (
		trade        models.TradeHistory
		result       gin.H
		layoutFormat = "2006-01-02 15:04:05"
	)

	lot, _ := strconv.ParseUint(c.PostForm("lot"), 10, 32)
	date, _ := time.Parse(layoutFormat, c.PostForm("date"))

	if c.PostForm("type") == "buy" {
		buy, _ := strconv.ParseFloat(c.PostForm("buy"), 32)
		trade.Buy = float32(buy)
		trade.Sell = 0
	} else if c.PostForm("type") == "sell" {
		sell, _ := strconv.ParseFloat(c.PostForm("sell"), 32)
		trade.Buy = 0
		trade.Sell = float32(sell)
	}

	trade.Symbol = c.PostForm("symbol")
	trade.Lot = uint32(lot)
	trade.Type = c.PostForm("type")
	trade.Date = &date

	m.Query.Create(&trade)
	result = gin.H{
		"result": trade,
	}

	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (m *Connection) EditTrade(c *gin.Context) {
	id := c.Query("id")

	var (
		trade        models.TradeHistory
		newTrade     models.TradeHistory
		result       gin.H
		layoutFormat = "2006-01-02 15:04:05"
	)

	err := m.Query.First(&trade, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	lot, _ := strconv.ParseUint(c.PostForm("lot"), 10, 32)
	date, _ := time.Parse(layoutFormat, c.PostForm("date"))

	if c.PostForm("type") == "buy" {
		buy, _ := strconv.ParseFloat(c.PostForm("buy"), 32)
		newTrade.Buy = float32(buy)
		newTrade.Sell = 0
	} else if c.PostForm("type") == "sell" {
		sell, _ := strconv.ParseFloat(c.PostForm("sell"), 32)
		newTrade.Buy = 0
		newTrade.Sell = float32(sell)
	}

	newTrade.Symbol = c.PostForm("symbol")
	newTrade.Lot = uint32(lot)
	newTrade.Type = c.PostForm("type")
	newTrade.Date = &date

	err = m.Query.Model(&trade).Updates(newTrade).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (m *Connection) DeleteTrade(c *gin.Context) {
	var (
		trade  models.TradeHistory
		result gin.H
	)
	id := c.Param("id")
	err := m.Query.First(&trade, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = m.Query.Delete(&trade, id).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
