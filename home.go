package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Route string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
	Action    string
	Auth      bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))

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

	actionText := "Log In"

	if isAuthenticated {
		actionText = "App"
	}

	data := TodoPageData{
		PageTitle: "Home",
		Todos: []Todo{
			{Title: "Home", Route: "/"},
			{Title: "App", Route: "/app"},
			{Title: "Form", Route: "/form"},
		},
		Action: actionText,
		Auth:   isAuthenticated,
	}
	tmpl.Execute(w, data)
}
