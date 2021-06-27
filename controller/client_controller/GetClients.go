package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
)

func GetClients(ctx *fiber.Ctx) {
	var output []interface{}
	db := database.DBConn

	var companies []model.Company
	db.Preload("Client").Find(&companies)

	for _, company := range companies {
		output = append(output, company)
	}

	var people []model.Person
	db.Preload("Client").Find(&people)

	for _, person := range people {
		output = append(output, person)
	}

	ctx.JSON(output)
}
