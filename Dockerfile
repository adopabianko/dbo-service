
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dbo-service .

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/dbo-service .

COPY .env .

EXPOSE 8080

CMD ["./dbo-service"]