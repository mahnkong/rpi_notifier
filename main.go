package main

import (
	"fmt"
	"os"
	"strconv"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	rpi           = raspi.NewAdaptor()
	ledController = NewLedController(rpi)
	server        = NewWebServer(ledController)
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %v config.json\n", os.Args[0])
		os.Exit(1)
	}

	conf := NewConfigFromFile(os.Args[1])

	for color, ledNumbers := range conf.Leds {
		for _, ledNumber := range ledNumbers {
			ledController.AddLed(color, ledNumber)
		}
	}

	button := gpio.NewButtonDriver(rpi, strconv.Itoa(conf.ButtonPin))
	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			ledController.ClearLeds()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{rpi},
		[]gobot.Device{button},
		work,
	)

	go func() {
		server.Run(conf.Port)
	}()
	robot.Start()
}
