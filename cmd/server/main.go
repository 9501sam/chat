package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var (
	port = "8080"
)

func handler(ws *websocket.Conn) {
	fmt.Printf("got a message\n")
}

func connect(port string) error {
	mux := http.NewServeMux()

	mux.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handler(ws)
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
