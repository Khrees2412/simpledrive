package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	FirstName string `json:"first_name" gorm:"unique"`
	LastName  string `json:"last_name" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
}

// Claims represent the structure of the JWT token
type Claims struct {
	jwt.RegisteredClaims
	ID uint `gorm:"primaryKey"`
}

// GenerateISOString generates a time string equivalent to Date.now().toISOString in JavaScript
func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// Base contains common columns for all tables
type Base struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Count     int    `json:"count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	// uuid.New() creates a new random UUID or panics.
	base.ID = uuid.NewString()

	// generate timestamps
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerateISOString()
	return nil
}
