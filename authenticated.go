package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Authenticated(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		filepath.Join("./templates/", "header.html"),
		filepath.Join("./templates/", "authenticated.html"),
	}
	tmpl := template.Must(template.New("authenticated").ParseFiles(paths...))
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	data := RoutePageData{
		PageTitle: "App",
		Routes:    routes,
		Action:    "",
		Auth:      session.Values["authenticated"].(bool),
	}
	tmpl.Execute(w, data)
}
