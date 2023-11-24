package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Login(w http.ResponseWriter, r *http.Request) {
	paths := []string{
		filepath.Join("./templates/", "meta.html"),
		filepath.Join("./templates/", "header.html"),
		filepath.Join("./templates/", "login.html"),
	}
	tmpl := template.Must(template.New("login").ParseFiles(paths...))

	session, _ := store.Get(r, "cookie-name")
	isAuthenticated := session.Values["authenticated"] == true

	if r.Method == http.MethodPost {
		if isAuthenticated {
			http.Redirect(w, r, "/app", http.StatusFound)
		} else {
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/app", http.StatusFound)
		}
	}

	data := RoutePageData{
		PageTitle: "Welcome!",
		Routes:    routes,
		Action:    "Log In",
		Auth:      isAuthenticated,
	}

	tmpl.Execute(w, data)
}
