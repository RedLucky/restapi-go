package main

import (
	"fast-trade/config"
	"fast-trade/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.Connection{Query: db}
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

	r.GET("/trade/:id", inDB.GetTradeHistory)
	r.Run(":3000")
}
