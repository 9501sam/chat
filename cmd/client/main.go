package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connection:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server!")
}
