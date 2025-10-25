package api

import (
	"github.com/CaioMtho/unstock/internal/product"
	"github.com/gin-gonic/gin"
)

func SetupEndpoints() *gin.Engine {
	r := gin.Default()

    r.Static("/static", "./web")

    r.GET("/", func(c *gin.Context) {
        c.File("./web/index.html")
    })

	productGroup := r.Group("api/products/")
    {
        productGroup.GET("/", product.GetAllHandler)
        productGroup.GET("/:id", product.GetByIDHandler)
        productGroup.POST("/", product.CreateHandler)
        productGroup.PATCH("/:id/stock", product.UpdateStockHandler)
        productGroup.GET("/alerts", product.GetLowStockHandler)
        productGroup.DELETE("/:id", product.DeleteHandler)
        productGroup.PUT("/:id", product.UpdateHandler)
    }

	r.GET("api/ws/alerts", product.WebSocketAlerts)

	return r
}