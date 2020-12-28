test:
	go test -v ./...

run:
	docker-compose up -d  --build

generate:
	go run github.com/99designs/gqlgen generate

#create migrate:
#	migrate create -ext sql -dir mysql -seq create_users_table