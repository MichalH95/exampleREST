package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
)

func GetClients(ctx *fiber.Ctx) {
	db := database.DBConn

	var clients []model.Client
	db.Preload("Company").Preload("Person").Find(&clients)

	ctx.JSON(clients)
}
