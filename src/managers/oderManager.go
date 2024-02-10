package Managers

import (
	"fmt"
	"github.com/UnMaxDeJoie/GoResto/entities"
	"time"
)

func createOrder(Mydb *DBController, userID, truckID int, time time.Time) (int, error) {
	query := `INSERT INTO orders (user_id, truck_id, time) VALUES (?, ?, ?)`

	result, err := Mydb.DB.Exec(query, userID, truckID, time)
	if err != nil {
		return 0, fmt.Errorf("createOrder: error when inserting new order: %v", err)
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createOrder: error getting last insert ID: %v", err)
	}

	return int(orderID), nil
}

func getOrderByID(Mydb *DBController, orderID int) (entities.Order, error) {
	var order entities.Order
	query := `SELECT order_id, user_id, truck_id, time FROM orders WHERE order_id = ?`

	err := Mydb.DB.QueryRow(query, orderID).Scan(&order.OrderID, &order.UserID, &order.TruckID, &order.Time)
	if err != nil {
		return order, fmt.Errorf("getOrderByID: error when getting order with ID %d: %v", orderID, err)
	}

	return order, nil
}

func updateOrder(Mydb *DBController, orderID, userID, truckID int, time time.Time) error {
	query := `UPDATE orders SET user_id = ?, truck_id = ?, time = ? WHERE order_id = ?`

	result, err := Mydb.DB.Exec(query, userID, truckID, time, orderID)
	if err != nil {
		return fmt.Errorf("updateOrder: error when updating order with ID %d: %v", orderID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateOrder: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("updateOrder: no order found with ID %d", orderID)
	}

	return nil
}

func deleteOrder(Mydb *DBController, orderID int) error {
	query := `DELETE FROM orders WHERE order_id = ?`

	result, err := Mydb.DB.Exec(query, orderID)
	if err != nil {
		return fmt.Errorf("deleteOrder: error when deleting order with ID %d: %v", orderID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteOrder: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("deleteOrder: no order found with ID %d", orderID)
	}

	return nil
}
