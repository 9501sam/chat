package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

var (
	port = "8080"
)

func main() {
	conn, err := websocket.Dial(fmt.Sprintf("ws://localhost:%s", port), "", "http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connection:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server!")
}
