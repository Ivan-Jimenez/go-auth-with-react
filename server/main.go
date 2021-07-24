package main

import (
	"github.com/Ivan-Jimenez/go-auth-with-react/server/database"
	"github.com/Ivan-Jimenez/go-auth-with-react/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)
	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	app.Listen(":4000")
}
