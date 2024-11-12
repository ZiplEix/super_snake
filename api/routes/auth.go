package routes

import (
	"github.com/ZiplEix/super_snake/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)
}
