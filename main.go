package main

import (
	"net/http"
	"procurement/config"
	"procurement/database"

	"github.com/gin-gonic/gin"
)

func init() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
}

func main() {

	// db := database.DB

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping-pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
