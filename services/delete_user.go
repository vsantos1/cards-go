package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
)

func DeleteUser(ctx *fiber.Ctx) error {
	id, errs := ctx.ParamsInt("id")

	u := repository.GetById(id)

	if u.Err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	
	if errs != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	res := repository.Delete(id)

	if res.Err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": res.Err.Error(),
		})
	}

	return ctx.Status(http.StatusNoContent).JSON(fiber.Map{})
}