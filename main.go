package main

import (
	"fmt"
	"github.com/MichalH95/exampleREST/controller/client_controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/clients", client_controller.GetClients)
	app.Post("/clients", client_controller.AddClient)
	app.Put("/clients/:id", client_controller.UpdateClient)
	app.Delete("/clients/:id", client_controller.DeleteClient)

	app.Post("/sample", client_controller.AddSampleData)
}

func initDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Dbname)

	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection to database opened")
	database.DBConn.AutoMigrate(&model.Company{}, &model.Person{}, &model.Client{})
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	sqlDB, err := database.DBConn.DB()
	if err == nil {
		defer sqlDB.Close()
	}
}
