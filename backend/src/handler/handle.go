package handler

import (
	"database/sql"
	"net/http"

	"NewGoResto/src/entities"
	myhttp "NewGoResto/src/http"

	// "github.com/Watsuk/go-food/src/entity"
	// myhttp "github.com/Watsuk/go-food/src/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func NewHandler(db *sql.DB, ref entities.Reference) *HandlerReference {

	handlers := &HandlerReference{
		chi.NewRouter(),
		ref.User,
		ref.Truck,
		ref.Chart,
		ref.Order,
		ref.OrderDetail,
	}

	// Cors middleware, the goal is to allow the front-end to access the API
	handlers.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Met l'URL de ton front-end ici
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	handlers.Use(middleware.Logger)

	handlers.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API Go-Food"))
	})

	handlers.Post("/truck", myhttp.CreateTrucksEndpoint(db))
	handlers.Get("/trucks", myhttp.GetAllTrucksEndpoint(db))
	handlers.Get("/truck/{truckID:[0-9]+}", myhttp.GetTruckByIDEndpoint(db))
	handlers.Delete("/truck/{truckID:[0-9]+}", myhttp.DeleteTruckEndpoint(db))
	handlers.Patch("/truck/{truckID:[0-9]+}", myhttp.UpdateTruckEndpoint(db))

	handlers.Post("/user", myhttp.CreateUserEndpoint(db))
	handlers.Get("/user/{userID:[0-9]+}", myhttp.GetUserByIdEndpoint(db))
	handlers.Delete("/user/{userID:[0-9]+}", myhttp.DeleteUserEndpoint(db))

	handlers.Post("/order_detail", myhttp.CreateOrderDetailsEndpoint(db))
	handlers.Get("/order_detail/order/{orderID:[0-9]+}", myhttp.GetOrderDetailsByOrderIDEndpoint(db))
	handlers.Get("/order_detail/truck/{truckID:[0-9]+}", myhttp.GetOrderDetailsByTruckIDEndpoint(db))
	handlers.Get("/order_detail/consumable/{consumableID:[0-9]+}", myhttp.GetOrderDetailsByConsumableIDEndpoint(db))
	handlers.Patch("/order_detail/order/{orderID:[0-9]+}", myhttp.UpdateOrderDetailEndpoint(db))
	handlers.Delete("/order_detail/order/{orderID:[0-9]+}/consumable/{consumableID:[0-9]+}", myhttp.DeleteOrderDetailEndpoint(db))

	handlers.Post("/order", myhttp.CreateOrderEndpoint(db))
	handlers.Get("/order/{orderID:[0-9]+}", myhttp.GetOrderByIDEndpoint(db))
	handlers.Patch("/order/{orderID:[0-9]+}", myhttp.UpdateOrderEndpoint(db))
	handlers.Delete("/order/{orderID:[0-9]+}", myhttp.DeleteOrderEndpoint(db))

	handlers.Post("/chart", myhttp.CreateChartEndpoint(db))
	handlers.Get("/chart/{chartID:[0-9]+}", myhttp.GetChartByIDEndpoint(db))
	handlers.Get("/charts/truck/{truckID:[0-9]+}", myhttp.GetChartsByTruckIDEndpoint(db))
	handlers.Patch("/chart/{chartID:[0-9]+}", myhttp.UpdateChartEndpoint(db))
	handlers.Delete("/chart/consumable/{consumableID:[0-9]+}", myhttp.DeleteChartEndpoint(db))

	return handlers
}

type HandlerReference struct {
	*chi.Mux
	user        *entities.User
	truck       *entities.Truck
	chart       *entities.Chart
	order       *entities.Order
	orderDetail *entities.OrderDetail
}
