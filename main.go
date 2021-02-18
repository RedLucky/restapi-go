package main

import (
	"fast-trade/config"
	"fast-trade/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	controller := &controllers.Connection{Query: db}
	r := gin.Default()

	/* r.GET("/", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)

	}) */

	v1 := r.Group("/trade")
	{
		v1.GET("/:id", controller.GetTradeById)
		v1.GET("", controller.GetTradeHistory)
		v1.POST("", controller.AddTrade)
		v1.PUT("", controller.EditTrade)
		r.DELETE("/:id", controller.DeleteTrade)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Run(":3000")
}
