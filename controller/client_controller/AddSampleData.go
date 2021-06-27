package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/helper"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"time"
)

func AddSampleData(ctx *fiber.Ctx) {
	db := database.DBConn

	company1 := model.Company{
		Name:             "Alza",
		ICO:              "54321",
		ContactFirstName: "Josef",
		ContactLastName:  "Chodounsky",
		Client:           helper.NewClientAsCompany(),
	}
	company2 := model.Company{
		Name:             "CZC.cz",
		ICO:              "37952",
		ContactFirstName: "Petr",
		ContactLastName:  "Nemec",
		Client:           helper.NewClientAsCompany(),
	}
	person1 := model.Person{
		FirstName: "Jan",
		LastName:  "Novak",
		BirthDate: time.Now().AddDate(-25, 7, 2),
		Client:    helper.NewClientAsPerson(),
	}
	person2 := model.Person{
		FirstName: "Jaromir",
		LastName:  "Meduna",
		BirthDate: time.Now().AddDate(-32, 4, 18),
		Client:    helper.NewClientAsPerson(),
	}

	db.Create(&company1)
	db.Create(&company2)
	db.Create(&person1)
	db.Create(&person2)

	output := []interface{}{company1, company2, person1, person2}

	ctx.JSON(output)
}
