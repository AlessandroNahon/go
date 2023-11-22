package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func Form(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form").ParseFiles("form.html", "header.html"))
	data := RoutePageData{
		PageTitle: "Contact",
		Routes:    routes,
		Action:    "",
		Auth:      false,
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, data)
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	_ = details

	tmpl.Execute(w, nil)
}
