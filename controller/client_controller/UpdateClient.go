package client_controller

import (
	"github.com/MichalH95/exampleREST/controller/error_controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/MichalH95/exampleREST/model/error_response"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
	"strconv"
)

func UpdateClient(ctx *fiber.Ctx) {
	idStr := ctx.Params("id")
	idUint64, convErr := strconv.ParseUint(idStr, 10, 32)

	if convErr != nil {
		error_controller.ClientErrorResponse(ctx, error_response.NegativeClientId)
		return
	}

	id := uint(idUint64)
	db := database.DBConn

	// delete old client

	var client model.Client
	db.Preload(clause.Associations).First(&client, id)
	if client.ClientType == "" {
		error_controller.ClientErrorResponse(ctx, error_response.NoClientWithThisId)
		return
	}

	db.Select(clause.Associations).Delete(&client)

	// insert new client

	client = model.Client{}
	err := ctx.BodyParser(&client)

	if err != nil {
		error_controller.ServerErrorResponse(ctx, err.Error())
		return
	}

	if client.ClientType != model.ClientTypeCompany && client.ClientType != model.ClientTypePerson {
		// received data doesn't have valid client type
		error_controller.ClientErrorResponse(ctx, error_response.InvalidClientType)
		return
	}

	client.ID = id
	client.Company.ID = 0
	client.Person.ID = 0

	db.Create(&client)

	ctx.JSON(client)
}
