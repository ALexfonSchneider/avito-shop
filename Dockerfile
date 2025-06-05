FROM golang:1.24.2-alpine as app

WORKDIR /app

COPY . .

RUN go mod download  \
    && go mod tidy

ENV APP_CONFIG_PATH=/app/config
ENV APP_ENV=docker

EXPOSE 8080

CMD go run cmd/migrate/main.go; go run cmd/init/main.go; go run cmd/service/main.go

FROM app as tests

WORKDIR /app

ENV APP_CONFIG_PATH=/app/config
ENV APP_ENV=test

CMD go run cmd/migrate/main.go; go run cmd/init/main.go; go test -tags=integrational /app/test/integration/...