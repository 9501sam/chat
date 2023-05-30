deps:
	go get github.com/gorilla/websocket
ser:
	@go run cmd/server/main.go
cli:
	@go run cmd/client/main.go
