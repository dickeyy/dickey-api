package routes

import (
	"context"

	"github.com/aidarkhanov/nanoid"
	"github.com/dickeyy/dickey-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type PasteRequest struct {
	Text string `json:"text"`
}

// CreatePaste handles POST /paste - creates a new paste and stores it in Redis
func CreatePaste(c *fiber.Ctx) error {
	var req PasteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Text == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Text field is required",
		})
	}

	// Generate a unique ID for the paste
	alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, _ := nanoid.Generate(alphabet, 12) // 778M IDs before 1% chance of collision

	// Store the paste in Redis
	ctx := context.Background()
	err := services.Redis.Set(ctx, "paste:"+id, req.Text, 0).Err()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to store paste",
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// GetPaste handles GET /paste/:id - retrieves a paste from Redis
func GetPaste(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	// Retrieve the paste from Redis
	ctx := context.Background()
	text, err := services.Redis.Get(ctx, "paste:"+id).Result()
	if err != nil {
		if err == redis.Nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Paste not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve paste",
		})
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"text": text,
	})
}
