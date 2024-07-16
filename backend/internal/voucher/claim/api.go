package claim

import (
	"errors"
	"fmt"

	"github.com/andikaraditya/horus-test/backend/internal/api"
	"github.com/gofiber/fiber/v2"
)

func CreateClaimVoucher(c *fiber.Ctx) error {
	req := new(ClaimVoucher)

	req.VoucherId = c.Params("voucherId")

	if err := Service.createClaimVoucher(req); err != nil {
		fmt.Println("Error: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"data": req,
	})
}

func GetClaimVoucher(c *fiber.Ctx) error {
	req := new(ClaimVoucher)

	req.ID = c.Params("claimVoucherId")

	if err := Service.getClaimVoucher(req); err != nil {
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

func GetClaimVouchers(c *fiber.Ctx) error {
	vouchers, err := Service.getClaimVouchers()
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

func DeleteClaimVoucher(c *fiber.Ctx) error {
	req := new(ClaimVoucher)

	req.ID = c.Params("claimVoucherId")

	if err := Service.deleteClaimVoucher(req); err != nil {
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

func GetClaimSummary(c *fiber.Ctx) error {
	summary, err := Service.getClaimSummary()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": summary,
	})
}
