package utils

import (
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Error(err)
	}
	return string(bytes)
}
