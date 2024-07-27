export APIKEY=<Gemini APIKEY>
export ENV="CI"
export FILEPATH="_example/example.md"

.PHONY: build
build:
	@echo "Building..."
	@GOOS=linux GOARCH=amd64 go build -o build/main main.go

.PHONY: run
run:
	@echo "Running..."
	@go run main.go
