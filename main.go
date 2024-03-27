package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/ws", func(response http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(response, request, nil)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		// 处理WebSocket连接
		for {
			// 读取消息
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Println(messageType, "Received message:", string(p))

			// 发送消息
			err = conn.WriteMessage(messageType, []byte("Hello, world!"))
			if err != nil {

				return
			}
		}

	})

	http.ListenAndServe(":8080", handler)
}
