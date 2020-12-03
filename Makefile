test:
	go test ./...

run:
	go run server.go

migrate:
	migrate -database mysql://user:password@/zikanwarikun -path internal/db/migrations/mysql up