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
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	promotion := Promotion{
		Nom:     "Promotion 2023",
		Filiere: "Informatique",
		Niveau:  "Bac+3",
		Nombre:  5,
		Etudiants: []Student{
			{"Doe", "John", 23, "homme"},
			{"Smith", "Jane", 21, "femme"},
			{"Johnson", "Bob", 23, "homme"},
			{"Williams", "Alice", 22, "femme"},
			{"Brown", "Charlie", 24, "homme"},
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
