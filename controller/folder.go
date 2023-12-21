package controller

//
//import (
//	"fmt"
//
//	db "github.com/khrees2412/simpledrive/database"
//	"github.com/khrees2412/simpledrive/model"
//	"github.com/khrees2412/simpledrive/types"
//
//	"github.com/gofiber/fiber/v2"
//	log "github.com/sirupsen/logrus"
//)
//
//func CreateFolder(c *fiber.Ctx) error {
//	folder := new(model.Folder)
//	userId := fmt.Sprintf("%s", c.Locals("id"))
//
//	if err := c.BodyParser(folder); err != nil {
//		log.Error(err)
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "Please review your input data",
//		})
//	}
//	folder.UserID = userId
//
//	if err := db.DB.Create(&folder).Error; err != nil {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "Unable to create folder",
//			Data:    err.Error(),
//		})
//	}
//
//	return c.JSON(types.GenericResponse{
//		Success: true,
//		Message: fmt.Sprintf("New folder created %s", folder.Name),
//	})
//}
