package Managers

import (
	"fmt"
	"github.com/UnMaxDeJoie/GoResto/entities"
	"log"
)

func GetChartById(Mydb *DBController, Chartid int) (entities.Chart, error) {
	query := "SELECT truck_id, label, description , price FROM chart WHERE consumable_id= ?"

	var Chart entities.Chart

	err := Mydb.DB.QueryRow(query, Chartid).Scan(&Chart.TruckID, &Chart.Label, &Chart.Description, &Chart.Price)
	if err != nil {

		return entities.Chart{}, err
	}

	return Chart, nil
}

func GetTrucksChart(Mydb *DBController, Trucksid int) ([]entities.Chart, error) {
	query := "SELECT consumable_id, label, description, price FROM chart WHERE truck_id = ?"
	/*truck-id doit etre mis en cookie*/
	rows, err := Mydb.DB.Query(query, Trucksid)
	if err != nil {
		log.Printf("Erreur lors de l'exécution de la requête : %v", err)
		return nil, err
	}
	defer rows.Close()

	var charts []entities.Chart
	for rows.Next() {
		var chart entities.Chart
		if err := rows.Scan(&chart.ConsumableID, &chart.Label, &chart.Description, &chart.Price); err != nil {
			log.Printf("Erreur lors du scan d'une ligne : %v", err)
			return nil, err
		}
		charts = append(charts, chart)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur lors de l'itération sur les lignes : %v", err)
		return nil, err
	}

	return charts, nil
}

func deleteChart(Mydb *DBController, conusId int) {
	query := "DELETE FROM chart WHERE consumable_id= ?"

	result, err := Mydb.DB.Exec(query, conusId)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {

		log.Fatal(err)
	}

	if rowsAffected == 0 {

		log.Fatal(fmt.Errorf("aucun utilisateur trouvé avec l'ID %d", conusId))
	}
}

func createChart(Mydb *DBController, consumableID int, truckID int, label string, description string, price float64) error {

	query := `INSERT INTO chart (consumable_id, truck_id, label, description, price) VALUES (?, ?, ?, ?, ?)`

	_, err := Mydb.DB.Exec(query, consumableID, truckID, label, description, price)
	if err != nil {

		return fmt.Errorf("createChart: error when inserting new chart entry: %v", err)
	}

	return nil
}

func updateChart(Mydb *DBController, consumableID int, truckID int, label string, description string, price float64) error {

	query := `UPDATE chart SET truck_id = ?, label = ?, description = ?, price = ? WHERE consumable_id = ?`

	result, err := Mydb.DB.Exec(query, truckID, label, description, price, consumableID)
	if err != nil {

		return fmt.Errorf("updateChart: error when updating chart entry with consumableID %d: %v", consumableID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateChart: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {

		return fmt.Errorf("updateChart: no chart found with consumableID %d", consumableID)
	}

	return nil
}
