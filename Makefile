build:
	@go build -o bin/jot-fusion

run: build
	@./bin/jot-fusion
	
test:
	@go test -v ./...
