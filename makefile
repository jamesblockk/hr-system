.PHONY: build run test docker-compose-up docker-compose-up-db-only 

test:
	go test ./employee-service/internal/...

mod:
	go mod tidy && go mod vendor

swag:
	swag init --generalInfo ./api-gateway/main.go

docker-compose-up:
	docker-compose  -f ./docker/docker-compose.yaml up --build

docker-compose-up-db-only:
	docker-compose  -f ./docker/docker-compose-db-only.yaml up --build

gorm-gen:
	go run common/database/gorm-gen/gen.go