package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:changeforproduction@localhost/postgres?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("failed to connect to database. %v", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
