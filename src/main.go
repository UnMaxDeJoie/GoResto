package main

import (
	Managers "GoResto/Managers"
	"fmt"
	"log"
)

type Menu struct {
	Food  string
	Drink string
}

func main() {
	// Étape 1: Initialiser la connexion
	Managers.NewDBController()

	// Étape 2: Récupérer l'instance de connexion
	truc := Managers.GetDBController()
	if truc == nil {
		log.Fatal("Impossible d'établir une connexion à la base de données.")
	}

	// Étape 3: Exécuter une requête de test
	var testQuery string
	err := truc.QueryRow("SELECT 'Connexion réussie'").Scan(&testQuery)
	if err != nil {
		log.Fatalf("Échec de la requête de test : %v", err)
	}

	fmt.Println(testQuery)
}
