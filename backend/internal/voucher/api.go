package voucher

import (
	"errors"
	"fmt"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/gofiber/fiber/v2"
)

func CreateVoucher(c *fiber.Ctx) error {
	req := new(Voucher)

	if err := c.BodyParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	validationError := api.ValidateRequest(req)
	if len(validationError) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": validationError,
		})
	}

	if err := Service.createVoucher(req); err != nil {
		if errors.Is(err, api.ErrPayload) {
			return c.Status(400).JSON(fiber.Map{
				"error": api.ErrPayload.Error(),
			})
		}
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"data": req,
	})
}

func GetVoucher(c *fiber.Ctx) error {
	req := new(Voucher)

	req.ID = c.Params("voucherId")

	if err := Service.getVoucher(req); err != nil {
		if errors.Is(err, api.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": api.ErrNotFound.Error(),
			})
		}
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": req,
	})
}

func GetVouchers(c *fiber.Ctx) error {
	param := c.Query("category")

	vouchers, err := Service.getVouchers(param)
	if err != nil {
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": vouchers,
	})
}

func UpdateVoucher(c *fiber.Ctx) error {
	req := new(Voucher)

	req.ID = c.Params("voucherId")

	updatedFields, err := api.GetUpdatedField(c.BodyRaw())
	if err != nil {
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	if err := c.BodyParser(req); err != nil {
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	if err := Service.updateVoucher(req, updatedFields); err != nil {
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": req,
	})
}

func DeleteVoucher(c *fiber.Ctx) error {
	req := new(Voucher)

	req.ID = c.Params("voucherId")

	if err := Service.deleteVoucher(req); err != nil {
		if errors.Is(err, api.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": api.ErrNotFound.Error(),
			})
		}
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}
