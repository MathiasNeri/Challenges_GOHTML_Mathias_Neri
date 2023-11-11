package main

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	userData     UserData
	userDataLock sync.Mutex
)

// UserData représente les données d'utilisateur
type UserData struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Sexe          string
}

func main() {
	http.HandleFunc("/user/init", initHandler)
	http.HandleFunc("/user/treatment", treatmentHandler)
	http.HandleFunc("/user/display", displayHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/init.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func treatmentHandler(w http.ResponseWriter, r *http.Request) {
	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	dateNaissance := r.FormValue("date_naissance")
	sexe := r.FormValue("sexe")

	userDataLock.Lock()
	defer userDataLock.Unlock()

	userData = UserData{
		Nom:           nom,
		Prenom:        prenom,
		DateNaissance: dateNaissance,
		Sexe:          sexe,
	}

	http.Redirect(w, r, "/user/display", http.StatusSeeOther)
}

func displayHandler(w http.ResponseWriter, r *http.Request) {
	userDataLock.Lock()
	defer userDataLock.Unlock()

	tmpl, err := template.ParseFiles("templates/display.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, userData)
}
