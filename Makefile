include app.env
export
run:
	@go run cmd/main.go
migrate_create:
	migrate create -ext sql -dir ./internal/storage/persistant/migration -seq delivery1
migrate_up:
	migrate -path ./internal/storage/persistant/migration  -database  "cockroach://root@:26257/delivery?sslmode=disable" -verbose up
migrate_down:
	migrate -path ./internal/storage/persistant/migration  -database  "cockroach://root@:26257/delivery?sslmode=disable" -verbose down
sqlc:
	sqlc generate