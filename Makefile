test:
	go test -v ./...

run:
	docker-compose up -d

generate:
	go run github.com/99designs/gqlgen generate
