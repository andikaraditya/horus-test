package main

import (
	"log"
	"os"
	"time"

	"github.com/andikaraditya/horus-test/backend/internal/user"
	"github.com/andikaraditya/horus-test/backend/internal/voucher"
	"github.com/andikaraditya/horus-test/backend/internal/voucher/claim"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${queryParams} | ${error}\n",
		TimeFormat:    "15:04:05",
		TimeZone:      "Local",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: false,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", user.CreateUser)
	app.Post("/login", user.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Get("/restricted", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Login!")
	})

	app.Post("/vouchers", voucher.CreateVoucher)
	app.Get("/vouchers", voucher.GetVouchers)
	app.Get("/vouchers/:voucherId", voucher.GetVoucher)
	app.Put("/vouchers/:voucherId", voucher.UpdateVoucher)
	app.Delete("/vouchers/:voucherId", voucher.DeleteVoucher)
	app.Post("/vouchers/:voucherId/claim", claim.CreateClaimVoucher)

	app.Get("/claims", claim.GetClaimVouchers)
	app.Get("/claims/summary", claim.GetClaimSummary)
	app.Get("/claims/:claimVoucherId", claim.GetClaimVoucher)
	app.Delete("/claims/:claimVoucherId", claim.DeleteClaimVoucher)

	log.Fatal(app.Listen(":3000"))
}
