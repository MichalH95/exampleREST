package client_controller

import (
	"database/sql"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
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

	// first delete old company/person, then add the new (updated) one

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

	}
	if client.ClientType == 2 {
		// client to delete is person
		var person model.Person
		db.Preload("Client").First(&person, client.PersonId.Int64)
		db.Delete(&person)
	}

	// old company/person deleted, now add the new (updated) one

	// make new client with the same ID but without client type and company/person id
	client = model.Client{ID: id}

	// parse as company, then check for Client.ClientType
	company := model.Company{}
	err := ctx.BodyParser(&company)
	if err != nil {
		ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
		return
	}
	// check if received data has client type
	if company.Client.ClientType != 1 && company.Client.ClientType != 2 {
		// received data doesn't have valid client type
		ctx.Status(400).Send(helper.ErrorMessageJson("Invalid Client.ClientType, specify either 1 for company or 2 for person"))
		return
	}
	// check if received data is person
	if company.Client.ClientType == model.ClientTypePerson {
		// received data is person
		person := model.Person{}
		err := ctx.BodyParser(&person)
		if err != nil {
			ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
			return
		}

		person.Model = gorm.Model{}
		person.Client = model.Client{}

		// insert person
		db.Create(&person)
		// update client
		client.ClientType = model.ClientTypePerson
		client.PersonId = sql.NullInt64{
			Int64: int64(person.ID),
			Valid: true,
		}
		db.Save(&client)

		// send updated person
		person.Client = client

		ctx.JSON(person)
		return
	}
	// received data is company
	company.Model = gorm.Model{}
	company.Client = model.Client{}

	// insert company
	db.Create(&company)
	// update client
	client.ClientType = model.ClientTypeCompany
	client.CompanyId = sql.NullInt64{
		Int64: int64(company.ID),
		Valid: true,
	}
	db.Save(&client)

	// send updated company
	company.Client = client

	ctx.JSON(company)
}
