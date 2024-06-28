package routes

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Reverse a string (text) query param
func Reverse(c *fiber.Ctx) error {
	text := c.Query("text")

	reversed := ""

	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}

	return c.JSON(fiber.Map{
		"reversed": reversed,
		"original": text,
	})
}

// Calculate the length of a string (text) query param
func Length(c *fiber.Ctx) error {
	text := c.Query("text")

	words := len(strings.Split(text, " "))
	characters := len(text)

	return c.JSON(fiber.Map{
		"length": fiber.Map{
			"words":      words,
			"characters": characters,
		},
		"text": text,
	})
}

// Return upercase of a string (text) query param
func Uppercase(c *fiber.Ctx) error {
	text := c.Query("text")

	uppercase := strings.ToUpper(text)

	return c.JSON(fiber.Map{
		"uppercase": uppercase,
		"original":  text,
	})
}

// Return lowercase of a string (text) query param
func Lowercase(c *fiber.Ctx) error {
	text := c.Query("text")

	lowercase := strings.ToLower(text)

	return c.JSON(fiber.Map{
		"lowercase": lowercase,
		"original":  text,
	})
}

// Replace all occurences of ?search with ?replace in ?text
func Replace(c *fiber.Ctx) error {
	text := c.Query("text")
	search := c.Query("search")
	replace := c.Query("replace")

	replaced := strings.Replace(text, search, replace, -1)

	return c.JSON(fiber.Map{
		"replaced": replaced,
		"original": text,
	})
}

// Generate a given number of lorem ipsum words
func LoremIpsum(c *fiber.Ctx) error {
	length, err := strconv.Atoi(c.Query("length"))
	if err != nil {
		return c.SendStatus(400)
	}

	lorem := "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur Excepteur sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim id est laborum"
	loremWords := strings.Split(lorem, " ")

	result := ""
	for i := 0; i < length; i++ {
		result += loremWords[i%len(loremWords)] + " "
	}

	return c.JSON(fiber.Map{
		"lorem":  result,
		"length": length,
	})
}
