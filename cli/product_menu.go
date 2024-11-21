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
		fmt.Println("1. List All Products")
		fmt.Println("2. Create New Product")
		fmt.Println("3. Update Product Stock")
		fmt.Println("4. Delete a Product")
		fmt.Println("5. Return to Main Menu")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		//		switch choice {
		//		case 1:
		//			// List all products
		//			err := c.handler.ShowProductTable()
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//			fmt.Println("\nPress Enter to return to the Update Item Stock Menu...")
		//			fmt.Scanln()
		//			fmt.Scanln() // Untuk menangani buffer input
		//
		//		case 2:
		//			// Create a new product
		//			fmt.Println("Create New Product")
		//			var name, description string
		//			var categoryID, stockQuantity int
		//			var price float64
		//			fmt.Print("Enter Product Name: ")
		//			fmt.Scan(&name)
		//			fmt.Print("Enter Category ID: ")
		//			fmt.Scan(&categoryID)
		//			fmt.Print("Enter Price: ")
		//			fmt.Scan(&price)
		//			fmt.Print("Enter Stock Quantity: ")
		//			fmt.Scan(&stockQuantity)
		//			fmt.Print("Enter Description: ")
		//			fmt.Scan(&description)
		//			err := c.handler.CreateProduct(name, categoryID, price, stockQuantity, description)
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//			fmt.Println("Product created successfully!")
		//
		//		case 3:
		//			// Update stock for a product
		//			fmt.Println("Update Product Stock")
		//			var productID, stockQuantity int
		//			fmt.Print("Enter Product ID: ")
		//			fmt.Scan(&productID)
		//			fmt.Print("Enter Stock Quantity: ")
		//			fmt.Scan(&stockQuantity)
		//			err := c.handler.UpdateProduct(productID, stockQuantity)
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//			fmt.Println("Product stock updated successfully!")
		//
		//		case 4:
		//			// Delete a product
		//			fmt.Println("Delete a Product")
		//			var productID int
		//			fmt.Print("Enter Product ID: ")
		//			fmt.Scan(&productID)
		//			err := c.handler.DeleteProduct(productID)
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//			fmt.Println("Product deleted successfully!")
		//
		//		case 5:
		//			// Return to main menu
		//			return
		//
		//		default:
		//			fmt.Println("Invalid choice. Please try again.")
		//		}
		//
		//		fmt.Println("\nPress Enter to continue...")
		//		fmt.Scanln()
		//		fmt.Scanln() // Untuk menangani buffer input
		//
		//	}
		//}

		switch choice {
		case 1:
			// List all products
			err := c.handler.ShowProductTable()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\nPress Enter to return to the Update Item Stock Menu...")
			fmt.Scanln()
			fmt.Scanln() // Untuk menangani buffer input

		case 2:
			// Create a new product
			fmt.Println("Create New Product")
			var name, description string
			var categoryID, stockQuantity int
			var price float64
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Category ID: ")
			fmt.Scan(&categoryID)
			fmt.Print("Enter Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Stock Quantity: ")
			fmt.Scan(&stockQuantity)
			fmt.Print("Enter Description: ")
			fmt.Scan(&description)
			err := c.handlerProduct.CreateProductX(name, categoryID, price, stockQuantity, description)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Product created successfully!")

		case 3:
			// Update stock for a product
			fmt.Println("Update Product Stock")
			var productID, stockQuantity int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&productID)
			fmt.Print("Enter Stock Quantity: ")
			fmt.Scan(&stockQuantity)
			err := c.handlerProduct.UpdateProductX(productID, stockQuantity)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Product stock updated successfully!")

		case 4:
			// Delete a product
			fmt.Println("Delete a Product")
			var productID int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&productID)
			err := c.handlerProduct.DeleteProductX(productID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Product deleted successfully!")

		case 5:
			// Return to main menu
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
		fmt.Scanln() // Untuk menangani buffer input

	}
}
