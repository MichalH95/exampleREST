package controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"time"
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

func AddClient(ctx *fiber.Ctx) {
	db := database.DBConn

	// parse as company, then check for Client.ClientType
	company := new(model.Company)
	err := ctx.BodyParser(company)
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
		person := new(model.Person)
		err := ctx.BodyParser(person)
		if err != nil {
			ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
			return
		}

		person.Model = gorm.Model{}
		person.Client = helper.NewClientAsPerson()

		db.Create(person)

		ctx.JSON(person)
		return
	}
	// received data is company
	company.Model = gorm.Model{}
	company.Client = helper.NewClientAsCompany()

	db.Create(company)

	ctx.JSON(company)
}

func UpdateClient(ctx *fiber.Ctx) {

}

func DeleteClient(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn

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
		ctx.JSON(company)
	}
	if client.ClientType == 2 {
		// client to delete is person
		var person model.Person
		db.Preload("Client").First(&person, client.PersonId.Int64)
		db.Delete(&person)
		ctx.JSON(person)
	}

	db.Delete(&client)
}

func PostSampleData(ctx *fiber.Ctx) {
	var output []interface{}
	db := database.DBConn

	var company1 model.Company
	company1.Name = "Alza"
	company1.ICO = "54321"
	company1.ContactFirstName = "Josef"
	company1.ContactLastName = "Chodounsky"
	company1.Client = helper.NewClientAsCompany()
	db.Create(&company1)

	output = append(output, company1)

	var company2 model.Company
	company2.Name = "CZC.cz"
	company2.ICO = "37952"
	company2.ContactFirstName = "Petr"
	company2.ContactLastName = "Nemec"
	company2.Client = helper.NewClientAsCompany()
	db.Create(&company2)

	output = append(output, company2)

	var person1 model.Person
	person1.FirstName = "Jan"
	person1.LastName = "Novak"
	person1.BirthDate = time.Now().AddDate(-25, 7, 2)
	person1.Client = helper.NewClientAsPerson()
	db.Create(&person1)

	output = append(output, person1)

	var person2 model.Person
	person2.FirstName = "Jaromir"
	person2.LastName = "Meduna"
	person2.BirthDate = time.Now().AddDate(-32, 4, 18)
	person2.Client = helper.NewClientAsPerson()
	db.Create(&person2)

	output = append(output, person2)

	ctx.JSON(output)
}
