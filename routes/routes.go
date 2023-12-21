package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/simpledrive/controller"
)

func RegisterRoutes(router *fiber.App) {
	controller.NewAuthController().RegisterRoutes(router)
}

//private.Post("/create-folder", controller.CreateFolder)
//
//private.Post("/upload/:folder", controller.StoreFileInFolder)
//
//private.Post("/upload", controller.StoreFile)
//
//private.Get("/download/:filename", controller.DownloadFile)
//
//private.Get("/view/files", controller.GetFile)
//
//private.Delete("/:fileID", controller.DeleteFile)
