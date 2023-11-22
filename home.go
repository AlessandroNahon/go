package main

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").ParseFiles("home.html", "header.html"))

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
		PageTitle: "Home",
		Routes:    routes,
		Action:    "Log In",
		Auth:      isAuthenticated,
	}
	tmpl.Execute(w, data)
}
