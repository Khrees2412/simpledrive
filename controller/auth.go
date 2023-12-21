package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/simpledrive/service"
	"github.com/khrees2412/simpledrive/types"
	"github.com/khrees2412/simpledrive/util"
)

type IAuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type authController struct {
	authService service.IAuthService
}

// NewAuthController instantiates Auth Controller
func NewAuthController() IAuthController {
	return &authController{
		authService: service.NewAuthService(),
	}
}

func (ctl *authController) RegisterRoutes(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/register", ctl.Register)
	auth.Post("/login", ctl.Login)
}

func (ctl *authController) Register(ctx *fiber.Ctx) error {
	var body types.RegisterRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	res, err := ctl.authService.Register(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "User successfully registered",
		Data:    res,
	})

}
func (ctl *authController) Login(ctx *fiber.Ctx) error {
	var body types.LoginRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	res, err := ctl.authService.Login(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "User successfully logged in",
		Data:    res,
	})
}
