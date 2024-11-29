package api

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Permettre les connexions CORS pour tests
    },
}

func WebSocketHandler(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to upgrade WebSocket"})
        return
    }
    defer conn.Close()

    // Simule un envoi de données toutes les secondes
    for {
        price := "50000.00" // Exemple : Remplacez par des données du bot
        if err := conn.WriteJSON(gin.H{"price": price}); err != nil {
            break
        }
        time.Sleep(1 * time.Second)
    }
}
