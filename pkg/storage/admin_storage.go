package storage

import (
	"github.com/AntonyCarl/OMA-Library/internal/models"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
)

func (storage *Storage) AddAdmin(req models.SignUpRequest) error {
	_, err := storage.db.Exec("INSERT INTO admins (username, email, password) VALUES ($1, $2, $3)",
		req.Username, req.Email, req.Password)
	if err != nil {
		logger.Logger.Error(err)
	}
	return err
}

func (storage *Storage) CheckExist(email string) bool {
	var exists bool
	err := storage.db.QueryRow("SELECT EXISTS(SELECT 1 FROM admins WHERE email=$1)", email).Scan(&exists)
	if err != nil {
		logger.Logger.Error(err)
	}
	return exists
}
