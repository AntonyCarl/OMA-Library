package psql

import "log"

func Create(o Omafile) error {
	db, err := database.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO files (brand, model, info, directory) VALUES (?, ?, ?, ?)",
		o.Brand, o.Model, o.Info, o.Directory)
	return err
}
