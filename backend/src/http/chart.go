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

func CreateChartEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chart entities.Chart
		err := json.NewDecoder(r.Body).Decode(&chart)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = managers.CreateChart(db, chart.ConsumableID, chart.TruckID, chart.Label, chart.Description, chart.Price)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Chart created"))
	}
}

func GetChartByIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chartIDString := chi.URLParam(r, "chartID")

		chartID, err := strconv.Atoi(chartIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Chart ID", http.StatusBadRequest)
			return
		}
		chart, err := managers.GetChartById(db, chartID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(chart)
	}
}

func GetChartsByTruckIDEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		truckIDString := chi.URLParam(r, "truckID")

		truckID, err := strconv.Atoi(truckIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Truck ID", http.StatusBadRequest)
			return
		}
		charts, err := managers.GetTrucksChart(db, truckID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(charts)
	}
}

func UpdateChartEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var chart entities.Chart
		err := json.NewDecoder(r.Body).Decode(&chart)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = managers.UpdateChart(db, chart.ConsumableID, chart.TruckID, chart.Label, chart.Description, chart.Price)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Chart updated"))
	}
}

func DeleteChartEndpoint(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		consumableIDString := chi.URLParam(r, "consumableID")

		consumableID, err := strconv.Atoi(consumableIDString)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid Chart ID", http.StatusBadRequest)
			return
		}

		managers.DeleteChart(db, consumableID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Chart deleted"))
	}
}
