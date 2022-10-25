package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/services"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/user/:id", services.GetUserById)
	app.Get("/users",services.GetAllUsers)
	app.Post("/user",services.CreateUser)
	
}