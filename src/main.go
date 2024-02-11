package main

import (
	"GoResto/managers"
	"html/template"
	"log"
	"net/http"
)

type Menu struct {
	Food  string
	Drink string
}

func main() {
	/*
		// Étape 1: Initialiser la connexion
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
		}

		fmt.Println(testQuery)
	*/

	goResto := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/template/index.html"))
		menus := map[string][]Menu{
			"Menus": {
				{Food: "Burger", Drink: "Water"},
				{Food: "Potatoes", Drink: "Koka"},
				{Food: "Mochi", Drink: "Peace Ti"},
			},
		}
		err := tmpl.Execute(w, menus)
		if err != nil {
			return
		}
	}
	http.HandleFunc("/", goResto)

	http.HandleFunc("/send-mail", managers.SendMailHandler)

	// termine le programme si le serveur ne peut pas se lancer sur le port 80
	log.Fatal(http.ListenAndServe(":8080", nil))
}
