package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func sendMsg(ws *websocket.Conn, text string) error {
	m := Message{
		Text:      text,
		Timestamp: time.Now(),
		SenderID:  "asdf",
		Code:      "hi",
	}
	err := websocket.JSON.Send(ws, m)
	return err
}

func send(ws *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "":
			continue
		default:
			err := sendMsg(ws, text)
			if err != nil {
				log.Println("Error sending message: ", err.Error())
				break
			}
		}
	}
}

func main() {
	conn, err := websocket.Dial(fmt.Sprintf("ws://localhost:%s", port), "",
		"http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connection:", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server!")

	// receive
	go receive(conn)

	// send
	send(conn)
}
