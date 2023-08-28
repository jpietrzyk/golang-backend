package main

import (
	"context"
	"jck/golangbackend/api"
	db "jck/golangbackend/db/sqlc"
	"jck/golangbackend/util"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannoct connect to db:", err)
	}

	store := db.NewStore(connPool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	if err != nil {
		log.Fatal("sever panic:", err)
	}
}
