package handler

import (
	"GoResto/managers"
	"encoding/json"
	"net/http"
)

func TrucksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		// Si la méthode de la requête n'est pas GET, renvoyez une erreur 405 Method Not Allowed
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	Mydb := managers.GetDBController()

	// Utilisez la fonction GetAllTrucks de votre package Managers pour récupérer les trucks
	trucks, err := managers.GetAllTrucks(Mydb)
	if err != nil {
		http.Error(w, "Impossible de récupérer les trucks", http.StatusInternalServerError)
		return
	}

	// Définir le Content-Type de la réponse à application/json
	w.Header().Set("Content-Type", "application/json")

	// Encoder les trucks en JSON et les écrire dans la réponse
	if err := json.NewEncoder(w).Encode(trucks); err != nil {
		http.Error(w, "Erreur lors de l'encodage des trucks en JSON", http.StatusInternalServerError)
	}
}
