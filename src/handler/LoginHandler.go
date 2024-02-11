package handler

import (
	"GoResto/managers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	Mydb := managers.GetDBController()
	fmt.Printf("je suisn ici")
	// Parse le corps de la requête pour extraire les données du formulaire
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse des données du formulaire", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Utilisez la fonction login que vous avez fournie pour authentifier l'utilisateur
	uid, _, err := managers.Login(Mydb, password, email) // Remplacez GetDBController() par votre instance *sql.DB si nécessaire
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Utilisez la fonction createToken pour générer le token JWT
	tokenString, err := managers.CreateToken(email, uid)
	if err != nil {
		log.Printf("Erreur lors de la création du token : %v", err)
		http.Error(w, "Erreur serveur interne", http.StatusInternalServerError)
		return
	}

	// Définissez le token dans un cookie et envoyez-le dans la réponse
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
}
