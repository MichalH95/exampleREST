package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
)

func AddClient(ctx *fiber.Ctx) {
	db := database.DBConn

	client := model.Client{}
	err := ctx.BodyParser(&client)

	if err != nil {
		ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
		return
	}

	if client.ClientType != model.ClientTypeCompany && client.ClientType != model.ClientTypePerson {
		// received data doesn't have valid client type
		ctx.Status(400).Send(helper.ErrorMessageJson("Invalid Client.ClientType, specify either 1 for company or 2 for person"))
		return
	}

	client.ID = 0
	client.Company.ID = 0
	client.Person.ID = 0

	db.Create(&client)
	ctx.JSON(client)
}
