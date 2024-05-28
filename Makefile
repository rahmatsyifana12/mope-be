include .env

APP_NAME := go-boilerplate
SOURCE_PATH := ./src/
MIGRATION_DIR := ./migrations

DBMATE_URL := postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

.PHONY: build
build:
	go build -v -o bin/${APP_NAME} ./src

.PHONY: start
start:
	./bin/${APP_NAME}

.PHONY: run
run:
	go run ./src

.PHONY: compile
compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./src
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 ./src

.PHONY: migration-up
migration-up:
	migrate -database ${DBMATE_URL} -path migrations up

.PHONY: migration-down
migration-down:
	migrate -database ${DBMATE_URL} -path migrations down

.PHONY: migration-create
migration-create:
	@read -p "Enter the migration name: " MIGRATION_NAME; \
	migrate create -ext sql -dir $(MIGRATION_DIR) $$MIGRATION_NAME

.PHONY: migration-down-1
migration-down-1:
	migrate -database ${DBMATE_URL} -path migrations down 1