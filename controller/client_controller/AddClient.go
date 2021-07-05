package client_controller

import (
	"github.com/MichalH95/exampleREST/controller/error_controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/MichalH95/exampleREST/model/error_response"
	"github.com/gofiber/fiber"
)

func AddClient(ctx *fiber.Ctx) {
	db := database.DBConn

	client := model.Client{}
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

	client.ID = 0
	client.Company.ID = 0
	client.Person.ID = 0

	db.Create(&client)
	ctx.JSON(client)
}
