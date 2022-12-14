package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/services"
)


func SetupRoutes(app *fiber.App) {
	// USER ROUTES 

	app.Get("/",services.Tester)
	app.Get("/user/:id", services.GetUserById)
	app.Get("/users",services.GetAllUsers)
	app.Post("/user",services.CreateUser)
	app.Put("/user/:id",services.UpdateUser)
	app.Delete("/user/:id",services.DeleteUser)

	// CARD ROUTES

	app.Post("/card",services.CreateCard)


	
}