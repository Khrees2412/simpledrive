package controllers

import (
	// "fmt"

	"fmt"
	db "simpledrive/database"
	"simpledrive/models"
	"simpledrive/types"
	"simpledrive/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	input := new(models.User)

	if err := c.BodyParser(input); err != nil {
		log.Error(err)
		return c.JSON(types.GenericResponse{
			Success: false,
			Message: "Please review your input",
		})
	}
	var user models.User

	if res := db.DB.Where(
		"email = ?", input.Email).Find(&user); res.RowsAffected <= 0 {
		return c.JSON(types.GenericResponse{
			Success: false,
			Message: "Invalid Credentials",
		})
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Error(err)
		return c.JSON(types.GenericResponse{
			Success: false,
			Message: "Invalid Credentials",
		})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := utils.GenerateTokens(user.ID)
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(types.GenericResponse{
		Success: true,
		Message: "Successfully generated tokens",
		Data:    fmt.Sprintf("access_token: %s, refresh_token: %s", accessToken, refreshToken),
	})
}
