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

func CreateOrderDetailsEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderDetailModel entities.OrderDetail
		err := json.NewDecoder(r.Body).Decode(&orderDetailModel)
		if err != nil {
			fmt.Printf("Error when decoding order detail model: %v", err)
			http.Error(w, "Invalid order detail data", http.StatusBadRequest)
			return
		}
		err = managers.CreateOrderDetail(db, orderDetailModel.OrderID, orderDetailModel.ConsumableID, orderDetailModel.Quantity, orderDetailModel.Comment)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order detail created"))
	}
}

func GetOrderDetailsByOrderIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDString := chi.URLParam(r, "orderID")

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Order ID", http.StatusBadRequest)
			return
		}
		orderDetails, err := managers.GetOrderDetailsByOrderID(db, orderID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orderDetails)
	}
}

func GetOrderDetailsByTruckIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")

		truckID, err := strconv.Atoi(truckIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Truck ID", http.StatusBadRequest)
			return
		}
		orderDetails, err := managers.GetOrderDetailsByTruckID(db, truckID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orderDetails)
	}
}

func GetOrderDetailsByConsumableIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		consumableIDString := chi.URLParam(r, "consumableID")

		consumableID, err := strconv.Atoi(consumableIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Consumable ID", http.StatusBadRequest)
			return
		}
		orderDetail, err := managers.GetOrderDetailsByConsumableID(db, consumableID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orderDetail)
	}
}

func UpdateOrderDetailEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderDetailModel entities.OrderDetail
		err := json.NewDecoder(r.Body).Decode(&orderDetailModel)
		if err != nil {
			fmt.Printf("Error when decoding order detail model: %v", err)
			http.Error(w, "Invalid order detail data", http.StatusBadRequest)
			return
		}
		orderIDString := chi.URLParam(r, "orderID")

		orderID, err := strconv.Atoi(orderIDString)
		if err != nil {
			http.Error(w, "Invalid Order Detail ID", http.StatusBadRequest)
			return
		}
		err = managers.UpdateOrderDetail(db, orderID, orderDetailModel.ConsumableID, orderDetailModel.Quantity, orderDetailModel.Comment)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order detail updated"))
	}
}

func DeleteOrderDetailEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderDetailIDString := chi.URLParam(r, "orderDetailID")

		orderDetailID, err := strconv.Atoi(orderDetailIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Order Detail ID", http.StatusBadRequest)
			return
		}

		consumableString := chi.URLParam(r, "consumableID")
		consumableID, err := strconv.Atoi(consumableString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Consumable ID", http.StatusBadRequest)
			return
		}

		err = managers.DeleteOrderDetail(db, orderDetailID, consumableID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Order detail deleted"))
	}
}
