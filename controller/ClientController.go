package controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
)

type Client model.Client

func GetClients(ctx *fiber.Ctx) {
	db := database.DBConn
	var clients []Client
	db.Find(&clients)
	ctx.JSON(clients)
}

func PostClient(ctx *fiber.Ctx) {
	db := database.DBConn
	var client Client
	client.FirstName = "Jan"
	client.Surname = "Novak"
	db.Create(&client)
	ctx.JSON(client)
}
