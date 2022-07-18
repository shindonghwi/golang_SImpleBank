package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	println("TestMain Start")
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}