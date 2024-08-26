package main

import (
	"log"
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal"
	"github.com/AntonyCarl/OMA-Library/pkg/database"
	_ "github.com/lib/pq"
)

func main() {

	// init database connection
	db, err := database.DbConnection(database.ConnectionParametrs{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "0000",
		DBName:   "omalibdb",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	internal.RunWeb()
	http.ListenAndServe(":8080", nil)
}
