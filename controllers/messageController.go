package controllers

import (
	"compare/models"
	"fmt"

	"github.com/gorilla/websocket"
)

// MessageHandler 是一個處理玩家消息的函數型別
type MessageHandler func(conn *websocket.Conn, playerId int, message string)

// messageHandlers 是消息類型到處理函數的映射
var messageHandlers = map[string]MessageHandler{
	"ready": HandleReady,
	// 在這裡添加更多的消息類型和處理函數...
}

// HandlePlayerMessage 處理玩家傳來的消息
func HandlePlayerMessage(conn *websocket.Conn, playerId int, message string) {
	// 檢查消息類型是否有對應的處理函數
	if handler, exists := messageHandlers[message]; exists {
		// 呼叫處理函數處理消息
		handler(conn, playerId, message)
	} else {
		fmt.Printf("Player%d 傳入未知訊息: %s\n", playerId, message)
	}
}

// handleReady 是處理 "ready" 消息的函數
func HandleReady(conn *websocket.Conn, playerId int, message string) {
	// 在這裡處理 "ready" 消息的邏輯
	models.PlayersReady[playerId] = true
	SendMessageToAll(fmt.Sprintf("玩家%d準備完成", playerId))
	// 你可以在這裡添加更多的處理邏輯...
}

func SendMessageToAll(message string) {
	for id := range connections {
		SendMessageToClient(id, message)
	}
}

func SendMessageToClient(playerId int, message string) {
	connectedMu.Lock()
	defer connectedMu.Unlock()

	for id, conn := range connections {
		if id == playerId {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fmt.Println(err)
			}
			break
		}
	}
}
