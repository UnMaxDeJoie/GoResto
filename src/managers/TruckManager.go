package Managers

import (
	"GoResto/entities"
	"database/sql"
	"fmt"
	"time"
)

func createTruck(Mydb *sql.DB, name string, slotBuffer uint8, opening, closing time.Time) error {
	query := `INSERT INTO trucks (name, slot_buffer, opening, closing) VALUES (?, ?, ?, ?)`

	_, err := Mydb.Exec(query, name, slotBuffer, opening, closing)
	if err != nil {
		return fmt.Errorf("createTruck: error when inserting new truck: %v", err)
	}

	return nil
}

func getTruckByID(Mydb *sql.DB, id int) (entities.Truck, error) {
	var truck entities.Truck
	query := `SELECT id, name, slot_buffer, opening, closing FROM trucks WHERE id = ?`

	err := Mydb.QueryRow(query, id).Scan(&truck.ID, &truck.Name, &truck.SlotBuffer, &truck.Opening, &truck.Closing)
	if err != nil {
		return truck, fmt.Errorf("getTruckByID: error when getting truck with ID %d: %v", id, err)
	}

	return truck, nil
}

func updateTruck(Mydb *sql.DB, id int, name string, slotBuffer uint8, opening, closing time.Time) error {
	query := `UPDATE trucks SET name = ?, slot_buffer = ?, opening = ?, closing = ? WHERE id = ?`

	result, err := Mydb.Exec(query, name, slotBuffer, opening, closing, id)
	if err != nil {
		return fmt.Errorf("updateTruck: error when updating truck with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateTruck: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("updateTruck: no truck found with ID %d", id)
	}

	return nil
}

func deleteTruck(Mydb *sql.DB, id int) error {
	query := `DELETE FROM trucks WHERE id = ?`

	result, err := Mydb.Exec(query, id)
	if err != nil {
		return fmt.Errorf("deleteTruck: error when deleting truck with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteTruck: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("deleteTruck: no truck found with ID %d", id)
	}

	return nil
}
