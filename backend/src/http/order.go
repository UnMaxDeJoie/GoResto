package http

import (
	"NewGoResto/src/entities"
	"NewGoResto/src/managers"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order entities.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		orderID, err := managers.CreateOrder(db, order.UserID, order.TruckID, order.Time)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("Order created with ID %d", orderID)))
	}
}

func GetOrderByIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid Order ID", http.StatusBadRequest)
			return
		}
		order, err := managers.GetOrderByID(db, orderID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(order)
	}
}

func UpdateOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order entities.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = managers.UpdateOrder(db, order.OrderID, order.UserID, order.TruckID, order.Time)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order updated"))
	}
}

func DeleteOrderEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid Order ID", http.StatusBadRequest)
			return
		}
		err = managers.DeleteOrder(db, orderID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order deleted"))
	}
}
