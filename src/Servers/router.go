package Servers

/*
import (
	"fmt"
	"github.com/UnMaxDeJoie/GoResto/managers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func router() {
	r := chi.NewRouter()

	// Login
	r.Route("/", func(r chi.Router) {
		r.Post("/login", managers.Login)
		r.Post("/createCompte", CreateCompte)
	})

	// Client
	r.Route("/client", func(r chi.Router) {
		r.Get("/home", GetTrucks)
		r.Get("/{idTruck}", GetTruck)
		r.Post("/{idTruck}/newOrder", ControllersBack.CreateOrder)
	})

	// Restaurant
	r.Route("/Restaurant", func(r chi.Router) {
		r.Get("/allOrders", ControllerFront.GetOrders)
		r.Post("/updateTruck", ControllersBack.UpdateTruck)
	})

	// Admin
	r.Route("/admin", func(r chi.Router) {
		r.Get("/", ControllerFront.GetTrucks)
		r.Post("/newTruck", ControllersBack.CreateTruck)
		r.Delete("/admin/{idTruck}", ControllersBack.DeleteTruck)
	})

	err := http.ListenAndServe("localhost:3000", r)
	if err != nil {
		fmt.Println("Could not start the server", err)
}
*/
