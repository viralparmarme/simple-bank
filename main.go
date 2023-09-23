package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/viralparmarme/simple-bank/api"
	db "github.com/viralparmarme/simple-bank/db/sqlc"
	"github.com/viralparmarme/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Can't connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
