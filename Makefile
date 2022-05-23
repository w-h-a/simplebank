postgres:
	docker run --name postgres12 -p 3500:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

up:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:3500/simple_bank?sslmode=disable" --verbose up

upOne:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:3500/simple_bank?sslmode=disable" --verbose up 1

down:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:3500/simple_bank?sslmode=disable" --verbose down

downOne:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:3500/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb up upOne down downOne sqlc test server