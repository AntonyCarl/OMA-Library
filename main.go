package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal/config"
	"github.com/AntonyCarl/OMA-Library/internal/handlers"
	"github.com/AntonyCarl/OMA-Library/internal/storage"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.SetConfig()
	logger.Init()
	storage, _ := storage.NewStorage(cfg)
	handlers.RunWeb(storage)
	http.ListenAndServe(":8080", nil)
}
