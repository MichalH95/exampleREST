package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
)

func DeleteClient(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn

	var client model.Client
	db.Preload("Person").Preload("Company").First(&client, id)
	if client.ClientType == "" {
		ctx.Status(400).Send(helper.ErrorMessageJson("No client found with this ID"))
		return
	}

	ctx.JSON(client)
	db.Delete(&client)
}
