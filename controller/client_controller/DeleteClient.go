package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
)

func DeleteClient(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn

	var client model.Client
	db.Preload(clause.Associations).First(&client, id)
	if client.ClientType == "" {
		ctx.Status(400).Send(helper.ErrorMessageJson("No client found with this ID"))
		return
	}

	db.Select(clause.Associations).Delete(&client)

	ctx.JSON(client)
}
