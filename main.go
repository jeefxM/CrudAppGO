package main

import (
	"goLang/db"
	"goLang/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    client := db.NewClient()

    if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
    defer client.Prisma.Disconnect()

    app := fiber.New()

    routes.RegisterTodoRoutes(app, client) // This registers all your todo routes
    
    app.Listen(":3000")
}
