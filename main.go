package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Home)
	r.HandleFunc("/form", Form)
	r.HandleFunc("/app", Authenticated)
	r.HandleFunc("/logout", Logout)

	http.ListenAndServe(":80", r)
}
