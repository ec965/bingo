server:
	go build -o server cmd/server/main.go

.PHONY: dev fmt

dev:
	go run cmd/server/main.go

fmt:
	gofmt -s -w .
