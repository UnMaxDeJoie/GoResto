package handler

import (
	"GoResto/managers"
	"fmt"
	"log"

	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Servir la page de login pour les requêtes GET
		http.ServeFile(w, r, "template/login.html") // Assurez-vous que le chemin est correct
	case "POST":
		// La logique pour traiter la soumission du formulaire de login
		Mydb := managers.GetDBController()
		fmt.Println("Traitement de la soumission du formulaire de login")

		// Parse le corps de la requête pour extraire les données du formulaire
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de l'analyse des données du formulaire", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		uid, _, err := managers.Login(Mydb, password, email) // Assurez-vous que cette fonction existe dans votre package "managers"
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		tokenString, err := managers.CreateToken(email, uid) // Assurez-vous que cette fonction existe dans votre package "managers"
		if err != nil {
			log.Printf("Erreur lors de la création du token : %v", err)
			http.Error(w, "Erreur serveur interne", http.StatusInternalServerError)
			return
		}

		// Définir le token dans un cookie et l'envoyer au client
		http.SetCookie(w, &http.Cookie{
			Name:     "auth-token",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true, // HttpOnly pour augmenter la sécurité en empêchant l'accès via JS
			Path:     "/",  // Le cookie est valide pour toute l'application
		})

		// Envoyez une réponse pour indiquer que le login a réussi
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login réussi et token envoyé dans un cookie."))
	default:
		// Méthode HTTP non prise en charge
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
