package main

import (
	"fmt"
	"github.com/MichalH95/exampleREST/controller"
	"github.com/MichalH95/exampleREST/database"
	"github.com/MichalH95/exampleREST/model"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Client model.Client

func setupRoutes(app *fiber.App) {
	app.Get("/clients", controller.GetClients)
	app.Post("/clients", controller.PostClient)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection to database opened")
	database.DBConn.AutoMigrate(&Client{})
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
