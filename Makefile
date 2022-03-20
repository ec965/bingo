server:
	go build -o server cmd/server/main.go

.PHONY: dev fmt db

dev:
	go run cmd/server/main.go

fmt:
	gofmt -s -w .

db:
	psql postgresql://postgres:postgres@localhost:5432/bingo
