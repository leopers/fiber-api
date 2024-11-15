package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leopers/fiber-api.git/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func main(){
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}