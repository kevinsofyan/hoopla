package cli

import (
	"fmt"
)

func (c *CLI) showProductMenu() {
	for {
		c.clearScreen()
		fmt.Println(Logo)
		fmt.Println("================================================")
		fmt.Println("UPDATE ITEM STOCK MENU")
		fmt.Println("================================================")
		fmt.Println("1. List All Products")
		fmt.Println("2. Create New Product")
		fmt.Println("3. Update Product Stock")
		fmt.Println("4. Delete a Product")
		fmt.Println("5. Return to Main Menu")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			// Return to main menu
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
		fmt.Scanln()
	}
}
