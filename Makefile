include .env

export 


run: postgres-up
	@echo "Running the application"
	@go run ./cmd/main.go

build:
	@echo "Building the application"
	@go build -o ./dist/main ./cmd/main.go

postgres-up:
	@echo "Starting postgres container"
	@docker-compose -f ./script/docker-compose-postgres.yml up -d --force-recreate

postgres-down:
	@echo "Stopping postgres container"
	@docker-compose -f ./script/docker-compose-postgres.yml down -v --remove-orphans