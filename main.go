package main

import (
	"github.com/AntonyCarl/OMA-Library/internal/config"
	"github.com/AntonyCarl/OMA-Library/internal/handlers"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	"github.com/AntonyCarl/OMA-Library/pkg/storage"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.SetConfig()
	logger.Init()
	storage, _ := storage.NewStorage(cfg)
	handlers.RunWeb(storage)
}
