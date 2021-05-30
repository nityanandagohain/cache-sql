package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbDriver = "mysql"
	dbSource = "root:pass@tcp(localhost:3306)/cache"
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
