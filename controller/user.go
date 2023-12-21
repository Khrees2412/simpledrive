package controller

//
//import (
//	db "github.com/khrees2412/simpledrive/database"
//	"github.com/khrees2412/simpledrive/model"
//	"github.com/khrees2412/simpledrive/types"
//
//	"github.com/gofiber/fiber/v2"
//)
//
//func GetUserData(c *fiber.Ctx) error {
//	id := c.Locals("id")
//
//	u := new(model.User)
//	if res := db.DB.Where("id = ?", id).First(&u); res.RowsAffected <= 0 {
//		return c.JSON(types.GenericResponse{
//			Success: false,
//			Message: "User not found",
//			Data:    u,
//		})
//	}
//
//	return c.JSON(u)
//}
