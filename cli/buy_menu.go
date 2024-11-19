package cli

import (
	"fmt"
	"log"
)

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
	fmt.Println("Item bought successfully.")
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}
