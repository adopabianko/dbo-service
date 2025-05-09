include .env
DB_URL_MIGRATION = $(DB_CONNECTION_MIGRATION)
BIN = build/dbo-service
FLAGS = -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo

.PHONY: setup
setup:
	@go install github.com/air-verse/air@v1.61.7
	@go install github.com/swaggo/swag/cmd/swag@v1.16.4
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2
	@go mod download
	@$(MAKE) swaggo

.PHONY: swaggo
swaggo:
	@/bin/rm -f ./docs/docs.go
	@`go env GOPATH`/bin/swag init --parseDependency -g ./main.go

.PHONY: run
run: 
	@`go env GOPATH`/bin/swag init --parseDependency -g ./main.go && `go env GOPATH`/bin/air

.PHONY: build
build:
	@go build ${FLAGS} -o $(BIN) .

.PHONY: migrate-up
migrate-up:
	@`go env GOPATH`/bin/migrate -path db/migrations/ -database $(DB_URL_MIGRATION) -verbose up

.PHONY: migrate-down
migrate-down:
	@`go env GOPATH`/bin/migrate -path db/migrations/ -database $(DB_URL_MIGRATION) -verbose down

.PHONY: docker-up
docker-up:
	cp -R .env.example .env
	docker compose up --build

.PHONY: docker-down
docker-down:
	docker compose down --rmi all -v --remove-orphans
