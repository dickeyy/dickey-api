package routes

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Calculates if n is prime and returns the divisors
func Prime(c *fiber.Ctx) error {
	// cast to int
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.SendStatus(400)
	}

	isPrime := true
	divisors := []int{}
	msg := "is"

	// Handle special cases
	if n < 2 {
		isPrime = false
		msg = "is not"
	}

	// Find divisors and check primality
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			// If we find a divisor other than 1 and n, it's not prime
			if i != 1 && i != n {
				isPrime = false
				msg = "is not"
			}
		}
	}

	return c.JSON(fiber.Map{
		"isPrime":  isPrime,
		"divisors": divisors,
		"message":  fmt.Sprintf("%d %s prime", n, msg),
	})
}

// Calculates the first n digits in the Fibonacci Sequence
func Fibonacci(c *fiber.Ctx) error {
	// cast to int
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.SendStatus(400)
	}

	// constraints
	if n < 0 {
		return c.SendStatus(400)
	} else if n == 0 {
		return c.JSON(fiber.Map{
			"sequence": []int{0},
		})
	} else if n > 1476 {
		return c.SendStatus(400)
	}

	sequence := []int{}
	sequence = append(sequence, 0, 1) // add first two numbers

	for i := 2; i <= n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}

	return c.JSON(fiber.Map{
		"sequence": sequence,
		"messge":   fmt.Sprintf("First %d numbers in the Fibonacci sequence", n),
	})
}

// Generate random-number between query params min and max
func RandomNumber(c *fiber.Ctx) error {
	minStr := c.Query("min")
	maxStr := c.Query("max")

	min := 0
	max := 1000

	if minStr != "" || maxStr != "" {
		min, _ = strconv.Atoi(minStr)
		max, _ = strconv.Atoi(maxStr)
	}

	if min > max {
		return c.SendStatus(400)
	}

	rn := rand.Intn(max-min) + min

	return c.JSON(fiber.Map{
		"randomNumber": rn,
		"message":      fmt.Sprintf("Random number between %d and %d", min, max),
		"range":        []int{min, max},
	})
}

// Calculate n! (factorial)
func Factorial(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	} else if n > 200 {
		return c.SendStatus(400)
	}

	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return c.JSON(fiber.Map{
		"number":  factorial,
		"message": fmt.Sprintf("Factorial of %d is %d", n, factorial),
	})
}

// Calculate the square root of a number
func SquareRoot(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	sqrt := math.Sqrt(float64(n))

	return c.JSON(fiber.Map{
		"number":  sqrt,
		"message": fmt.Sprintf("Square root of %d is %.2f", n, sqrt),
	})
}

// Calculate the absolute value of a number
func AbsoluteValue(c *fiber.Ctx) error {
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	abs := math.Abs(float64(n))

	return c.JSON(fiber.Map{
		"number":  abs,
		"message": fmt.Sprintf("Absolute value of %d is %.0f", n, abs),
	})
}

// Round n to the nearest integer
func Round(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	rounded := math.Round(n)

	return c.JSON(fiber.Map{
		"number":  rounded,
		"message": fmt.Sprintf("Rounded value of %f is %.0f", n, rounded),
	})
}

// Ceil n to the nearest integer
func Ceil(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	ceiled := math.Ceil(n)

	return c.JSON(fiber.Map{
		"number":  ceiled,
		"message": fmt.Sprintf("Ceiled value of %f is %.0f", n, ceiled),
	})
}

// Floor n to the nearest integer
func Floor(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	floored := math.Floor(n)

	return c.JSON(fiber.Map{
		"number":  floored,
		"message": fmt.Sprintf("Floored value of %f is %.0f", n, floored),
	})
}

// Calculate sin value of n
func Sin(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	sin := math.Sin(n * math.Pi / 180)

	return c.JSON(fiber.Map{
		"number":  sin,
		"message": fmt.Sprintf("Sin value of %f is %.2f", n, sin),
	})
}

// Calculate cos value of n
func Cos(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	cos := math.Cos(n * math.Pi / 180)

	return c.JSON(fiber.Map{
		"number":  cos,
		"message": fmt.Sprintf("Cos value of %f is %.2f", n, cos),
	})
}

// Calculate tan value of n
func Tan(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	tan := math.Tan(n * math.Pi / 180)

	return c.JSON(fiber.Map{
		"number":  tan,
		"message": fmt.Sprintf("Tan value of %f is %.2f", n, tan),
	})
}

// Calculate the logarithm of a number
func Log(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	log := math.Log(n)

	return c.JSON(fiber.Map{
		"number":  log,
		"message": fmt.Sprintf("Log value of %f is %.2f", n, log),
	})
}

// Calculate the log base 10 of a number
func Log10(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	log10 := math.Log10(n)

	return c.JSON(fiber.Map{
		"number":  log10,
		"message": fmt.Sprintf("Log10 value of %f is %.2f", n, log10),
	})

}

// Calculate the log base 2 of a number
func Log2(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	log2 := math.Log2(n)

	return c.JSON(fiber.Map{
		"number":  log2,
		"message": fmt.Sprintf("Log2 value of %f is %.2f", n, log2),
	})
}

// Calculate the exp value of a number
func Exp(c *fiber.Ctx) error {
	n, err := strconv.ParseFloat(c.Params("n"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if n < 0 {
		return c.SendStatus(400)
	}

	exp := math.Exp(n)

	return c.JSON(fiber.Map{
		"number":  exp,
		"message": fmt.Sprintf("Exp value of %f is %.2f", n, exp),
	})
}

// Calculate the pow of ?base and ?exponent
func Pow(c *fiber.Ctx) error {
	base, err := strconv.ParseFloat(c.Query("base"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	exponent, err := strconv.ParseFloat(c.Query("exponent"), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	if base < 0 {
		return c.SendStatus(400)
	}
	if exponent < 0 {
		return c.SendStatus(400)
	}

	pow := math.Pow(base, exponent)

	return c.JSON(fiber.Map{
		"number":  pow,
		"message": fmt.Sprintf("Pow value of %f to the power of %f is %.2f", base, exponent, pow),
	})
}
