.PHONY: postgres migrate sqlc

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=changeforproduction postgres
	
migrate:
	migrate -path db/migration \
			-database postgres://postgres:changeforproduction@localhost/postgres?sslmode=disable up

migrate-down:
	migrate -path db/migration \
			-database postgres://postgres:changeforproduction@localhost/postgres?sslmode=disable down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...