package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 启动一个 goroutine 读取服务器发送的消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("接收到服务器消息: %s\n", message)
		}
	}()

	// 从终端读取用户输入并发送到服务器
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// 去除换行符
		input = strings.TrimSpace(input)

		// 发送消息到服务器
		err = conn.WriteMessage(websocket.TextMessage, []byte(input))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
