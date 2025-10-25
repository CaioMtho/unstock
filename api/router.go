package api

import (
	"github.com/CaioMtho/unstock/internal/product"
	"github.com/gin-gonic/gin"
)

func SetupEndpoints() *gin.Engine {
	r := gin.Default()

	productGroup := r.Group("/products")
    {
        productGroup.GET("/", product.GetAllHandler)
        productGroup.GET("/:id", product.GetByIDHandler)
        productGroup.POST("/", product.CreateHandler)
        productGroup.PATCH("/:id/stock", product.UpdateStockHandler)
        productGroup.GET("/alerts", product.GetLowStockHandler)
        productGroup.DELETE("/:id", product.DeleteHandler)
        productGroup.PUT("/:id", product.UpdateHandler)
    }

	r.GET("ws/alerts", product.WebSocketAlerts)

	return r
}