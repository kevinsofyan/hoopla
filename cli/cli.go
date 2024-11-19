package cli

import (
	"fmt"
)

type CLI struct {
}

func showMenu() {
	fmt.Println("\n\n\n\n\n\n\n\n================================================")
	fmt.Println("GAMES REPORT TRACKER ADMIN CLI")
	fmt.Println("================================================")
	fmt.Println("Please select report to generate:")
	fmt.Println("1. Total Games Sales Report")
	fmt.Println("2. Most Popular Game Report")
	fmt.Println("3. Total Revenue Per Game Report")
	fmt.Println("4. Player Count Per Game Report")
	fmt.Println("5. Exit")
	fmt.Print("Enter the number of the report you want to generate:")

	var choice int
	fmt.Scan(&choice)
	fmt.Println("")

	switch choice {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Println("\nPress Enter to return to the menu...")
	fmt.Scanln()

	showMenu()

}
func Init() {
	showMenu()
}
