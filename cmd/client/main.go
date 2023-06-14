package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"chat-room/shared"

	"golang.org/x/net/websocket"
)

var (
	port = "8080"
	id   string
)

func generatedIP() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	id = fmt.Sprintf("http://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
	return id
}

func connect(port string) (*websocket.Conn, error) {
	conn, err := websocket.Dial(fmt.Sprintf("ws://localhost:%s", port), "", generatedIP())
	return conn, err
}

// Receives message from server
func receive(ws *websocket.Conn) {
	var m shared.Message
	for {
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println("You have disconnected")
		}
		fmt.Printf("m = %v\n", m)
	}
}

func sendMsg(ws *websocket.Conn, text string) error {
	m := shared.Message{
		Text:      text,
		Timestamp: time.Now(),
		SenderID:  id,
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
	// connect
	conn, err := connect(port)
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
