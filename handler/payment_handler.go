package handler

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
)

type PaymentHandler interface {
	ShowProductTable() error
	ProcessOrder(userID, productID, quantity int, totalAmount float64) error
	ShowOrderTable() error
	ShowPaymentTable() error
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
	rows, err := ph.DB.Query("SELECT ProductID, ProductName, CategoryID, Price, StockQuantity FROM Product")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Product Table:")
	fmt.Printf("%-10s %-30s %-10s %-10s %-10s\n", "ID", "Name", "CategoryID", "Price", "Stock")
	for rows.Next() {
		var id, categoryID, stockQuantity int
		var name string
		var price float64
		err := rows.Scan(&id, &name, &categoryID, &price, &stockQuantity)
		if err != nil {
			return err
		}
		fmt.Printf("%-10d %-30s %-10d %-20.2f %-10d\n", id, name, categoryID, price, stockQuantity)
	}
	return nil
}

func (ph *PaymentHandlerImpl) ProcessOrder(userID, productID, quantity int, totalAmount float64) error {
	// Check if the user exists
	var exists bool
	if err := ph.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE UserID = ?)", userID).Scan(&exists); err != nil || !exists {
		return fmt.Errorf("user %d does not exist", userID)
	}

	// Start transaction
	tx, err := ph.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Create the order
	orderResult, err := tx.Exec("INSERT INTO `Orders` (UserID, OrderDate, TotalAmount) VALUES (?, NOW(), 0)", userID)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	orderID, err := orderResult.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to fetch order ID: %v", err)
	}

	// Process the provided product and quantity
	var unitPrice float64
	if err := ph.DB.QueryRow("SELECT Price FROM Product WHERE ProductID = ?", productID).Scan(&unitPrice); err != nil {
		return fmt.Errorf("invalid product %d: %v", productID, err)
	}

	if _, err := tx.Exec(
		"INSERT INTO OrderDetails (OrderID, ProductID, Quantity, UnitPrice) VALUES (?, ?, ?, ?)",
		orderID, productID, quantity, unitPrice,
	); err != nil {
		return fmt.Errorf("failed to add order detail: %v", err)
	}

	totalAmount += float64(quantity) * unitPrice

	// Check if more items need to be added
	for {
		var addMore string
		fmt.Print("Add another item? (yes/no): ")
		fmt.Scan(&addMore)
		if addMore != "yes" {
			break
		}

		fmt.Print("Enter ProductID: ")
		fmt.Scan(&productID)
		fmt.Print("Enter Quantity: ")
		fmt.Scan(&quantity)

		// Fetch product price for the new item
		if err := ph.DB.QueryRow("SELECT Price FROM Product WHERE ProductID = ?", productID).Scan(&unitPrice); err != nil {
			return fmt.Errorf("invalid product %d: %v", productID, err)
		}

		if _, err := tx.Exec(
			"INSERT INTO OrderDetails (OrderID, ProductID, Quantity, UnitPrice) VALUES (?, ?, ?, ?)",
			orderID, productID, quantity, unitPrice,
		); err != nil {
			return fmt.Errorf("failed to add order detail: %v", err)
		}

		totalAmount += float64(quantity) * unitPrice
	}

	// Update order total
	if totalAmount == 0 {
		fmt.Println("No items added. Order canceled.")
		return nil
	}

	if _, err := tx.Exec("UPDATE `Orders` SET TotalAmount = ? WHERE OrderID = ?", totalAmount, orderID); err != nil {
		return fmt.Errorf("failed to update total: %v", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Printf("Order processed. OrderID: %d, Total: Rp %.2f\n", orderID, totalAmount)
	return nil
}

func (ph *PaymentHandlerImpl) ShowOrderTable() error {
	// Query to show the Orders table
	orderRows, err := ph.DB.Query(`
        SELECT o.OrderID, u.Name, o.OrderDate, o.TotalAmount, o.Paid
        FROM Orders o
        JOIN User u ON o.UserID = u.UserID
    `)
	if err != nil {
		return err
	}
	defer orderRows.Close()

	fmt.Println("Orders Table:\n")
	fmt.Printf("%-10s %-20s %-20s %-10s\n", "OrderID", "UserName", "OrderDate", "Total Amount", "Paid")
	for orderRows.Next() {
		var orderID int
		var userName string
		var orderDate string
		var totalAmount float64
		var paid bool
		err := orderRows.Scan(&orderID, &userName, &orderDate, &totalAmount, &paid)
		if err != nil {
			return err
		}
		fmt.Printf("%-10d %-20s %-20s Rp %-15.2f %-10t\n", orderID, userName, orderDate, totalAmount, paid)
	}

	// Prompt the user to enter an OrderID
	var selectedOrderID int
	fmt.Print("\nEnter OrderID to view details: ")
	_, err = fmt.Scan(&selectedOrderID)
	if err != nil {
		return err
	}

	// Query to show the OrderDetails table for the selected OrderID
	rows, err := ph.DB.Query(`
        SELECT od.OrderDetailID, p.ProductName, od.Quantity, od.UnitPrice, (od.Quantity * od.UnitPrice) as Total
        FROM OrderDetails od
        JOIN Product p ON od.ProductID = p.ProductID
        WHERE od.OrderID = ?
    `, selectedOrderID)
	if err != nil {
		return err
	}
	defer rows.Close()
	clearScreen()
	fmt.Println("\nOrder Details: \n")
	fmt.Printf("%-15s %-30s %-10s %-10s %-10s\n", "OrderDetailID", "ProductName", "Quantity", "UnitPrice", "Total")
	for rows.Next() {
		var orderDetailID, quantity int
		var productName string
		var unitPrice, total float64
		err := rows.Scan(&orderDetailID, &productName, &quantity, &unitPrice, &total)
		if err != nil {
			return err
		}
		fmt.Printf("%-15d %-30s %-10d Rp %-10.2f Rp %-10.2f\n", orderDetailID, productName, quantity, unitPrice, total)
	}

	// Prompt the user to enter an OrderDetailID
	var selectedOrderDetailID int
	fmt.Print("\nEnter OrderDetailID to update or delete: ")
	_, err = fmt.Scan(&selectedOrderDetailID)
	if err != nil {
		return err
	}

	// Prompt the user to choose an action
	var action string
	fmt.Print("Do you want to update quantity or delete the item? (update/delete): ")
	fmt.Scan(&action)

	if action == "update" {
		// Prompt the user to enter the new quantity
		var newQuantity int
		fmt.Print("Enter new quantity: ")
		fmt.Scan(&newQuantity)

		err := ph.updateOrderQuantity(selectedOrderDetailID, newQuantity)
		if err != nil {
			return err
		}
		fmt.Printf("\nOrder detail updated successfully.\n\n")
	} else if action == "delete" {
		err := ph.deleteOrderDetail(selectedOrderDetailID)
		if err != nil {
			return err
		}
		fmt.Printf("\nOrder detail deleted successfully.\n\n")
	} else {
		fmt.Println("Invalid action.")
	}

	// Calculate the grand total for the selected OrderID
	var grandTotal float64
	err = ph.DB.QueryRow(`
        SELECT COALESCE(SUM(od.Quantity * od.UnitPrice), 0) as GrandTotal
        FROM OrderDetails od
        WHERE od.OrderID = ?
    `, selectedOrderID).Scan(&grandTotal)
	if err != nil {
		return err
	}

	fmt.Printf("\nGrand Total: Rp %.2f\n", grandTotal)
	return nil
}

func (ph *PaymentHandlerImpl) ShowPaymentTable() error {
	// Show all unpaid orders
	rows, err := ph.DB.Query(`
        SELECT o.OrderID, u.Name, o.OrderDate, o.TotalAmount, o.Paid
        FROM Orders o
        JOIN User u ON o.UserID = u.UserID
        WHERE o.Paid = FALSE
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Unpaid Orders:")
	fmt.Printf("%-10s %-20s %-20s %-15s %-10s\n", "OrderID", "UserName", "OrderDate", "TotalAmount", "Paid")
	for rows.Next() {
		var orderID int
		var userName string
		var orderDate string
		var totalAmount float64
		var paid bool
		err := rows.Scan(&orderID, &userName, &orderDate, &totalAmount, &paid)
		if err != nil {
			return err
		}
		fmt.Printf("%-10d %-20s %-20s %-15.2f %-10t\n", orderID, userName, orderDate, totalAmount, paid)
	}

	// Prompt the user to enter an OrderID to pay
	var selectedOrderID int
	fmt.Print("\nEnter OrderID to pay: ")
	_, err = fmt.Scan(&selectedOrderID)
	if err != nil {
		return err
	}

	// Prompt the user to enter a payment method
	fmt.Println("Available payment methods: [bank-transfer, va, gopay, cc]")
	var paymentMethod string
	fmt.Print("Enter Payment Method: ")
	fmt.Scan(&paymentMethod)

	// Validate payment method
	validMethods := map[string]bool{
		"bank-transfer": true,
		"va":            true,
		"gopay":         true,
		"cc":            true,
	}
	if !validMethods[paymentMethod] {
		return fmt.Errorf("Invalid payment method: %s", paymentMethod)
	}

	// Process payment for the selected order
	var totalAmount float64
	err = ph.DB.QueryRow("SELECT TotalAmount FROM Orders WHERE OrderID = ? AND Paid = FALSE", selectedOrderID).Scan(&totalAmount)
	if err != nil {
		return err
	}

	_, err = ph.DB.Exec("INSERT INTO Payment (OrderID, PaymentDate, PaymentMethod, PaymentAmount) VALUES (?, NOW(), ?, ?)", selectedOrderID, paymentMethod, totalAmount)
	if err != nil {
		return err
	}

	// Update the order to mark it as paid
	_, err = ph.DB.Exec("UPDATE Orders SET Paid = TRUE WHERE OrderID = ?", selectedOrderID)
	if err != nil {
		return err
	}

	fmt.Printf("Payment processed successfully for Order %d, Method %s, Amount %.2f\n", selectedOrderID, paymentMethod, totalAmount)
	return nil
}

func (ph *PaymentHandlerImpl) updateOrderQuantity(orderDetailID, newQuantity int) error {
	_, err := ph.DB.Exec("UPDATE OrderDetails SET Quantity = ? WHERE OrderDetailID = ?", newQuantity, orderDetailID)
	return err
}

func (ph *PaymentHandlerImpl) deleteOrderDetail(orderDetailID int) error {
	_, err := ph.DB.Exec("DELETE FROM OrderDetails WHERE OrderDetailID = ?", orderDetailID)
	return err
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
