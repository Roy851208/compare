package controllers

import (
	"compare/models"
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

func StartGame() {
	for {
		<-models.WaitingNum
		fmt.Println("新玩家入場")
		if len(models.Players) == 2 {
			fmt.Println("遊戲開始")
			GetCard()
			Result()
		}

	}
}

func GetCard() {
	for i := range models.Players {
		models.Players[i] = rand.Intn(13) + 1
		SendMessageToClient(i, fmt.Sprintf("models.Players%d 獲得點數為 %d\n", i, models.Players[i]))
		fmt.Printf("models.Players%d 獲得點數為 %d\n", i, models.Players[i])
	}
	if models.Players[0] == models.Players[1] {
		fmt.Println("點數相同，重新分配點數")
		GetCard()
	}
}

func Result() {
	result := 0
	winnerId := 0
	for i, v := range models.Players {
		if v > result {
			result = v
			winnerId = i
		}
	}
	fmt.Printf("贏家為player%d\n", winnerId)
	// 發送結果消息給客戶端
	for id := range models.Players {
		SendMessageToClient(id, fmt.Sprintf("贏家為player%d\n", winnerId))
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
