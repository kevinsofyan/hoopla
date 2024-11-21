package cli

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func (c *CLI) showOrders() {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println("ORDER TABLE")
	fmt.Println("================================================")
	err := c.paymentHandler.ShowOrderTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}

func (c *CLI) showBuyMenu() {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println("BUY ITEM MENU")
	fmt.Println("================================================")
	err := c.handler.ShowProductTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please enter the product details to buy the item:")
	var productID, quantity int
	fmt.Print("Enter Product ID: ")
	fmt.Scan(&productID)
	fmt.Print("Enter Quantity: ")
	fmt.Scan(&quantity)
	err = c.handler.BuyItem(productID, quantity)
	if err != nil {
		log.Fatal(err)
	}
	totalAmount := float64(quantity) * 10.0
	err = c.paymentHandler.ProcessOrder(c.user.UserID, productID, quantity, totalAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item bought and payment processed successfully.")
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}
