package controllers

import (
	"compare/models"
	"fmt"

	"github.com/gorilla/websocket"
)

func HandlePlayerMessage(conn *websocket.Conn, playerId int, message string) {
	switch message {
	case "ready":
		// 玩家準備
		models.PlayersReady[playerId] = true
		SendMessageToClient(playerId, fmt.Sprintf("玩家%d準備完成", playerId))
	default:
		fmt.Printf("Player%d 傳入錯誤訊息\n", playerId)
	}
}

func sendMessageToAll(message string) {
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
