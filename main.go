package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("layout.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", r)
}
