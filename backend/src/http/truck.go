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

func CreateTrucksEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var truckModel entities.Truck
		err := json.NewDecoder(r.Body).Decode(&truckModel)
		if err != nil {
			fmt.Printf("Error when decoding truck model: %v", err)
			http.Error(w, "Invalid truck data", http.StatusBadRequest)
			return
		}
		err = managers.CreateTruck(db, truckModel.Name, truckModel.SlotBuffer, truckModel.Opening, truckModel.Closing,truckModel.UserID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Truck created"))
	}
}

func GetAllTrucksEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		trucks, err := managers.GetAllTrucks(db)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(trucks)
	}
}

func GetTruckByIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")

		truckID, err := strconv.Atoi(truckIDString)
		if err != nil {
			http.Error(w, "Invalid Truck ID", http.StatusBadRequest)
			return
		}

		truck, err := managers.GetTruckByID(db, truckID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(truck)
	}
}

func DeleteTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")

		truckID, err := strconv.Atoi(truckIDString)
		if err != nil {
			http.Error(w, "Invalid Truck ID", http.StatusBadRequest)
			return
		}

		err = managers.DeleteTruck(db, truckID)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Truck deleted"))
	}
}

func UpdateTruckEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var truckModel entities.Truck
		err := json.NewDecoder(r.Body).Decode(&truckModel)
		if err != nil {
			fmt.Printf("Error when decoding truck model: %v", err)
			http.Error(w, "Invalid truck data", http.StatusBadRequest)
			return
		}
		err = managers.UpdateTruck(db, truckModel.ID, truckModel.Name, truckModel.SlotBuffer, truckModel.Opening, truckModel.Closing ,truckModel.UserID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Truck updated"))
	}
}
