.PHONY: all run build migrate test tidy deps clean

all: build

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

migrate:
	go run ./cmd/server

test:
	go test ./... -v -count=1

tidy:
	go mod tidy
	go mod verify

deps:
	go mod download

clean:
	rm -rf bin/ tmp/
