package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	"github.com/AntonyCarl/OMA-Library/pkg/psql"
	_ "github.com/lib/pq"
)

func main() {
	logger.Init()
	logger.Logger.Info("start")
	psql.DbConnection()
	internal.RunWeb()
	http.ListenAndServe(":8080", nil)
}
