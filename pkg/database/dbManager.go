package psql

import (
	"log"

	"github.com/AntonyCarl/OMA-Library/internal/domain"
)

func Create(o domain.Omafile) error {
	db, err := DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO files (brand, model, info, directory) VALUES ($1, $2, $3, $4)",
		o.Brand, o.Model, o.Info, o.Directory)
	return err
}
