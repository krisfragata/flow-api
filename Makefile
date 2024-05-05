build:
	@go build -o bin/flow-api

run: build
	@./bin/flow-api

test:
	@go test -v ./...