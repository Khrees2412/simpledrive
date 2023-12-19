package controllers

import (
	db "simpledrive/database"
	"simpledrive/models"
	"simpledrive/types"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) error {
	id := c.Locals("id")

	u := new(models.User)
	if res := db.DB.Where("uuid = ?", id).First(&u); res.RowsAffected <= 0 {
		return c.JSON(types.GenericResponse{
			Success: false,
			Message: "User not found",
			Data:    u,
		})
	}

	return c.JSON(u)
}
