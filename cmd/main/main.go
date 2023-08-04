package main

import (
	"database/sql"
	"log"

	"github.com/StoicGoose/simple-bank/pkg/api"
	"github.com/StoicGoose/simple-bank/pkg/util"
	db "github.com/StoicGoose/simple-bank/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}