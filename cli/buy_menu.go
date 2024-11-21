package cli

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func (c *CLI) showOrders() {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println("ORDER DETAILS TABLE")
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
	err := c.paymentHandler.ShowProductTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please enter the product details to buy the item:")
	var productID, quantity int
	fmt.Print("Enter Product ID: ")
	fmt.Scan(&productID)
	fmt.Print("Enter Quantity: ")
	fmt.Scan(&quantity)
	totalAmount := float64(quantity) * 10.0
	err = c.paymentHandler.ProcessOrder(c.user.UserID, productID, quantity, totalAmount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item bought and order processed successfully.")
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}

func (c *CLI) showPaymentMenu() {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println("PAYMENT MENU")
	fmt.Println("================================================")
	err := c.paymentHandler.ShowPaymentTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All payments processed successfully.")
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}
