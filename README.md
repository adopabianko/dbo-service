## Requirement
Install go version >= 1.24
```bash
https://go.dev/doc/install
```

## Project setup
```bash
# using go install
go install github.com/air-verse/air@v1.61.7
go install github.com/swaggo/swag/cmd/swag@v1.16.4
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2
go mod download
# generate swagger
swag init --parseDependency -g ./main.go

# or

# using make
make setup
```

## Compile and run the project

```bash
# using go run
go run main.go

# or

# using make
make run
```

## Run tests

```bash
go run test
```

## Migrations
Create a migration file
```bash
# using migration cli
migrate create -ext sql -dir db/migrations/ -seq customer

# or

# using make
make migrate-up
```

Apply migrations
```bash
# using migration cli
migrate -path db/migrations/ -database 'postgres://postgres:password@localhost:5432/dbo?sslmode=disable' -verbose up

# or

# using make
make migrate-down
```

Rollback migrations
```bash
migrate -path db/migrations/ -database 'postgres://postgres:password@localhost:5432/dbo?sslmode=disable' -verbose down
```

## Access URL
```bash
http://localhost:8080
```