package main

import (
	"net/http"
)

func Authenticated(w http.ResponseWriter, r *http.Request) {
	tmpl := parseLayoutTemplate("./web/templates/authenticated.html", true)
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
