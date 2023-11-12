package main

import (
    "fmt"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/stianeikeland/go-rpio/v4"
)

var (
    pin = rpio.Pin(10)
)

func main() {
    app := fiber.New()

    app.Get("/healthz", func(c *fiber.Ctx) error {
        return c.SendString("ok")
    })

    app.Get("/on", func(c *fiber.Ctx) error {
        if err := rpio.Open(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

	defer rpio.Close()
	pin.Output()

	pin.High()
	return c.SendString("Turning LED on")
    })

    app.Get("/oFF", func(c *fiber.Ctx) error {
        if err := rpio.Open(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

	defer rpio.Close()
	pin.Output()

	pin.Low()
	return c.SendString("Turning LED off")
    })
    app.Listen(":3000")
}
