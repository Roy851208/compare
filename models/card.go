package models

var Players = make(map[int]int)
var WaitingNum = make(chan int)
var PlayersReady = make(map[int]bool)
