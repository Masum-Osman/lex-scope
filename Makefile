.PHONY: build run test docker-up docker-down

build:
	go build -o lex-scope cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test ./... -coverprofile=coverage.out

docker-up:
	docker-compose -f build/docker-compose.yml up -d

docker-down:
	docker-compose -f build/docker-compose.yml down
