package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal/storage/postgres"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	logger.Init()
	storage, _ := postgres.NewStorage()
	logger.Logger.Info("start")
	handler.RunWeb(storage)
	http.ListenAndServe(":8080", nil)
}
