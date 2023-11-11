package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   string
}

type Promotion struct {
	Nom       string
	Filiere   string
	Niveau    string
	Nombre    int
	Etudiants []Student
}

func main() {
	http.HandleFunc("/promo", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	promotion := Promotion{
		Nom:     "Promotion 2023",
		Filiere: "Informatique",
		Niveau:  "Bac+3",
		Nombre:  3,
		Etudiants: []Student{
			{"RODRIGUES", "Cyril", 22, "homme"},
			{"MEDERREG", "Kheir-eddine", 22, "homme"},
			{"PHILIPIERT", "Alan", 26, "homme"},
		},
	}

	tmpl, err := template.ParseFiles("templates/promo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
