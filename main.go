package main

import (
	"database/sql"
	"log"

	"github.com/christy712/simplebank/api"
	db "github.com/christy712/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable"
	sereveradress = "0.0.0.0:8080"
)

func main() {

	cnn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connection failed :", err)
	}

	store := db.NewStore(cnn)
	server := api.NewServer(store)
	err = server.Start(sereveradress)
	if err != nil {
		log.Fatal("Cannot Start Server :", err)
	}
}
