package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(repository.GetAll())

}