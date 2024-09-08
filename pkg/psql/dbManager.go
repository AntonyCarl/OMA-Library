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

func GetByBrand(brand string) ([]domain.Omafile, error) {
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
	return forms, err
}
