package main

import (
	"os"

	"github.com/dickeyy/dickey-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var app = fiber.New()

func init() {
	// load .env file
	if err := godotenv.Load(".env.local"); err != nil {
		// log.Panic("No .env file found")
		return
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return "0.0.0.0:" + port
}

// main function
func main() {
	// Configure CORS to allow requests from any origin
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// WebSocket middleware
	app.Use("/spotify/current-track/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("docs at https://docs.kyle.so")
	})

	// math routes
	app.Get("/math/prime/:n", routes.Prime)
	app.Get("/math/fibonacci/:n", routes.Fibonacci)
	app.Get("/math/random-number", routes.RandomNumber)
	app.Get("/math/factorial/:n", routes.Factorial)
	app.Get("/math/sqrt/:n", routes.SquareRoot)
	app.Get("/math/abs/:n", routes.AbsoluteValue)
	app.Get("/math/round/:n", routes.Round)
	app.Get("/math/ceil/:n", routes.Ceil)
	app.Get("/math/floor/:n", routes.Floor)
	app.Get("/math/sin/:n", routes.Sin)
	app.Get("/math/cos/:n", routes.Cos)
	app.Get("/math/tan/:n", routes.Tan)
	app.Get("/math/log/:n", routes.Log)
	app.Get("/math/log10/:n", routes.Log10)
	app.Get("/math/log2/:n", routes.Log2)
	app.Get("/math/exp/:n", routes.Exp)
	app.Get("/math/pow", routes.Pow)

	// text routes
	app.Get("/text/reverse", routes.Reverse)
	app.Get("/text/length", routes.Length)
	app.Get("/text/uppercase", routes.Uppercase)
	app.Get("/text/lowercase", routes.Lowercase)
	app.Get("/text/replace", routes.Replace)
	app.Get("/text/lorem", routes.LoremIpsum)

	// time routes
	app.Get("/time/now", routes.Now)

	// spotify routes
	app.Get("/spotify/current-track", routes.GetCurrentTrack)
	app.Get("/spotify/current-track/ws", websocket.New(routes.HandleSpotifyWebSocket))

	// start server
	err := app.Listen(getPort())
	if err != nil {
		log.Panic(err)
	}
}
