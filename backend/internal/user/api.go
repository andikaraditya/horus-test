package user

import (
	"errors"
	"fmt"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	req := new(User)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	validationError := api.ValidateRequest(req)
	if len(validationError) > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"error": validationError,
		})
	}

	err := Service.CreateUser(req)
	if err != nil {
		if errors.Is(err, api.ErrPayload) {
			return ctx.Status(400).JSON(fiber.Map{
				"error": "email already exists",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}

func Login(ctx *fiber.Ctx) error {
	req := new(LoginRequest)

	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	validationError := api.ValidateRequest(req)
	if len(validationError) > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"error": validationError,
		})
	}

	token, err := Service.Login(req)
	if err != nil {
		fmt.Printf("Error: %v", err)
		if errors.Is(err, api.ErrPayload) {
			return ctx.Status(400).JSON(fiber.Map{
				"error": "incorrect password or email",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
