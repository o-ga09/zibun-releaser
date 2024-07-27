.PHONY: build
build:
	@echo "Building..."
	@GOOS=linux GOARC=amd64 go build -o build/main main.go
