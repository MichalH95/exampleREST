package controller

import (
	"fmt"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Client model.Client

func GetClients(ctx *fiber.Ctx) {
	db := database.DBConn
	var clients []Client
	db.Preload("Company").Find(&clients)
	for i, client := range clients {
		fmt.Println(i)
		fmt.Println("Name: " + client.FirstName)
		fmt.Println("Name: " + client.Surname)
		fmt.Println("Company: " + client.Company.Name)
		fmt.Printf("Company: %v\n", client.Company.ID)
	}
	ctx.JSON(clients)
}

func PostClient(ctx *fiber.Ctx) {
	db := database.DBConn
	var client Client
	client.FirstName = "Jan"
	client.Surname = "Novak"
	client.Company = model.Company{
		Model: gorm.Model{},
		Name:  "Alza",
		ICO:   "12345",
	}
	db.Create(&client)
	ctx.JSON(client)
}
