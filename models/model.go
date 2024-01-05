package models

var Players = make(map[int]int)
var WaitingNum = make(chan int)
var PlayersReady = make(map[int]bool)

const (
	GameStartMessage     = "遊戲即將開始，請輸入 ready 來完成準備"
	PlayerReadyMessage   = "玩家已準備完成，即將開始遊戲"
	PlayerNotReadyFormat = "有玩家尚未準備完成"
	GameRestartMessage   = "遊戲即將重新開始"
)
