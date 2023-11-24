package main

import (
	"net/http"

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
		{Title: "Home", Route: "/"},
		{Title: "App", Route: "/app"},
	}
)

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

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)
}
