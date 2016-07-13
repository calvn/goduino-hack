package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/joho/godotenv"
)

func main() {
	// Get credentials
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	// Initialize twitter client and stream
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	stream := api.UserStream(nil)

	// Arduino setup
	gbot := gobot.NewGobot()
	board := firmata.NewFirmataAdaptor("arduino", os.Args[1])

	// digital devices
	button := gpio.NewButtonDriver(board, "button", "2")
	blue := gpio.NewLedDriver(board, "blue", "3")

	work := func() {
		go func() {
			for streamObj := range stream.C {
				switch t := streamObj.(type) {
				case anaconda.Tweet:
					blue.On()
				default:
					log.Println("unhandled type %T", t)
				}
			}
		}()

		gobot.On(button.Event(gpio.Release), func(data interface{}) {
			fmt.Println("Off!")
			blue.Off()
		})
	}

	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
