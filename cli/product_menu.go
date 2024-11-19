package cli

import (
	"fmt"
	"log"
)

func (c *CLI) showProductMenu() {
	for {
		c.clearScreen()
		fmt.Println(Logo)
		fmt.Println("================================================")
		fmt.Println("UPDATE ITEM STOCK MENU")
		fmt.Println("================================================")
		err := c.handler.ShowProductTable()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Please enter the product details to update the stock:")
		var productID, stockQuantity int
		fmt.Print("Enter Product ID: ")
		fmt.Scan(&productID)
		fmt.Print("Enter Stock Quantity: ")
		fmt.Scan(&stockQuantity)
		err = c.handler.UpdateProductStock(productID, stockQuantity)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Product stock updated successfully.")
		fmt.Println("Do you want to update another item? (y/n): ")
		var choice string
		fmt.Scan(&choice)
		if choice != "y" {
			break
		}
	}
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
	fmt.Scanln()
}
