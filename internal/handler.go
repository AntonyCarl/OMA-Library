package internal

import (
	"html/template"
	"log"
	"net/http"

	"github.com/AntonyCarl/OMA-Library/repository"
)

const (
	footer = "templates/header_footer/footer.html"
	header = "templates/header_footer/header.html"
)

func RunWeb() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/upload", uploadFormHandler)
	http.HandleFunc("/upload_file", uploadFileHandler)

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
	repository.SaveFile(file, handler.Filename)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
