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
	db.First(&client, id)
	if client.ClientType == 0 {
		ctx.Status(400).Send(helper.ErrorMessageJson("No client found with this ID"))
		return
	}
	if client.ClientType == 1 {
		// client to delete is company
		var company model.Company
		db.Preload("Client").First(&company, client.CompanyId.Int64)
		db.Delete(&company)
		ctx.JSON(company)
	}
	if client.ClientType == 2 {
		// client to delete is person
		var person model.Person
		db.Preload("Client").First(&person, client.PersonId.Int64)
		db.Delete(&person)
		ctx.JSON(person)
	}

	db.Delete(&client)
}
