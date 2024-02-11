package handler

import (
	"GoResto/managers"
	"html/template"
	"net/http"
)

func TruckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Obtenez les données des "trucks" de la base de données
	Mydb := managers.GetDBController()         // Assurez-vous que cette fonction retourne une connexion DB valide
	trucks, err := managers.GetAllTrucks(Mydb) // Supposons que cette fonction renvoie une liste de "trucks" et une erreur
	if err != nil {
		http.Error(w, "Impossible de récupérer les données des trucks", http.StatusInternalServerError)
		return
	}

	// Chargez le template HTML
	tmpl, err := template.ParseFiles("template/trucks.html") // Assurez-vous que le chemin est correct
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page des trucks", http.StatusInternalServerError)
		return
	}

	// Exécutez le template en passant les données des "trucks"
	err = tmpl.Execute(w, trucks) // Assurez-vous que votre template HTML est prêt à utiliser ces données
	if err != nil {
		http.Error(w, "Erreur lors du rendu de la page des trucks", http.StatusInternalServerError)
	}
}
