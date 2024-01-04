package controllers

import (
	"compare/models"
	"fmt"
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

	id := connectedCnt
	models.Players[id] = 0
	fmt.Printf("Player%d connected.\n", id)

	models.WaitingNum <- connectedCnt

	// 添加到connections映射
	connectedMu.Lock()
	connections[id] = conn
	connectedMu.Unlock()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("玩家%d已連接，等待遊戲開始\n", id))); err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		HandlePlayerMessage(conn, id, string(msg))
	}
}
