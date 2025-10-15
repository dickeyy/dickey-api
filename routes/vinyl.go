package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type DiscogsResponse struct {
	Releases []interface{} `json:"releases"`
}

func GetVinylCollection(c *fiber.Ctx) error {
	token := os.Getenv("DISCOGS_TOKEN")
	if token == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "DISCOGS_TOKEN not set",
		})
	}

	url := "https://api.discogs.com/users/kdickey/collection/folders/0/releases?per_page=100"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request",
		})
	}

	req.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", token))
	req.Header.Set("User-Agent", "dickey-api/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch vinyl collection",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": "Failed to fetch vinyl collection from Discogs",
		})
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response body",
		})
	}

	var data DiscogsResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse Discogs response",
		})
	}

	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")

	return c.JSON(data.Releases)
}
