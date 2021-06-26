package controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	time "time"
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

func PostClient(ctx *fiber.Ctx) {

}

func PostSampleData(ctx *fiber.Ctx) {
	var output []interface{}
	db := database.DBConn

	var company1 model.Company
	company1.Name = "Alza"
	company1.ICO = "54321"
	company1.ContactFirstName = "Josef"
	company1.ContactLastName = "Chodounsky"
	company1.Client = helper.NewClientWithCompanyId()
	db.Create(&company1)

	output = append(output, company1)

	var company2 model.Company
	company2.Name = "CZC.cz"
	company2.ICO = "37952"
	company2.ContactFirstName = "Petr"
	company2.ContactLastName = "Nemec"
	company2.Client = helper.NewClientWithCompanyId()
	db.Create(&company2)

	output = append(output, company2)

	var person1 model.Person
	person1.FirstName = "Jan"
	person1.LastName = "Novak"
	person1.BirthDate = time.Now().AddDate(-25, 7, 2)
	person1.Client = helper.NewClientWithPersonId()
	db.Create(&person1)

	output = append(output, person1)

	var person2 model.Person
	person2.FirstName = "Jaromir"
	person2.LastName = "Meduna"
	person2.BirthDate = time.Now().AddDate(-32, 4, 18)
	person2.Client = helper.NewClientWithPersonId()
	db.Create(&person2)

	output = append(output, person2)

	ctx.JSON(output)
}
