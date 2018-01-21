package main

import (
	"log"
	"strconv"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type LedController struct {
	rpi  *raspi.Adaptor
	Leds map[string][]*gpio.LedDriver
}

func NewLedController(adaptor *raspi.Adaptor) *LedController {
	return &LedController{Leds: make(map[string][]*gpio.LedDriver), rpi: adaptor}
}

func (c LedController) ActivateLeds(ledType string) {
	log.Printf("activating leds with type '%v'\n", ledType)
	for _, led := range c.Leds[ledType] {
		led.On()
	}
}

func (c LedController) DeactivateLeds(ledType string) {
	log.Printf("deactivating leds with type '%v'\n", ledType)
	for _, led := range c.Leds[ledType] {
		led.Off()
	}
}

func (c LedController) ClearLeds() {
	log.Println("clearing leds")
	for _, leds := range c.Leds {
		for _, led := range leds {
			led.Off()
		}
	}
}

func (c LedController) AddLed(ledType string, ledNumber int) {
	log.Printf("registering %v pin %v.\n", ledType, ledNumber)
	c.Leds[ledType] = append(c.Leds[ledType], gpio.NewLedDriver(rpi, strconv.Itoa(ledNumber)))
}
