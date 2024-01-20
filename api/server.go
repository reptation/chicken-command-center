package main

import (
    "fmt"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/stianeikeland/go-rpio/v4"
)

var (
    // https://github.com/stianeikeland/go-rpio/blob/d8d85b35367c123e3bd3cb16c7fe0aa1f50bb19f/rpio.go#L304C48-L304C48
    // PWM is only on pins 12, 13, 18, 19
    pin = rpio.Pin(19)
    enable_pin = rpio.Pin(18)

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
	pin.Pwm()
	enable_pin.Output()
	enable_pin.High()
	pin.Freq(64000)
        pin.DutyCycle(0,32)

	// five times smoothly fade in and out
        for i := 0; i < 50; i++ {
                for i := uint32(0); i < 32; i++ { // increasing brightness
                        pin.DutyCycle(i, 32)
                        time.Sleep(time.Second/32)
                }
                for i := uint32(32); i > 0; i-- { // decreasing brightness
                        pin.DutyCycle(i, 32)
                        time.Sleep(time.Second/32)
                }
        }

	return c.SendString("Turning LED on")
    })

    app.Get("/off", func(c *fiber.Ctx) error {
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
