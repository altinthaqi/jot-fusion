build:
	@go build -o bin/jot-fusion

start-docker:
	@docker-compose up -d

run: start-docker build
	@./bin/jot-fusion
	
test:
	@go test -v ./...
