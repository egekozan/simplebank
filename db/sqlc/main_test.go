package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/egekozan/simplebank/db/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not open env file")
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not open database connection")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
