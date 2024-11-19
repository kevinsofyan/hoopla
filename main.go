package main

import (
	"hoopla/cli"
	"hoopla/config"
	"hoopla/handler"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()
	handler := handler.NewHandler(db)
	cli := cli.NewCLI(handler)
	cli.Init()
}
