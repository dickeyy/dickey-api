package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// get current time
func Now(c *fiber.Ctx) error {
	now := time.Now()

	return c.JSON(fiber.Map{
		"time": fiber.Map{
			"unix": fiber.Map{
				"seconds": now.Unix(),
				"ms":      now.UnixMilli(),
			},
			"utc":  now.UTC().Format(time.RFC3339),
			"iso":  now.Format(time.RFC3339),
			"date": now.Format("2006-01-02"),
		},
	})
}
