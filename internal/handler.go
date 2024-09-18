package internal

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AntonyCarl/OMA-Library/internal/domain"
	psql "github.com/AntonyCarl/OMA-Library/pkg/psql"
	"github.com/AntonyCarl/OMA-Library/repository"
	"github.com/gorilla/mux"
)

const (
	footer = "templates/header_footer/footer.html"
	header = "templates/header_footer/header.html"
	forms  = "templates/forms.html"
)

func RunWeb() {
	router := mux.NewRouter()
	router.HandleFunc("/", mainPageHandler).Methods("GET")
	router.HandleFunc("/upload", uploadFormHandler).Methods("GET")
	router.HandleFunc("/upload_file", uploadFileHandler).Methods("POST")
	router.HandleFunc("/search", searchHandler).Methods("GET")
	router.HandleFunc("/oma/{id:[0-9]+}", dowloadHandler)

	http.Handle("/", router)

}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", footer, header)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func uploadFormHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/upload_form.html", footer, header)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "upload", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("uploaded_file")
	if err != nil {
		log.Fatal(err)
	}
	path := repository.SaveFile(file, handler.Filename)
	omafile := domain.NewOmafile(r.FormValue("Brand"), r.FormValue("Model"), r.FormValue("Description"), path)

	err = psql.Create(omafile)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", footer, header, forms)
	if err != nil {
		log.Fatal(err)
	}

	brand := r.URL.Query().Get("brand")
	model := r.URL.Query().Get("model")
	var files []domain.Omafile = nil

	if brand != "" && model != "" {
		files = psql.GetByBrandAndModel(brand, model)
	} else if brand != "" {
		files = psql.GetByBrand(brand)
	} else if model != "" {
		files = psql.GetByModel(model)
	}

	t.ExecuteTemplate(w, "forms", files)
}

func dowloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	oma := psql.GetById(vars["id"])

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(oma.Directory))
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, oma.Directory)
}
