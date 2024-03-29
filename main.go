package main

import (
	"database/sql"
	"log"

	"github.com/egekozan/simplebank/api"
	db "github.com/egekozan/simplebank/db/sqlc"
	"github.com/egekozan/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("config file not load")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not open database connection")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server")
	}
}
