postgres:
	docker run --name some-postgresalpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:alpine  

createdb:
	docker exec -it some-postgresalpine createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it some-postgresalpine dropdb simple_bank

migrateup:
	/home/christy/go/bin/migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup_CI:
	migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	/home/christy/go/bin/migrate -path db/migration -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test migrateup_CI
 
