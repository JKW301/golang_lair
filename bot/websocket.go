package bot

import (
    "log"
    "github.com/gorilla/websocket"
)

// StreamBinance connecte au flux WebSocket de Binance
func StreamBinance(symbol string, ch chan<- float64) {
    url := "wss://stream.binance.com:9443/ws/" + symbol + "@trade"
    conn, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        log.Fatalf("Failed to connect to Binance WebSocket: %v", err)
    }
    defer conn.Close()

    for {
        var message map[string]interface{}
        if err := conn.ReadJSON(&message); err != nil {
            log.Printf("Error reading Binance WebSocket: %v", err)
            return
        }

        price, ok := message["p"].(string)
        if ok {
            // Convert string price to float and send to channel
            ch <- price
        }
    }
}
