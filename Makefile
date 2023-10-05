.SILENT:

.DEFAULT_GOAL := run

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up -d

.PHONY: test
test:
	go test -v ./...

.PHONY: migrate
migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

.PHONY: create-migration
create-migration:
	migrate create -ext sql -dir schema/ -seq $(NAME)

.PHONY: swag
swag:
	swag init -g cmd/main.go


