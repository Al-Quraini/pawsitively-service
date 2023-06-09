postgres:
	docker run --name postgres14 --network pawsitively-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root pawsitively

dropdb:
	docker exec -it postgres14 dropdb pawsitively

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pawsitively?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/pawsitively?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/action.go github.com/alquraini/pawsitively/db/sqlc Action

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock