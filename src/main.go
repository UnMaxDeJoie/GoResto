package main

import (
	"html/template"
	"log"
	"net/http"
)

type Menu struct {
	Food  string
	Drink string
}

func main() {
	goResto := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		menus := map[string][]Menu{
			"Menus": {
				{Food: "Burger", Drink: "Water"},
				{Food: "Potatoes", Drink: "Koka"},
				{Food: "Mochi", Drink: "Peace Ti"},
			},
		}
		tmpl.Execute(w, menus)
	}
	http.HandleFunc("/", goResto)

	// termine le programme si le serveur ne peut pas se lancer sur le port 80
	log.Fatal(http.ListenAndServe(":3000", nil))
}
