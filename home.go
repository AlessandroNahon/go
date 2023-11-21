package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
	Action    string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("home.html"))

	session, _ := store.Get(r, "cookie-name")

	if r.Method == http.MethodPost {
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/secret", http.StatusFound)
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
		Action: "Log In",
	}
	tmpl.Execute(w, data)
}
