package main

import (
    "fmt"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/stianeikeland/go-rpio/v4"
)

var (
    pin = rpio.Pin(19)
    enable_pin = rpio.Pin(18)
    blk_pin = rpio.Pin(4)
    grn_pin = rpio.Pin(17)
    red_pin = rpio.Pin(23)
    blu_pin = rpio.Pin(24)

    step_count = 4
    seq = [][]bool {{true,true,false,false},
		     {false,true,true,false},
		     {false,false,true,true},
		     {true,false,false,true}}
    )

func set_step(w1 bool, w2 bool, w3 bool, w4 bool) {
	if w1 == true {
		blk_pin.High()
	}else {
		blk_pin.Low()
	}

	if w2 == true{
		red_pin.High()
	}else {
		red_pin.Low()
	}

	if w3  == true{
		grn_pin.High()
	}else {
		grn_pin.Low()
	}

	if w4 == true{
		blu_pin.High()
	}else {
		blu_pin.Low()
	}
}

func forward(delay time.Duration, steps int){
	for i:=0; i<steps; i++{
		for j:=0; j<step_count; j++{
			set_step(seq[j][0], seq[j][1],seq[j][2], seq[j][3])
			time.Sleep(delay * time.Millisecond)
		}
	}
}

func backward(delay time.Duration, steps int){
	for i:=0; i<steps; i++{
		for j:=step_count; j>0; j--{
			set_step(seq[j][0], seq[j][1],seq[j][2], seq[j][3])
			time.Sleep(delay * time.Second)
		}
	}
}

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

    app.Get("/clockwise", func(c *fiber.Ctx) error {
        if err := rpio.Open(); err != nil {
            fmt.Println(err)
	    os.Exit(1)
	}
    enable_pin.Output()
    blk_pin.Output()
    grn_pin.Output()
    red_pin.Output()
    blu_pin.Output()
    enable_pin.High()

   forward(10, 1000)
       return c.SendString("Running stepper clockwise")
    })

    app.Listen(":3000")
}
