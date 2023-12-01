package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

func List(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	if (title != "") || (director != "") {
		tmpl := template.Must(template.ParseFiles(filepath.Join("./templates/", "authenticated.html")))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}
}
