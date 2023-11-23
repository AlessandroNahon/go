package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Route struct {
	Title string
	Route string
}

type RoutePageData struct {
	PageTitle string
	Routes    []Route
	Action    string
	Auth      bool
}

var (
	key    = []byte("super-secret-key")
	store  = sessions.NewCookieStore(key)
	routes = []Route{
		{Title: "Login", Route: "/"},
		{Title: "App", Route: "/app"},
	}
)

func parseLayoutTemplate(f string, h bool) *template.Template {
	header := "./web/templates/header.html"
	name := strings.Split(filepath.Base(f), ".")[0]

	if h == false {
		header = ""
	}
	return template.Must(template.New(name).ParseFiles(f, header))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Login)
	r.HandleFunc("/app", Authenticated)
	r.HandleFunc("/logout", Logout)

	http.ListenAndServe(":80", r)
}
