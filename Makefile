run:
	go run src/cmd/api/main.go

build:
	go build -o bin/api src/cmd/api/main.go

swagger:
	swag init -g src/cmd/api/main.go --parseInternal=true

dev:
	air