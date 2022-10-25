package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/models"
	"github.com/vsantos1/bank-go/repository"
	"github.com/vsantos1/bank-go/utils"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
}

user.Password = utils.MD5(user.Password)

rep := repository.Create(user)
if rep.Err != nil {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": "Email already exists",
	})
}


return ctx.Status(http.StatusCreated).JSON(fiber.Map{
	"message": "User created",
})

}