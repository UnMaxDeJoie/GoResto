package managers

import (
	"NewGoResto/src/entities"
	"database/sql"
	"fmt"
	"log"
)

func CreateTruck(Mydb *sql.DB, name string, slotBuffer uint8, opening string, closing string) error {
	query := `INSERT INTO trucks (name, slot_buffer, opening, closing) VALUES (?, ?, ?, ?)`

	_, err := Mydb.Exec(query, name, slotBuffer, opening, closing)
	if err != nil {
		return fmt.Errorf("createTruck: error when inserting new truck: %v", err)
	}

	return nil
}
func GetAllTrucks(db *sql.DB) ([]entities.Truck, error) {
	var trucks []entities.Truck

	// Remplacez 'your_trucks_table' par le nom réel de votre table de trucks
	query := "SELECT id, name, slot_buffer, opening, closing FROM trucks"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Erreur lors de la requête pour obtenir tous les trucks: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var truck entities.Truck
		// Assurez-vous d'adapter le Scan en fonction de la structure de votre table
		if err := rows.Scan(&truck.ID, &truck.Name, &truck.SlotBuffer, &truck.Opening, &truck.Closing); err != nil {
			log.Printf("Erreur lors du scan d'un truck: %v", err)
			return nil, err
		}
		trucks = append(trucks, truck)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur lors de l'itération sur les résultats des trucks: %v", err)
		return nil, err
	}

	return trucks, nil
}

func GetTruckByID(Mydb *sql.DB, id int) (entities.Truck, error) {
	var truck entities.Truck
	query := `SELECT id, name, slot_buffer, opening, closing FROM trucks WHERE id = ?`

	err := Mydb.QueryRow(query, id).Scan(&truck.ID, &truck.Name, &truck.SlotBuffer, &truck.Opening, &truck.Closing)
	if err != nil {
		return truck, fmt.Errorf("getTruckByID: error when getting truck with ID %d: %v", id, err)
	}

	return truck, nil
}

func UpdateTruck(Mydb *sql.DB, id int, name string, slotBuffer uint8, opening string, closing string) error {
	query := `UPDATE trucks SET name = ?, slot_buffer = ?, opening = ?, closing = ? WHERE id = ?`

	result, err := Mydb.Exec(query, name, slotBuffer, opening, closing, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("updateTruck: error when updating truck with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("updateTruck: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {

		return fmt.Errorf("updateTruck: no truck found with ID %d", id)
	}

	return nil
}

func DeleteTruck(Mydb *sql.DB, id int) error {
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