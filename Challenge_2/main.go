package main

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	counter     = 0
	counterLock sync.Mutex
)

func main() {
	http.HandleFunc("/change", changeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	// Verrouillage du compteur pour assurer une incrémentation atomique
	counterLock.Lock()
	defer counterLock.Unlock()

	// Incrémentation du compteur à chaque visite
	counter++

	// Détermination si le compteur est pair ou impair
	isEven := counter%2 == 0

	// Définition du message en fonction de la parité du compteur
	message := "Le compteur est impair."
	if isEven {
		message = "Le compteur est pair."
	}

	// Chargement du modèle HTML
	tmpl, err := template.ParseFiles("templates/change.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Création des données à passer au modèle
	data := struct {
		Counter int
		Message string
		Even    bool
	}{
		Counter: counter,
		Message: message,
		Even:    isEven,
	}

	// Exécution du modèle avec les données
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
