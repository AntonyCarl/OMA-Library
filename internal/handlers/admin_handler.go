package handlers

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal/models"
	"github.com/AntonyCarl/OMA-Library/internal/utils"
	"github.com/AntonyCarl/OMA-Library/pkg/storage"
	"github.com/labstack/echo/v4"
)

func RegisterAdmin(storage *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		var signUpReq models.SignUpRequest
		err := c.Bind(&signUpReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "wrong data"})
		}
		err = c.Validate(signUpReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}
		exist := storage.CheckExist(signUpReq.Email)
		if exist {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "Email already exist"})
		}
		signUpReq.Password = utils.HashPassword(signUpReq.Password)
		err = storage.AddAdmin(signUpReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cant create admin"})
		}
		return c.JSON(http.StatusCreated, nil)
	}
}
