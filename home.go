package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := parseHeaderLayout("./web/templates/home.html", true)

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
