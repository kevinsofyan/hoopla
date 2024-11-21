package handler

import (
	"database/sql"
	"fmt"
)

type PaymentHandler interface {
	ShowProductTable() error
	ProcessOrder(userID, productID, quantity int, totalAmount float64) error
	ShowOrderTable() error
}

type PaymentHandlerImpl struct {
	DB *sql.DB
}

func NewPaymentHandler(db *sql.DB) *PaymentHandlerImpl {
	return &PaymentHandlerImpl{
		DB: db,
	}
}

func (ph *PaymentHandlerImpl) ShowProductTable() error {
	rows, err := ph.DB.Query("SELECT ProductID, ProductName, Price, StockQuantity FROM Product")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Product Table:")
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	for rows.Next() {
		var id int
		var name string
		var price float64
		var stock int
		err := rows.Scan(&id, &name, &price, &stock)
		if err != nil {
			return err
		}
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", id, name, price, stock)
	}
	return nil
}

func (ph *PaymentHandlerImpl) ProcessOrder(userID, productID, quantity int, totalAmount float64) error {
	// Check if UserID exists
	var exists bool
	err := ph.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE UserID = ?)", userID).Scan(&exists)
	if err != nil {

		return err
	}
	if !exists {
		return fmt.Errorf("UserID %d does not exist", userID)
	}

	// Start a transaction
	tx, err := ph.DB.Begin()
	if err != nil {
		return err
	}

	// Insert into Order table
	orderResult, err := tx.Exec("INSERT INTO `Order` (UserID, OrderDate, TotalAmount) VALUES (?, NOW(), ?)", userID, totalAmount)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Get the last inserted OrderID
	orderID, err := orderResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into OrderDetails table
	_, err = tx.Exec("INSERT INTO OrderDetails (OrderID, ProductID, Quantity, UnitPrice) VALUES (?, ?, ?, ?)", orderID, productID, quantity, totalAmount/float64(quantity))
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Order processed successfully for User %d, Product %d, Quantity %d, Total Amount %.2f\n", userID, productID, quantity, totalAmount)
	return nil
}

func (ph *PaymentHandlerImpl) ShowOrderTable() error {
	rows, err := ph.DB.Query("SELECT OrderDetailID, OrderID, ProductID, Quantity, UnitPrice FROM OrderDetails")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Order Details Table:")
	fmt.Printf("%-15s %-10s %-10s %-10s %-10s\n", "OrderDetailID", "OrderID", "ProductID", "Quantity", "UnitPrice")
	for rows.Next() {
		var orderDetailID, orderID, productID, quantity int
		var unitPrice float64
		err := rows.Scan(&orderDetailID, &orderID, &productID, &quantity, &unitPrice)
		if err != nil {
			return err
		}
		fmt.Printf("%-15d %-10d %-10d %-10d %-10.2f\n", orderDetailID, orderID, productID, quantity, unitPrice)
	}
	return nil
}
