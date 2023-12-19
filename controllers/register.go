package controllers

import (
	"math/rand"

	db "simpledrive/database"
	"simpledrive/models"

	"simpledrive/utils"
	
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		log.Error(err)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Please review your input data",
		})
	}

	// validate if the email, username and password are in correct format
	errors := utils.ValidateRegister(user)
	if errors.Err {
		log.Error(errors)
		return c.JSON(errors)
	}

	// &models.User{Email: user.Email}
	if count := db.DB.Where("email = ?", user.Email).First(&user).RowsAffected; count > 0 {
		errors.Err, errors.Email = true, "Email is already registered"
	}
	// if count := database.Where(&models.User{Username: u.Username}).First(new(models.User)).RowsAffected; count > 0 {
	// 	errors.Err, errors.Username = true, "Username is already registered"
	// }
	if errors.Err {
		log.Error(errors)
		return c.JSON(errors)
	}

	// Hashing the password with a random salt
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(
		password,
		rand.Intn(bcrypt.MaxCost-bcrypt.MinCost)+bcrypt.MinCost,
	)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		log.Error(err)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong, please try again later. 😕",
		})
	}

	// setting up the authorization cookies
	accessToken, refreshToken := utils.GenerateTokens(user.ID)
	accessCookie, refreshCookie := utils.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
