package cli

import (
	"fmt"
	"hoopla/handler"
	"log"
	"os"
	"os/exec"
)

type CLI struct {
	paymentHandler handler.PaymentHandler
	productHandler handler.HandlerProduct
	handler        handler.Handler
	user           *handler.User
}

func NewCLI(h handler.Handler, ph handler.PaymentHandler, hp handler.HandlerProduct) *CLI {
	return &CLI{handler: h, paymentHandler: ph, productHandler: hp}
}

func (c *CLI) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (c *CLI) showMenu(title string, options []string) {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println(title)
	fmt.Println("================================================")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Print("Enter the number of the option you want to select: ")
}

func (c *CLI) Init() {
	c.chooseUser()
	for {
		c.showMainMenu()
	}
}

func (c *CLI) chooseUser() {
	c.clearScreen()
	fmt.Println(Logo)
	fmt.Println("================================================")
	fmt.Println("USER SELECTION")
	fmt.Println("================================================")
	err := c.handler.ShowUserTable()
	if err != nil {
		log.Fatal(err)
	}
	userIDs, err := c.handler.GetUserIDs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please select a user by entering the User ID:")
	for {
		var userID int
		fmt.Print("Enter User ID: ")
		fmt.Scan(&userID)
		if contains(userIDs, userID) {
			user, err := c.handler.GetUserByID(userID)
			if err != nil {
				log.Fatal(err)
			}
			c.user = user
			break
		} else {
			fmt.Println("Invalid User ID. Please try again.")
		}
	}
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (c *CLI) showMainMenu() {
	var options []string
	if c.user.Role == "admin" {
		options = []string{
			"Reporting",
			"Update Item Stock",
			"Buy Item",
			"Show Orders",
			"Exit",
		}
	} else {
		options = []string{
			"Buy Item",
			"Show Orders",
			"Exit",
		}
	}
	c.showMenu("Dribbling through data, dunking on profits!!", options)
	fmt.Println("")
	var choice int
	fmt.Scan(&choice)
	c.handleMainMenuChoice(choice)
}

func (c *CLI) handleMainMenuChoice(choice int) {
	if c.user.Role == "admin" {
		switch choice {
		case 1:
			c.showReportingMenu()
		case 2:
			c.showProductMenu()
		case 3:
			c.showBuyMenu()
		case 4:
			c.showOrders()
		case 5:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	} else {
		switch choice {
		case 1:
			c.showBuyMenu()
		case 2:
			c.showOrders()
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
