package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsantos1/bank-go/repository"
	"github.com/vsantos1/bank-go/utils"
)

func UpdateUser(ctx *fiber.Ctx) error {
	id,errs := ctx.ParamsInt("id")
	if  errs != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}

	rep := repository.GetById(id)
	
	
	if err := ctx.BodyParser(rep.User); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}
	rep.User.Password = utils.MD5(rep.User.Password)
	
	res := repository.Update(id,rep.User)
	
	if res.Err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": res.Err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "User updated",
	})


}