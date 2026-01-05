run:
	go run src/cmd/api/main.go

build:
	go build -o bin/api src/cmd/api/main.go

swagger:
	swag init -g src/cmd/api/main.go --parseInternal=true

dev:
	air

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-ps:
	docker-compose ps

docker-restart:
	docker-compose restart

docker-rebuild:
	docker-compose up -d --build