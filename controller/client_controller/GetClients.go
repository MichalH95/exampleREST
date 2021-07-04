package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
)

func GetClients(ctx *fiber.Ctx) {
	db := database.DBConn

	var clients []model.Client
	db.Preload(clause.Associations).Find(&clients)

	ctx.JSON(clients)
}
