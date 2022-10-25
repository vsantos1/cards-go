package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
	
} 