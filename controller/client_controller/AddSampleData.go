package client_controller

import (
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"time"
)

func AddSampleData(ctx *fiber.Ctx) {
	db := database.DBConn

	client1 := model.Client{
		Company: model.Company{
			Name:             "Alza",
			ICO:              "54321",
			ContactFirstName: "Josef",
			ContactLastName:  "Chodounsky",
		},
		ClientType: model.ClientTypeCompany,
	}

	client2 := model.Client{
		Company: model.Company{
			Name:             "CZC.cz",
			ICO:              "37952",
			ContactFirstName: "Petr",
			ContactLastName:  "Nemec",
		},
		ClientType: model.ClientTypeCompany,
	}

	client3 := model.Client{
		Person: model.Person{
			FirstName: "Jan",
			LastName:  "Novak",
			BirthDate: time.Now().AddDate(-25, 7, 2),
		},
		ClientType: model.ClientTypePerson,
	}

	client4 := model.Client{
		Person: model.Person{
			FirstName: "Jaromir",
			LastName:  "Meduna",
			BirthDate: time.Now().AddDate(-32, 4, 18),
		},
		ClientType: model.ClientTypePerson,
	}

	db.Create(&client1)
	db.Create(&client2)
	db.Create(&client3)
	db.Create(&client4)

	output := []interface{}{client1, client2, client3, client4}

	ctx.JSON(output)
}
