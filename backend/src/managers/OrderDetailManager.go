package managers

import (
	"NewGoResto/src/entities"
	"database/sql"
	"fmt"
)

func CreateOrderDetail(Mydb *sql.DB, orderID, consumableID, quantity int, comment string) error {
	query := `INSERT INTO order_detail (order_id, consumable_id, quantity, comment) VALUES (?, ?, ?, ?)`

	_, err := Mydb.Exec(query, orderID, consumableID, quantity, comment)
	if err != nil {
		return fmt.Errorf("createOrderDetail: error when inserting new order detail: %v", err)
	}

	return nil
}

func GetOrderDetailsByOrderID(Mydb *sql.DB, orderID int) ([]entities.OrderDetail, error) {
	query := `SELECT order_id, consumable_id, quantity, comment FROM order_detail WHERE order_id = ?`
	rows, err := Mydb.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("getOrderDetailsByOrderID: error when getting order details: %v", err)
	}
	defer rows.Close()

	var orderDetails []entities.OrderDetail
	for rows.Next() {
		var detail entities.OrderDetail
		if err := rows.Scan(&detail.OrderID, &detail.ConsumableID, &detail.Quantity, &detail.Comment); err != nil {
			return nil, fmt.Errorf("getOrderDetailsByOrderID: error scanning order detail: %v", err)
		}
		orderDetails = append(orderDetails, detail)
	}

	return orderDetails, nil
}

func GetOrderDetailsByTruckID(Mydb *sql.DB, truckID int) ([]entities.OrderDetail, error) {
	// Ajustez la requête pour sélectionner en fonction de TruckID
	query := `SELECT order_id, consumable_id, quantity, comment FROM order_detail WHERE truck_id = ?`
	rows, err := Mydb.Query(query, truckID)
	if err != nil {
		return nil, fmt.Errorf("getOrderDetailsByTruckID: error when getting order details: %v", err)
	}
	defer rows.Close()

	var orderDetails []entities.OrderDetail
	for rows.Next() {
		var detail entities.OrderDetail
		if err := rows.Scan(&detail.OrderID, &detail.ConsumableID, &detail.Quantity, &detail.Comment); err != nil {
			return nil, fmt.Errorf("getOrderDetailsByTruckID: error scanning order detail: %v", err)
		}
		orderDetails = append(orderDetails, detail)
	}

	return orderDetails, nil
}

func GetOrderDetailsByConsumableID(Mydb *sql.DB, consumableID int) ([]entities.OrderDetail, error) {
	// Ajustez la requête pour sélectionner en fonction de ConsumableID
	query := `SELECT order_id, consumable_id, quantity, comment FROM order_detail WHERE consumable_id = ?`
	rows, err := Mydb.Query(query, consumableID)
	if err != nil {
		return nil, fmt.Errorf("getOrderDetailsByConsumableID: error when getting order details: %v", err)
	}
	defer rows.Close()

	var orderDetails []entities.OrderDetail
	for rows.Next() {
		var detail entities.OrderDetail
		if err := rows.Scan(&detail.OrderID, &detail.ConsumableID, &detail.Quantity, &detail.Comment); err != nil {
			return nil, fmt.Errorf("getOrderDetailsByConsumableID: error scanning order detail: %v", err)
		}
		orderDetails = append(orderDetails, detail)
	}

	return orderDetails, nil
}

func UpdateOrderDetail(Mydb *sql.DB, orderID, consumableID, quantity int, comment string) error {
	query := `UPDATE order_detail SET quantity = ?, comment = ? WHERE order_id = ? AND consumableID = ?`

	result, err := Mydb.Exec(query, quantity, comment, orderID, consumableID)
	if err != nil {
		return fmt.Errorf("updateOrderDetail: error when updating order detail: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateOrderDetail: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("updateOrderDetail: no order detail found with OrderID %d and ConsumableID %d", orderID, consumableID)
	}

	return nil
}

func DeleteOrderDetail(Mydb *sql.DB, orderID, consumableID int) error {
	query := `DELETE FROM order_detail WHERE order_id = ? AND consumable_id = ?`

	result, err := Mydb.Exec(query, orderID, consumableID)
	if err != nil {
		return fmt.Errorf("deleteOrderDetail: error when deleting order detail: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteOrderDetail: error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("deleteOrderDetail: no order detail found with OrderID %d and ConsumableID %d", orderID, consumableID)
	}

	return nil
}
