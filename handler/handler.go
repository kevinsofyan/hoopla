package handler

import (
	"database/sql"
	"fmt"
)

type Handler interface {
	TotalSalesReport() error
	MostPopularProductReport() error
	TotalRevenuePerProductReport() error
	CustomerCountPerCityReport() error
	UpdateProductStock(productID, stockQuantity int) error
	BuyItem(productID, quantity int) error
	ShowProductTable() error
	ShowUserTable() error
	GetUserIDs() ([]int, error)
	GetUserByID(userID int) (*User, error)
	CreateProduct(name string, categoryID int, price float64, stockQuantity int, description string) error
	UpdateProduct(productID, stockQuantity int) error
	DeleteProduct(productID int) error
}

type User struct {
	UserID  int
	Role    string
	Name    string
	Email   string
	Phone   string
	Address string
	City    string
}

type HandlerImpl struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *HandlerImpl {
	return &HandlerImpl{
		DB: db,
	}
}

func (h *HandlerImpl) TotalSalesReport() error {
	fmt.Println("Generating Total Sales Report...")
	return nil
}

func (h *HandlerImpl) MostPopularProductReport() error {
	fmt.Println("Generating Most Popular Product Report...")
	return nil
}

func (h *HandlerImpl) TotalRevenuePerProductReport() error {
	fmt.Println("Generating Total Revenue Per Product Report...")
	return nil
}

func (h *HandlerImpl) CustomerCountPerCityReport() error {
	fmt.Println("Generating Customer Count Per City Report...")
	return nil
}

func (h *HandlerImpl) UpdateProductStock(productID, stockQuantity int) error {
	fmt.Println("Updating Product Stock...")
	return nil
}

func (h *HandlerImpl) BuyItem(productID, quantity int) error {
	fmt.Println("Buying Item...")
	return nil
}

func (h *HandlerImpl) ShowProductTable() error {
	rows, err := h.DB.Query("SELECT ProductID, ProductName, CategoryID, Price, StockQuantity FROM Product")
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

func (h *HandlerImpl) ShowUserTable() error {
	rows, err := h.DB.Query("SELECT UserID, Role, Name, Email, Phone, Address, City FROM User")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("User Table:")
	fmt.Printf("%-10s %-10s %-20s %-30s %-15s %-30s %-20s\n", "ID", "Role", "Name", "Email", "Phone", "Address", "City")
	for rows.Next() {
		var id int
		var role, name, email, phone, address, city string
		err := rows.Scan(&id, &role, &name, &email, &phone, &address, &city)
		if err != nil {
			return err
		}
		fmt.Printf("%-10d %-10s %-20s %-30s %-15s %-30s %-20s\n", id, role, name, email, phone, address, city)
	}
	return nil
}

func (h *HandlerImpl) GetUserIDs() ([]int, error) {
	rows, err := h.DB.Query("SELECT UserID FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}
	return userIDs, nil
}

func (h *HandlerImpl) GetUserByID(userID int) (*User, error) {
	row := h.DB.QueryRow("SELECT UserID, Role, Name, Email, Phone, Address, City FROM User WHERE UserID = ?", userID)
	var user User
	err := row.Scan(&user.UserID, &user.Role, &user.Name, &user.Email, &user.Phone, &user.Address, &user.City)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (h *HandlerImpl) CreateProduct(name string, categoryID int, price float64, stockQuantity int, description string) error {
	query := "INSERT INTO Product (ProductName, CategoryID, Price, StockQuantity, Description) VALUES (?, ?, ?, ?, ?)"
	_, err := h.DB.Exec(query, name, categoryID, price, stockQuantity, description)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlerImpl) UpdateProduct(productID, stockQuantity int) error {
	query := "UPDATE Product SET StockQuantity = ? WHERE ProductID = ?"
	_, err := h.DB.Exec(query, stockQuantity, productID)
	if err != nil {
		return fmt.Errorf("failed to update stock for ProductID %d: %v", productID, err)
	}
	return err
}

func (h *HandlerImpl) DeleteProduct(productID int) error {
	_, err := h.DB.Exec("DELETE FROM Product WHERE ProductID = ?", productID)
	if err != nil {
		return err
	}
	return nil
}

func (h *HandlerImpl) NewHandlerProduct(db *sql.DB) interface{} {

}
