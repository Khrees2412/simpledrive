package util

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/khrees2412/simpledrive/types"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const AuthUserContextKey string = "auth"

var utilToken = Token{}

// UserFromContext extracts the user_id from context
func UserFromContext(ctx *fiber.Ctx) (string, error) {
	id := ctx.GetRespHeader(AuthUserContextKey)
	if &id == nil {
		return "", ctx.Status(fiber.StatusUnauthorized).JSON("unable to fetch user info from token")
	}

	return id, nil
}

// SecureAuth returns a middleware which secures all the private routes
func SecureAuth() func(*fiber.Ctx) error {
	jwtKey := os.Getenv("JWT_KEY")

	return func(c *fiber.Ctx) error {
		accessToken, err := utilToken.ExtractBearerToken(c.Get("Authorization"))
		if err != nil {
			logrus.Error("error while extracting token: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
				Success: false,
				Message: err.Error(),
			})
		}
		claims := new(types.Claims)
		token, err := jwt.ParseWithClaims(accessToken, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtKey), nil
			})
		if err != nil {
			logrus.Error("error while parsing claims: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if token.Valid {
			if claims.ExpiresAt.Unix() < time.Now().Unix() {
				return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
					Success: false,
					Message: "token expired",
				})
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return c.SendStatus(fiber.StatusForbidden)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return c.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
					Success: false,
					Message: "token expired or not yet active",
				})
			} else {
				// cannot handle this token
				return c.Status(fiber.StatusForbidden).JSON(types.GenericResponse{
					Success: false,
					Message: "unable to handle this token or invalid token",
				})
			}
		}

		c.Set(AuthUserContextKey, claims.UserId)
		return c.Next()
	}
}

// IsEmpty checks if a string is empty
func IsEmpty(str string) bool {
	if valid.HasWhitespaceOnly(str) && str != "" {
		return true
	}
	return false
}
