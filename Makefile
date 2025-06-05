
generate:
	go generate ./...

initdb:
	go run cmd/init/main.go

migrate:
	go run cmd/migrate/main.go

unit-tests:
	go clean -testcache && go test ./internal/application/...

local-run:
	export APP_ENV=local && make migrate && make initdb && go run cmd/service/main.go

docker-drop:
	docker compose down -v db

docker-up-db:
	docker compose up --build -d db

docker-up:
	docker compose up --build -d

docker-run-integration:
	docker compose -f docker-compose-integration.yaml down -v && docker compose -f docker-compose-integration.yaml up --build avito-shop-integration-tests && docker compose -f docker-compose-integration.yaml down