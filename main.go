package main

import (
	"fmt"
	"github.com/MichalH95/exampleREST/controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "examplerest_db"
)

func setupRoutes(app *fiber.App) {
	app.Get("/clients", controller.GetClients)
	app.Post("/clients", controller.AddClient)
	app.Put("/clients/:id", controller.UpdateClient)
	app.Delete("/clients/:id", controller.DeleteClient)

	app.Post("/sample", controller.PostSampleData)
}

func initDatabase() {
	var err error

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		host, port, user, password, dbname)
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection to database opened")
	database.DBConn.AutoMigrate(&model.Company{}, &model.Person{}, &model.Client{})
	//database.DBConn.Model(&model.Company{}).AddForeignKey("client_id", "clients(id)", "RESTRICT", "RESTRICT")
	//database.DBConn.Model(&model.Client{}).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
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
