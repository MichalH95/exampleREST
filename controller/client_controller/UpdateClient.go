package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
	"strconv"
)

func UpdateClient(ctx *fiber.Ctx) {
	idStr := ctx.Params("id")
	idUint64, convErr := strconv.ParseUint(idStr, 10, 32)

	if convErr != nil {
		ctx.Status(400).Send(helper.ErrorMessageJson("Negative client id"))
		return
	}

	id := uint(idUint64)
	db := database.DBConn

	// delete old client

	var client model.Client
	db.Preload(clause.Associations).First(&client, id)
	if client.ClientType == "" {
		ctx.Status(400).Send(helper.ErrorMessageJson("No client found with this ID"))
		return
	}

	db.Select(clause.Associations).Delete(&client)

	// insert new client

	client = model.Client{}
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

	client.ID = id
	client.Company.ID = 0
	client.Person.ID = 0

	db.Create(&client)

	ctx.JSON(client)
}
