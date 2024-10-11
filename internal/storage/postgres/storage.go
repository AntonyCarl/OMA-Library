package postgres

import (
	"database/sql"
	"fmt"

	"github.com/AntonyCarl/OMA-Library/internal/domain"
	"github.com/AntonyCarl/OMA-Library/pkg/logger"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	const op = "storage.NewStorage"

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode)) // get from config
	if err != nil {
		logger.Logger.Fatal(err)
		return nil, fmt.Errorf(op)
	}

	return &Storage{db: db}, nil
}

func (storage *Storage) Create(o domain.Omafile) error {
	_, err := storage.db.Exec("INSERT INTO files (brand, model, info, directory) VALUES ($1, $2, $3, $4)",
		o.Brand, o.Model, o.Info, o.Directory)

	if err != nil {
		logger.Logger.Error(err)
	}
	return err
}

func (storage *Storage) GetById(id string) domain.Omafile {
	rows, err := storage.db.Query("SELECT * FROM files WHERE id = $1", id)
	if err != nil {
		logger.Logger.Error(err)
	}
	defer rows.Close()

	form := domain.Omafile{}
	for rows.Next() {
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			logger.Logger.Error(err)
		}
	}
	return form
}

func (storage *Storage) GetByBrand(brand string) []domain.Omafile {
	rows, err := storage.db.Query("SELECT * FROM files WHERE brand = $1", brand)
	if err != nil {
		logger.Logger.Error(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			logger.Logger.Error(err)
		}
		forms = append(forms, form)
	}
	return forms
}

func (storage *Storage) GetByBrandAndModel(brand string, model string) []domain.Omafile {
	rows, err := storage.db.Query("SELECT * FROM files WHERE brand = $1, model = $2", brand, model)
	if err != nil {
		logger.Logger.Error(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			logger.Logger.Error(err)
		}
		forms = append(forms, form)
	}
	return forms
}

func (storage *Storage) GetByModel(model string) []domain.Omafile {
	rows, err := storage.db.Query("SELECT * FROM files WHERE model = $1", model)
	if err != nil {
		logger.Logger.Error(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			logger.Logger.Error(err)
		}
		forms = append(forms, form)
	}
	return forms
}
