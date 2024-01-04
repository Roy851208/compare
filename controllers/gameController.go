package controllers

import (
	"compare/models"
	"fmt"
	"math/rand"
	"time"
)

func StartGame() {
	for {
		<-models.WaitingNum
		// 玩家進入後告訴客戶端是否準備好
		if len(models.Players) == 2 {
			for {
				time.Sleep(1 * time.Second)
				sendMessageToAll(models.GameStartMessage)
				// 等待兩位玩家都準備好
				for !CheckReady() {
					time.Sleep(5 * time.Second)
				}
				GetCard()
				time.Sleep(2 * time.Second)
				Result()
				time.Sleep(2 * time.Second)
				ResetReady()
			}
		}
	}
}

func GetCard() {
	for id := range models.Players {
		models.Players[id] = rand.Intn(13) + 1
		sendMessageToAll(fmt.Sprintf("Players%d 獲得點數為 %d", id, models.Players[id]))
	}
	if models.Players[0] == models.Players[1] {
		sendMessageToAll("點數相同，重新分配點數")
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
	sendMessageToAll(fmt.Sprintf("贏家為player%d\n遊戲即將重新開始\n", winnerId))
}

func CheckReady() bool {
	if models.PlayersReady[1] && models.PlayersReady[2] {
		sendMessageToAll(models.PlayerReadyMessage)
		return true
	} else {
		sendMessageToAll(models.PlayerNotReadyFormat)
		return false
	}
}

func ResetReady() {
	for id := range models.PlayersReady {
		models.PlayersReady[id] = false
	}
}
