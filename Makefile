.PHONY: default run build test clean

APP_NAME=devbook

default: run

run:
	@go run cmd/main.go
build:
	@go build -o $(APP_NAME) cmd/main.go
test:
	@go test ./ ...
clean:
	@rm -rf $(APP_NAME)