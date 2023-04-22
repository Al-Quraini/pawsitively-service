package main

import (
	"database/sql"
	"log"

	"github.com/alquraini/pawsitively/api"
	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	action := db.NewAction(conn)
	server, err := api.NewServer(config, action)
	if err != nil {
		log.Fatal("connot cerate server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
