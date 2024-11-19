package cli

import (
	"fmt"
	"log"
)

func (c *CLI) showReportingMenu() {
	options := []string{
		"Total Sales Report",
		"Most Popular Product Report",
		"Total Revenue Per Product Report",
		"Customer Count Per City Report",
		"Back to Main Menu",
	}
	c.showMenu("REPORTING MENU", options)
	var choice int
	fmt.Scan(&choice)
	c.handleReportingMenuChoice(choice)
}

func (c *CLI) handleReportingMenuChoice(choice int) {
	switch choice {
	case 1:
		err := c.handler.TotalSalesReport()
		if err != nil {
			log.Fatal(err)
		}
	case 2:
		err := c.handler.MostPopularProductReport()
		if err != nil {
			log.Fatal(err)
		}
	case 3:
		err := c.handler.TotalRevenuePerProductReport()
		if err != nil {
			log.Fatal(err)
		}
	case 4:
		err := c.handler.CustomerCountPerCityReport()
		if err != nil {
			log.Fatal(err)
		}
	case 5:
		return
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Println("\nPress Enter to return to the reporting menu...")
	fmt.Scanln()
	fmt.Scanln()
}
