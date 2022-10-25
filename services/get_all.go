package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	users := repository.GetAll()

	return ctx.Status(http.StatusOK).JSON(users)

}