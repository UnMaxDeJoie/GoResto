package main

import (
	"GoResto/handler"
	"log"
	"net/http"
)

type Menu struct {
	Food  string
	Drink string
}

func main() {
	/*// Étape 1: Initialiser la connexion
	managers.NewDBController()

	// Étape 2: Récupérer l'instance de connexion
	truc := managers.GetDBController()
	if truc == nil {
		log.Fatal("Impossible d'établir une connexion à la base de données.")
	}

	// Étape 3: Exécuter une requête de test
	var testQuery string
	err := truc.QueryRow("SELECT 'Connexion réussie'").Scan(&testQuery)
	if err != nil {
		log.Fatalf("Échec de la requête de test : %v", err)
	}*/

	/*fmt.Println(testQuery)*/

	// Définir le handler de login
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/truck", handler.TruckHandler)

	// Démarrer le serveur HTTP
	log.Println("Le serveur écoute sur le port :3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur :", err)
	}
}
