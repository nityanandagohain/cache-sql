## Using postgres as a cache



## steps to run
* `make postgres`
* `make migrate`
* `make sqlc`
* test `go test ./db/sqlc`
* benchmark `go test -bench=. -benchtime 2000x ./db/sqlc`