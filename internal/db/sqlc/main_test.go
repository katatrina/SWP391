package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB    *sql.DB
	testStore *Store
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgres://postgres:12345@localhost:5432/bird_service_platform?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(testDB)

	// Run all tests by orders.

	os.Exit(m.Run())
}
