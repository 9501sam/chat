package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

var (
	port = "8080"
)

func handler(ws *websocket.Conn) {
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
	connect(port)
}
