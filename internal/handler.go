package internal

import (
	"html/template"
	"log"
	"net/http"
)

const (
	footer = "templates/header_footer/footer.html"
	header = "templates/header_footer/header.html"
)

func RunWeb() {
	http.HandleFunc("/", showMainPage)
	http.HandleFunc("/upload", showUploadForm)

}

func showMainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", footer, header)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func showUploadForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/upload_form.html", footer, header)
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "upload", nil)
	if err != nil {
		log.Fatal(err)
	}
}
