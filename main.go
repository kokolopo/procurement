package main

import (
	"net/http"
	"procurement/config"
	"procurement/database"
	"procurement/domain/material"
	"procurement/domain/purchase_request"

	"github.com/gin-gonic/gin"
)

func init() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
}

func main() {

	db := database.DB // initial DB

	materialRepository := material.NewMaterialRepository(db)
	materialService := material.NewMaterialService(materialRepository)
	materialHandler := material.NewMaterialHandler(materialService)

	purchaseReqRepo := purchase_request.NewMaterialRepository(db)
	purchaseService := purchase_request.NewPRService(purchaseReqRepo, materialRepository)
	purchaseHandler := purchase_request.NewPurchaseRequestHandler(purchaseService)

	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":           "hay, wellcome",
			"/v1/materials":     "to get materials",
			"/v1/materials/:id": "to get material",
			"/v1/purchase-requests": "to get purchase request",
			"/v1/purchase-requests/:id": "to get purchase request",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/materials", materialHandler.GetAll)
		v1.GET("/materials/:id", materialHandler.GetById)

		v1.POST("/purchase-requests", purchaseHandler.CreateNewPurchaseRequest)
		v1.GET("/purchase-requests", purchaseHandler.GetAll)
		v1.GET("/purchase-requests/:id", purchaseHandler.GetById)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
