package main

import (
	"hoopla/cli"
	"hoopla/config"
	"hoopla/handler"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()
	h := handler.NewHandler(db)
	ph := handler.NewPaymentHandler(db)
	cli := cli.NewCLI(h, ph)
	cli.Init()
}
