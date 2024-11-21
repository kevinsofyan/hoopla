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
	productHandler := handler.NewHandlerProduct(db)
	//hPP := handler_product.NewHandlerProduct(db)

	cli := cli.NewCLI(handler)
	cli.ProductCLI(productHandler)
	cli.Init()

	//cli.ProductCLI(productHandler)

}
