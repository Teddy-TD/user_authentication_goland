package main

import (
	"log"
	"user_authentication_goland/routes"

	"../database"
	"./routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
    // Database connection
    database.Connect()
    
    // Initialize a new Fiber app
    app := fiber.New()

    routes.Setup(app)


    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}


