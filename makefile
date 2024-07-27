.PHONY: build
build:
	@echo "Building..."
	@GOOS=linux GOARCH=amd64 go build -o build/main main.go
