package main

import (
	"compare/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化Gin
	r := gin.Default()

	// 設定路由
	r.GET("/ws", controllers.HandleWebSocket)
	go controllers.StartGame()

	// 啟動伺服器
	r.Run(":8080")
}
