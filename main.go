package main

import (
	"net/http"

	"github.com/AntonyCarl/OMA-Library/internal"
)

func main() {

	internal.RunWeb()
	http.ListenAndServe(":8080", nil)
}
