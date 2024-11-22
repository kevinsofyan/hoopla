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
	InsertCategory(name string) error
	DeleteCategory(name string) error
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
	rows, err := h.DB.Query(`SELECT SUM(payment.PaymentAmount) FROM payment`)
	if err != nil {
		return err
	}
	defer rows.Close()
	var hasil float64;
	fmt.Println("Generating Total Sales Report...", hasil);
	for rows.Next() {
		err := rows.Scan(&hasil)
		if err != nil {
			return err
		}
		fmt.Printf("Total Sales Rp.: %f\n", hasil)
	}
	return nil
}

func (h *HandlerImpl) MostPopularProductReport() error {
	rows, err := h.DB.Query(`SELECT ProductName, SUM(Quantity) AS TotalQuantity
	FROM orderdetails
	GROUP BY ProductID
	ORDER BY TotalQuantity DESC`)
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("Generating Most Popular Product Report...")
	for rows.Next() {
		var productID, totalQuantity int
		err := rows.Scan(&productID, &totalQuantity)
		if err != nil {
			return err
		}
		fmt.Printf("Product ID: %d, Total Quantity: %d\n", productID, totalQuantity)
	}
	return nil
}

func (h *HandlerImpl) TotalRevenuePerProductReport() error {
	rows, err := h.DB.Query(`SELECT product.Price * orderdetails.Quantity 
	FROM product
	INNER JOIN orderdetails
	ON product.ProductID = orderdetails.ProductID`)
	if err != nil {
		return err
	}
	defer rows.Close()
	fmt.Println("Generating Total Revenue Per Product Report...")
	for rows.Next() {
		var totalRevenue float64;
		err := rows.Scan(&totalRevenue)
		if err != nil {
			return err
		}
		fmt.Printf("Total Revenue: %.2f\n", totalRevenue)
	}
	return nil
}

func (h *HandlerImpl) CustomerCountPerCityReport() error {
	fmt.Println("Generating Customer Count Per City Report...")
	row, err := h.DB.Query(`SELECT City, COUNT(*) AS CustomerCount
	FROM user
	WHERE Role = 'customer'
	GROUP BY City
	ORDER BY 
    CustomerCount DESC`)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		var city string
		var customerCount int
		err := row.Scan(&city, &customerCount)
		if err != nil {
			return err
		}
		fmt.Printf("City: %s, Customer Count: %d\n", city, customerCount)	
	}
	
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

// insert category
func (h *HandlerImpl) InsertCategory(name string) error {
	_, err := h.DB.Exec("INSERT INTO Category (CategoryName) VALUES (?)", name)
	if err != nil {
		return err
	}
	return nil
}

// delete category
func (h *HandlerImpl) DeleteCategory(name string) error {
	_, err := h.DB.Exec("DELETE FROM Category WHERE CategoryName = ?", name)
	if err != nil {
		return err
	}
	return nil
}
