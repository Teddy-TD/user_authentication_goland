package routes

import (
	"user_authentication_goland/controllers"

	"./controllers"
	"github.com/gofiber/fiber/v3"
)

func Setup (app *fiber.App) {

	 // Define a route for the GET method on the root path '/'
	 app.Get("/", controllers.Hello)

}