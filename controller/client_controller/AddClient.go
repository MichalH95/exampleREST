package client_controller

import (
	"github.com/gofiber/fiber"
)

func AddClient(ctx *fiber.Ctx) {
	//db := database.DBConn
	//
	//// parse as company, then check for Client.ClientType
	//company := model.Company{}
	//err := ctx.BodyParser(&company)
	//if err != nil {
	//	ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
	//	return
	//}
	//// check if received data has client type
	//if company.Client.ClientType != 1 && company.Client.ClientType != 2 {
	//	// received data doesn't have valid client type
	//	ctx.Status(400).Send(helper.ErrorMessageJson("Invalid Client.ClientType, specify either 1 for company or 2 for person"))
	//	return
	//}
	//// check if received data is person
	//if company.Client.ClientType == model.ClientTypePerson {
	//	// received data is person
	//	person := model.Person{}
	//	err := ctx.BodyParser(&person)
	//	if err != nil {
	//		ctx.Status(503).Send(helper.ErrorMessageJson(err.Error()))
	//		return
	//	}
	//
	//	person.Model = gorm.Model{}
	//	person.Client = helper.NewClientAsPerson()
	//
	//	db.Create(&person)
	//
	//	ctx.JSON(person)
	//	return
	//}
	//// received data is company
	//company.Model = gorm.Model{}
	//company.Client = helper.NewClientAsCompany()
	//
	//db.Create(&company)
	//
	//ctx.JSON(company)
}
