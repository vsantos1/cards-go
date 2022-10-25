package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
)

func GetUserById(ctx *fiber.Ctx) error {
	id,err := ctx.ParamsInt("id")
	
	user := repository.GetById(id)

	if user.Err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	
	return ctx.JSON(user)
}