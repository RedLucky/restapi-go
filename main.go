package main

import (
	"fast-trade/config"
	"fast-trade/controllers"
	"fast-trade/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	controller := &controllers.Connection{Query: db}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lucky"
		msg.Message = "Fernanda"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)

	})
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/trade")
	{
		v1.GET("/:id", services.Auth, controller.GetTradeById)
		v1.GET("", services.Auth, controller.GetTradeHistory)
		v1.POST("", services.Auth, controller.AddTrade)
		v1.PUT("", services.Auth, controller.EditTrade)
		r.DELETE("/:id", services.Auth, controller.DeleteTrade)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Run(":3000")
}
