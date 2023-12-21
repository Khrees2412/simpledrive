package types

import "github.com/golang-jwt/jwt/v4"

type (
	// Claims represent the structure of the JWT token
	Claims struct {
		Email  string `json:"email"`
		UserId string `json:"user_id"`
		jwt.RegisteredClaims
	}
	GenericResponse struct {
		Success bool        `json:"success"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
	RegisterRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required"`
		Password  string `json:"password" validate:"required"`
	}
	RegisterResponse struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	LoginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct {
		Token     *TokenResponse `json:"token"`
		FirstName string         `json:"first_name"`
		LastName  string         `json:"last_name"`
		Email     string         `json:"-"`
	}
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresAt   int64  `json:"expires_at"`
		Issuer      string `json:"issuer"`
	}

	CreateFolderRequest struct {
		Name string `json:"name" validate:"required"`
	}
)
