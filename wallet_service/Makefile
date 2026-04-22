.PHONY: run build test migrate-up migrate-down psql

DATABASE_URL ?= postgres://wallet:wallet@localhost:55432/wallet?sslmode=disable

run:
	DATABASE_URL="$(DATABASE_URL)" PORT=8080 go run ./cmd/server

build:
	go build ./...

test:
	go test ./... -race -count=1

migrate-up:
	docker compose run --rm migrate \
	  -path /migrations \
	  -database "postgres://wallet:wallet@postgres:5432/wallet?sslmode=disable" \
	  up

migrate-down:
	docker compose run --rm migrate \
	  -path /migrations \
	  -database "postgres://wallet:wallet@postgres:5432/wallet?sslmode=disable" \
	  down 1

psql:
	docker compose exec postgres psql -U wallet -d wallet
