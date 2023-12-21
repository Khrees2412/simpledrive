package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/simpledrive/controller"
)

func RegisterRoutes(router *fiber.App) {
	controller.NewAuthController().RegisterRoutes(router)
	controller.NewFileController().RegisterRoutes(router)
}
