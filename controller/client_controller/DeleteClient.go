package client_controller

import (
	"github.com/MichalH95/exampleREST/controller/error_controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/MichalH95/exampleREST/model/error_response"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
)

func DeleteClient(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn

	var client model.Client
	db.Preload(clause.Associations).First(&client, id)
	if client.ClientType == "" {
		error_controller.ClientErrorResponse(ctx, error_response.NoClientWithThisId)
		return
	}

	db.Select(clause.Associations).Delete(&client)

	ctx.JSON(client)
}
