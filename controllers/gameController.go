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
		if len(models.Players) == 3 {
			for {
				time.Sleep(1 * time.Second)
				SendMessageToAll(models.GameStartMessage)
				// 等待玩家都準備好
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
	checkReAssign := make(map[int]bool)
	for id := range models.Players {
		for {
			models.Players[id] = rand.Intn(13) + 1
			if _, exists := checkReAssign[models.Players[id]]; !exists {
				checkReAssign[models.Players[id]] = true
				break
			}
		}
		SendMessageToAll(fmt.Sprintf("Players%d 獲得點數為 %d", id, models.Players[id]))
	}
}

func Result() {
	result, winnerId := 0, 0
	for id, point := range models.Players {
		if point > result {
			result = point
			winnerId = id
		}
	}
	SendMessageToAll(fmt.Sprintf("贏家為player%d\n遊戲即將重新開始\n", winnerId))
}

func CheckReady() bool {
	for _, ready := range models.PlayersReady {
		if !ready {
			SendMessageToAll(models.PlayerNotReadyFormat)
			return false
		}
	}
	SendMessageToAll(models.PlayerReadyMessage)
	return true
}

func ResetReady() {
	for id := range models.PlayersReady {
		models.PlayersReady[id] = false
	}
}
