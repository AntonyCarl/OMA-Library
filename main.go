package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal"
	"github.com/AntonyCarl/OMA-Library/pkg/psql"
	_ "github.com/lib/pq"
)

func main() {
	psql.DbConnection()
	internal.RunWeb()
	http.ListenAndServe(":8080", nil)

}

// func GetByBrand(brand string) ([]domain.Omafile, error) {
// 	db, err := psql.DbConnection()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM files WHERE brand = $1", brand)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	forms := make([]domain.Omafile, 0)
// 	for rows.Next() {
// 		form := domain.Omafile{}
// 		err := rows.Scan(&form.Id, &form.Brand, &form.Model, &form.Info, &form.Directory)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		forms = append(forms, form)
// 	}
// 	return forms, err
// }
