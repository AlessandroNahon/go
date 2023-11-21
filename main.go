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

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Home)
	r.HandleFunc("/form", Form)
	r.HandleFunc("/secret", Secret)
	r.HandleFunc("/logout", logout)

	http.ListenAndServe(":80", r)
}
