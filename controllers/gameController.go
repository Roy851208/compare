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
		fmt.Println("新玩家入場")

		// 玩家進入後告訴客戶端是否準備好
		if len(models.Players) == 2 {
			for {
				for id := range models.Players {
					SendMessageToClient(id, "遊戲即將開始，請輸入 ready 來完成準備")
				}
				// 等待兩位玩家都準備好
				for !CheckReady() {
					time.Sleep(100 * time.Second)
					for id := range models.Players {
						SendMessageToClient(id, "遊戲即將開始，請輸入 ready 來完成準備")
					}
				}
				fmt.Println("遊戲開始")
				GetCard()
				Result()
				ResetReady()
			}
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

func CheckReady() bool {
	if models.PlayersReady[1] && models.PlayersReady[2] {
		for id := range models.Players {
			SendMessageToClient(id, "玩家已準備完成，即將開始遊戲")
		}
		return true
	} else if !models.PlayersReady[1] && models.PlayersReady[2] {
		for id := range models.Players {
			SendMessageToClient(id, fmt.Sprintf("玩家%d尚未準備完成", id))
		}
		return false
	} else if models.PlayersReady[1] && !models.PlayersReady[2] {
		for id := range models.Players {
			SendMessageToClient(id, fmt.Sprintf("玩家%d尚未準備完成", id))
		}
		return false
	} else {
		for id := range models.Players {
			SendMessageToClient(id, "雙方都尚未準備")
		}
		return false
	}
}

func ResetReady() {
	for id := range models.PlayersReady {
		models.PlayersReady[id] = false
	}
}
