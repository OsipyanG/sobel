# Variables
APP_NAME := sobel

# Default goal
.PHONY: all build run  lint fmt clean
all: run

# Build the application
build:
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

# Run the application locally
run:
	@go run cmd/$(APP_NAME)/main.go -image ./images/image.jpg -numWorkers 16

bench:
	@go test -benchmem  -bench=. sobel/internal/worker

lint:
	@golangci-lint run

fmt:
	@go fmt ./...

clean:
	@rm -rf bin/ out/

