package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/AntonyCarl/OMA-Library/internal/models"
	"github.com/AntonyCarl/OMA-Library/internal/utils"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
	"github.com/AntonyCarl/OMA-Library/pkg/repository"
	"github.com/AntonyCarl/OMA-Library/pkg/storage"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RunWeb(storage *storage.Storage) {
	e := echo.New()
	e.Renderer = utils.NewTemplate("templates/*.html")
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	resGroup := e.Group("/admin")
	resGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  jwtSecret,
		TokenLookup: "cookie:jwt",
	}))

	e.GET("/", mainPageHandler)
	//e.GET("/upload", uploadFormHandler)
	//e.POST("/upload_file", uploadFileHandler(storage))
	e.GET("/search", searchHandler(storage))
	e.GET("/oma/:id", dowloadHandler(storage))
	e.POST("/register", RegisterAdmin(storage))
	e.GET("/login", AdminLogin(storage))
	resGroup.GET("/upload", uploadFormHandler)
	resGroup.POST("/upload_file", uploadFileHandler(storage))

	e.Start(":8080") // add to config

}

func mainPageHandler(c echo.Context) error {
	if err := c.Render(http.StatusOK, "index", nil); err != nil {
		logger.Logger.Error(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}

func uploadFormHandler(c echo.Context) error {
	if err := c.Render(http.StatusOK, "upload", nil); err != nil {
		logger.Logger.Error(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}

// hendlers working with db:

func uploadFileHandler(storage *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		fileHeader, err := c.FormFile("uploaded_file")
		if err != nil {
			logger.Logger.Error(err)
			c.String(http.StatusBadRequest, "Cant get file: "+err.Error())
		}
		file, err := fileHeader.Open()
		if err != nil {
			logger.Logger.Error(err)
			c.String(http.StatusBadRequest, "Cant get file: "+err.Error())
		}

		path := repository.SaveFile(file, fileHeader.Filename)
		if !strings.HasSuffix(fileHeader.Filename, ".oma") {
			logger.Logger.Info("Not oma")
			c.String(http.StatusBadRequest, "Invalid file format. Only .oma files are allowed"+err.Error())
		}

		omafile := models.NewOmafile(c.FormValue("Brand"), c.FormValue("Model"), c.FormValue("Description"), path)
		err = storage.Create(omafile)
		if err != nil {
			logger.Logger.Error(err)
		}
		c.Redirect(200, "/")
		return nil
	}
}

func searchHandler(storage *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {

		brand := c.QueryParam("brand")
		model := c.QueryParam("model")
		var files []models.Omafile = nil

		if brand != "" && model != "" {
			files = storage.GetByBrandAndModel(brand, model)
		} else if brand != "" {
			files = storage.GetByBrand(brand)
		} else if model != "" {
			files = storage.GetByModel(model)
		}

		if err := c.Render(http.StatusOK, "forms", files); err != nil {
			logger.Logger.Error(err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	}
}

func dowloadHandler(storage *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		logger.Logger.Println(id)
		oma := storage.GetById(id)

		c.Response().Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(oma.Directory))
		c.Response().Header().Set("Content-Type", "application/octet-stream")

		return c.File(oma.Directory)
	}
}
