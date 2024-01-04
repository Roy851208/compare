package controllers

import (
	"compare/models"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader     = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	connectedMu  sync.Mutex
	connectedCnt int
	connections  = make(map[int]*websocket.Conn)
)

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// 新玩家進入，增加計數
	connectedMu.Lock()
	connectedCnt++
	connectedMu.Unlock()
	models.WaitingNum <- connectedCnt

	id := connectedCnt
	models.Players[id] = 0
	log.Printf("Player%d connected.\n", id)

	// 添加到connections映射
	connectedMu.Lock()
	connections[id] = conn
	connectedMu.Unlock()

	SendMessageToClient(id, fmt.Sprintf("玩家%d已連接，等待遊戲開始\n", id))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		HandlePlayerMessage(conn, id, string(msg))
	}
}
