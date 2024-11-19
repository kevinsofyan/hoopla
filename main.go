package main

import (
	"hoopla/cli"
	"hoopla/config"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	cli.Init()
}
