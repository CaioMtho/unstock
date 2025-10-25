package product

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketAlerts(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer conn.Close()

	for {
		lowStock, err := GetLowStockProducts()
		if err != nil {
			conn.WriteJSON(gin.H{"error": "Erro ao buscar produtos"})
			return
		}

		conn.WriteJSON(lowStock)

		time.Sleep(5 * time.Second)

	}
}