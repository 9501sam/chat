deps:
	go get github.com/gorilla/websocket
	go get github.com/9501sam/chat/shared
ser:
	@go run cmd/server/main.go
cli:
	@go run cmd/client/main.go
