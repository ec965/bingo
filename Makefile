pg_fallback = postgresql://postgres:postgres@localhost:5432/bingo

server:
	go build -o server cmd/server/main.go

.PHONY: dev fmt vet db migrate-up migrate-down tmux

# code
dev:
	go run cmd/server/main.go

fmt:
	gofmt -s -w .

vet:
	go vet ./...

# database
db:
	psql $(pg_url)

migrate-up:
	migrate -database \
		"$${DATABASE_URL:-$(pg_fallback)}?sslmode=disable" \
		-path db/migrations up

migrate-down:
	migrate -database \
		"$${DATABASE_URL:-$(pg_fallback)}?sslmode=disable" \
		-path db/migrations down

# dev
tmux:
	./scripts/tmux.sh

