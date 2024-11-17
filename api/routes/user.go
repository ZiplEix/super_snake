package routes

import (
	"github.com/ZiplEix/super_snake/api/controllers"
	"github.com/ZiplEix/super_snake/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/me", middleware.Protected, controllers.Me)
}
