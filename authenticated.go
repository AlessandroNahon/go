package main

import (
	"html/template"
	"net/http"
)

func Authenticated(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("authenticated").ParseFiles("authenticated.html", "header.html"))
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	data := RoutePageData{
		PageTitle: "App",
		Routes:    routes,
		Action:    "",
		Auth:      session.Values["authenticated"] == true,
	}
	tmpl.Execute(w, data)
}
