package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/models"
	"github.com/vsantos1/bank-go/repository"
)

func CreateCard(ctx *fiber.Ctx) error {

	card := new(models.Card)

	if err := ctx.BodyParser(card); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
}

rep_user := repository.GetById(card.UserId)

if rep_user.Err != nil {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": "User not found",
	})
}


rep := repository.CreateCard(card.UserId,card)

if rep.Err != nil {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": rep.Err.Error(),
	})
}


return ctx.Status(http.StatusCreated).JSON(fiber.Map{
	"message": "Card created",
})

}