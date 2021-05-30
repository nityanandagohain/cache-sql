.PHONY: postgres migrate sqlc

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=changeforproduction postgres

mysql:
	docker run --rm -ti -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=pass mysql:5.7.34

mysql-cache:
	docker run --rm -ti -p 3306:3306 -v /home/nityananda/projects/go/src/github.com/nityanandagohain/sql-cache/conf:/etc/mysql/conf.d --name mysql -e MYSQL_ROOT_PASSWORD=pass mysql:5.7.34
	
migrate:
	migrate -path db/migration \
			-database mysql://root:changeforproduction@tcp(localhost:3306)/mysql up

migrate-down:
	migrate -path db/migration \
			-database postgres://postgres:changeforproduction@localhost/postgres?sslmode=disable down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...