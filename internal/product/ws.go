package product

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var StockUpdateChannel = make(chan Product)

func WebSocketAlerts(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer conn.Close()

    go func() {
        for p := range StockUpdateChannel {
            conn.WriteJSON(gin.H{"update": p})
        }
    }()

    for {
        lowStock, err := GetLowStockProducts()
        if err != nil {
            conn.WriteJSON(gin.H{"error": "Erro ao buscar produtos"})
            return
        }

        conn.WriteJSON(gin.H{"low_stock": lowStock})
        time.Sleep(5 * time.Second)
    }
}
