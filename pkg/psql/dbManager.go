package psql

import (
	"log"

	"github.com/AntonyCarl/OMA-Library/internal/domain"
)

func Create(o domain.Omafile) error {
	_, err := DbConn.Exec("INSERT INTO files (brand, model, info, directory) VALUES ($1, $2, $3, $4)",
		o.Brand, o.Model, o.Info, o.Directory)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetById(id string) domain.Omafile {
	rows, err := DbConn.Query("SELECT * FROM files WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	form := domain.Omafile{}
	for rows.Next() {
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			log.Fatal(err)
		}
	}
	return form
}

func GetByBrand(brand string) []domain.Omafile {
	rows, err := DbConn.Query("SELECT * FROM files WHERE brand = $1", brand)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			log.Fatal(err)
		}
		forms = append(forms, form)
	}
	return forms
}

func GetByBrandAndModel(brand string, model string) []domain.Omafile {
	rows, err := DbConn.Query("SELECT * FROM files WHERE brand = $1, model = $2", brand, model)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			log.Fatal(err)
		}
		forms = append(forms, form)
	}
	return forms
}

func GetByModel(model string) []domain.Omafile {
	rows, err := DbConn.Query("SELECT * FROM files WHERE model = $1", model)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	forms := make([]domain.Omafile, 0)
	for rows.Next() {
		form := domain.Omafile{}
		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
		if err != nil {
			log.Fatal(err)
		}
		forms = append(forms, form)
	}
	return forms
}
