package handler

import (
	"database/sql"
	"fmt"
)

type HandlerProduct interface {
	ShowProduct() error
	CreateProductX(name string, categoryID int, price float64, stockQuantity int, description string) error
	UpdateProductX(productID, stockQuantity int) error
	DeleteProductX(productID int) error // Tambahkan ini jika belum ada untuk fitur delete
}

func NewHandlerProduct(db *sql.DB) *HandlerImpl {
	return &HandlerImpl{
		DB: db,
	}
}

func (hp *HandlerImpl) ShowProduct() error {
	rows, err := hp.DB.Query("SELECT ProductID, ProductName, CategoryID, Price, StockQuantity FROM Product")
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

func (hp *HandlerImpl) CreateProductX(name string, categoryID int, price float64, stockQuantity int, description string) error {
	query := "INSERT INTO Product (ProductName, CategoryID, Price, StockQuantity, Description) VALUES (?, ?, ?, ?, ?)"
	_, err := hp.DB.Exec(query, name, categoryID, price, stockQuantity, description)
	if err != nil {
		return err
	}
	return nil
}
func (hp *HandlerImpl) UpdateProductX(productID, stockQuantity int) error {
	query := "UPDATE Product SET StockQuantity = ? WHERE ProductID = ?"
	_, err := hp.DB.Exec(query, stockQuantity, productID)
	if err != nil {
		return fmt.Errorf("failed to update stock for ProductID %d: %v", productID, err)
	}
	return err
}

func (hp *HandlerImpl) DeleteProductX(productID int) error {
	_, err := hp.DB.Exec("DELETE FROM Product WHERE ProductID = ?", productID)
	if err != nil {
		return err
	}
	return nil
}
