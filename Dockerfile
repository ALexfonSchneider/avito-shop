FROM golang:1.24.2-alpine as app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

ENV APP_CONFIG_PATH=/app/config
ENV APP_ENV=docker

EXPOSE 8080

CMD go run cmd/migrate/main.go; go run cmd/init/main.go; go run cmd/service/main.go