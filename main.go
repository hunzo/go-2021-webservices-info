package main

import (
	"time"

	_ "time/tzdata"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"_golang_api": "information services",
			"_timestamp_": time.Now().Format("Jan _2 15:04:05"),
			":8001":       "ldap password reset",
			":8002":       "msgraph sendmail service (gin)",
			":8003":       "msgraph password reset",
			":8004":       "radius authentication (uapp)",
			":8005":       "certificate openssl",
			":8006":       "msgraph sendmail services (fiber)",
		})
	})
	app.Listen(":8080")
}
