package services

import "github.com/gofiber/fiber/v2"

func Tester(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"user": "lorem",
		"password": "123456",
		"email":"lorem@ipsum.com",

})
}