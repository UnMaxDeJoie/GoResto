package Managers

import (
	"GoResto/entities"
	"database/sql"
	"fmt"
	"time"
)

func createOrder(Mydb *sql.DB, userID, truckID int, time time.Time) (int, error) {
	query := `INSERT INTO orders (user_id, truck_id, time) VALUES (?, ?, ?)`

	result, err := Mydb.Exec(query, userID, truckID, time)
	if err != nil {
		return 0, fmt.Errorf("createOrder: error when inserting new order: %v", err)
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createOrder: error getting last insert ID: %v", err)
	}

	return int(orderID), nil
}

func getOrderByID(Mydb *sql.DB, orderID int) (entities.Order, error) {
	var order entities.Order
	query := `SELECT order_id, user_id, truck_id, time FROM orders WHERE order_id = ?`

	err := Mydb.QueryRow(query, orderID).Scan(&order.OrderID, &order.UserID, &order.TruckID, &order.Time)
	if err != nil {
		return order, fmt.Errorf("getOrderByID: error when getting order with ID %d: %v", orderID, err)
	}

	return order, nil
}

func updateOrder(Mydb *sql.DB, orderID, userID, truckID int, time time.Time) error {
	query := `UPDATE orders SET user_id = ?, truck_id = ?, time = ? WHERE order_id = ?`

	result, err := Mydb.Exec(query, userID, truckID, time, orderID)
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

func deleteOrder(Mydb *sql.DB, orderID int) error {
	query := `DELETE FROM orders WHERE order_id = ?`

	result, err := Mydb.Exec(query, orderID)
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
