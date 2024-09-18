package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal"
	"github.com/AntonyCarl/OMA-Library/pkg/psql"
	_ "github.com/lib/pq"
)

func main() {
	psql.DbConnection()
	internal.RunWeb()
	http.ListenAndServe(":8080", nil)
}
