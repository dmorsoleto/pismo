include .env

export 


run: postgres-up
	./bin/wait-for-it
	@echo "Running the application"
	@go run ./cmd/main.go

postgres-up:
	@echo "Starting postgres container"
	@docker-compose -f ./script/docker-compose-postgres.yml up -d --force-recreate

postgres-down:
	@echo "Stopping postgres container"
	@docker-compose -f ./script/docker-compose-postgres.yml down -v --remove-orphans

local-api:
	@echo "Starting local api"
	@docker-compose -f ./script/docker-compose-api.yml up -d --force-recreate