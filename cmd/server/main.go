package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
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

	room struct {
		clients          map[string]*websocket.Conn
		addClienChan     chan *websocket.Conn
		removeClientChan chan *websocket.Conn
		broadcastChan    chan Message
	}
)

var (
	port = "8080"
	once sync.Once
)

func newRoom() *room {
	return &room{
		clients:          make(map[string]*websocket.Conn),
		addClienChan:     make(chan *websocket.Conn),
		removeClientChan: make(chan *websocket.Conn),
	}
}

func (rm *room) addClient(conn *websocket.Conn) {
}

func (rm *room) removeClient(conn *websocket.Conn) {
}

func (rm *room) broadcast(m Message) {
	for _, conn := range rm.clients {
		if err := websocket.JSON.Send(conn, m); err != nil {
			log.Println("Error broadcasting")
		}
	}
}

func (rm *room) run() {
	for {
		select {
		case conn := <-rm.addClienChan:
			rm.addClient(conn)
		case conn := <-rm.removeClientChan:
			rm.removeClient(conn)
		case m := <-rm.broadcastChan:
			rm.broadcast(m)
		}
	}
}

func handler(ws *websocket.Conn, rm *room) {
	fmt.Printf("new client\n")
	go once.Do(rm.run)
	rm.addClienChan <- ws

	var m Message
	for {
		err := websocket.JSON.Receive(ws, &m)

		if err != nil {
			rm.removeClientChan <- ws
			return
		}

		rm.broadcastChan <- m
	}
}

func connect(port string) error {
	rm := newRoom()

	mux := http.NewServeMux()

	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws, rm)
	}))

	s := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	fmt.Println("Server running on port: ", port)
	return s.ListenAndServe()
}

func main() {
	if err := connect(port); err != nil {
		log.Println(err.Error())
	}
}
