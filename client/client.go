package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5112/chathub", nil)
	if err != nil {
		panic(err)
	}

	// 发送握手消息
	err = sendHandshakeMessage(conn, `{"protocol":"json","version":1}`)
	if err != nil {
		panic(err)
	}

	// ping
	go func() {
		for {
			data := []byte(`{"type":6}`)
			data = append(data, 0x1e)
			conn.WriteMessage(1, data)
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(p))
	}
}

func sendHandshakeMessage(conn *websocket.Conn, message string) error {
	data := []byte(message)
	data = append(data, 0x1e)
	return conn.WriteMessage(1, data)
}
