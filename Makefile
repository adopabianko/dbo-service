BIN = build/dbo-service
FLAGS = -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo

.PHONY: setup
setup:
        @go install github.com/swaggo/swag/cmd/swag@v1.16.4

.PHONY: swaggo
swaggo:
        @rm -f ./docs/docs.go
        @swag init --parseDependency -g ./main.go

.PHONY: docker-up
docker-up: setup swaggo
        @cp -R .env.example .env
        @docker compose up --build

.PHONY: docker-down
docker-down:
        @docker compose down --rmi all -v --remove-orphans