package handlers

import (
	"net/http"
	"time"

	"github.com/AntonyCarl/OMA-Library/internal/models"
	"github.com/AntonyCarl/OMA-Library/internal/utils"
	"github.com/AntonyCarl/OMA-Library/pkg/storage"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("secret_token") //Must be at config??

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

func AdminLogin(storage *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		var signInReq models.SignInRequset
		err := c.Bind(&signInReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "wrong data"})
		}
		err = c.Validate(signInReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}
		admin, err := storage.GetByEmail(signInReq.Email)
		if err != nil || utils.CheckPasswordHash(signInReq.Password, admin.Password) != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid email or password" + err.Error()})
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": admin.Email,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not generate token"})
		}
		c.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(time.Hour * 2),
		})

		return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
	}
}
