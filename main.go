package main

import (
	"database/sql"
	"log"

	"github.com/egekozan/simplebank/api"
	db "github.com/egekozan/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not open database connection")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
}
