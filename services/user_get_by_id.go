package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
)

func GetUserById(ctx *fiber.Ctx) error {
	id,err := ctx.ParamsInt("id")
	
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	user := repository.GetById(id)
	return ctx.JSON(user)
}