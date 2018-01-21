package wineMap

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/main.html"))
	tmpl.Execute(w, nil)
}
