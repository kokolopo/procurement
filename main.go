package main

import (
	"net/http"
	"procurement/config"
	"procurement/database"
	"procurement/domain/material"

	"github.com/gin-gonic/gin"
)

func init() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
}

func main() {

	db := database.DB // initial DB

	repository := material.NewMaterialRepository(db)
	service := material.NewMaterialService(repository)
	handler := material.NewClientHandler(service)

	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":        "hay, wellcome",
			"/materials":     "to get materials",
			"/materials/:id": "to get material",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/materials", handler.GetAll)
		v1.GET("/materials/:id", handler.GetById)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
