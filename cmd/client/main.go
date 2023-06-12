package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type (
	Message struct {
		Text      string    `json:"text"`
		Timestamp time.Time `json:"timestamp"`
		SenderID  string    `json:"senderid"`
		Code      string    `json:"code"`
	}
)

var (
	port = "8080"
)

// Receives message from server
func receive(ws *websocket.Conn) {
	var m Message
	for {
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println("You have disconnected")
		}
		fmt.Printf("m = %v\n", m)
	}
}

func main() {
	conn, err := websocket.Dial(fmt.Sprintf("ws://localhost:%s", port), "",
		"http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connection:", err)
	}
	defer conn.Close()

	// receive
	go receive(conn)

	fmt.Println("Connected to server!")
}
