package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/routes"
	"github.com/vsantos1/bank-go/utils"
)

func main() {
	app := fiber.New()
	fmt.Println(utils.Dec(3))
	routes.SetupRoutes(app)
	app.Listen(":3000")
	
} 