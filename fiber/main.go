package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main2() {
	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 10,
		ReadTimeout: time.Second * 10,
	})

	// Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 3000
	log.Fatal(app.Listen("localhost:3000"))
}