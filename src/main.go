package main

import (
	"html/template"
	"log"
	"net/http"

	Managers "github.com/UnMaxDeJoie/GoResto/managers"
)

type Menu struct {
	Food  string
	Drink string
}

func main() {
	// Définition de la fonction handler pour la page principale
	goRestoHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Exemple d'utilisation de la fonction SendNotification
		serviceAccountKeyPath := "path/to/serviceAccountKey.json"
		deviceToken := "token_de_l_appareil"
		title := "Titre de la notification"
		body := "Corps de la notification"

		// Envoi de la notification avec les informations récupérées
		err := Managers.SendNotification(serviceAccountKeyPath, deviceToken, title, body)
		if err != nil {
			http.Error(w, "Erreur lors de l'envoi de la notification", http.StatusInternalServerError)
			log.Printf("Erreur lors de l'envoi de la notification: %v\n", err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// Définition du handler pour la page principale
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			menus := map[string][]Menu{
				"Menus": {
					{Food: "Burger", Drink: "Water"},
					{Food: "Potatoes", Drink: "Koka"},
					{Food: "Mochi", Drink: "Peace Ti"},
				},
			}
			tmpl.Execute(w, menus)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Définition du handler pour le bouton de notification
	http.HandleFunc("/send-notification", goRestoHandler)

	// Lancement du serveur HTTP
	log.Fatal(http.ListenAndServe(":3000", nil))
}
