package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const pinGPIO = 10

var countInMin int

func pinRead(pin rpio.Pin) {
	res := pin.Read() // Read state from pin (High / Low)
	if res == 1 {
		//fmt.Println(time.Now().Local())
		countInMin++
	}
	fmt.Println("total count in this minute: " + strconv.Itoa(countInMin))
}
func inputRead() {
	err := rpio.Open()
	fmt.Println(err)
	pin := rpio.Pin(pinGPIO)
	fmt.Println("pin: " + strconv.Itoa(pinGPIO))

	pin.Input() // Input mode
	timeCount := 0
	countInMin = 0
	for {
		pinRead(pin)
		sleepSec(1)
		timeCount++
		if timeCount > 10 && countInMin > 0 { //for a minute
			sendRegister(time.Now().Local(), float64(countInMin))
			timeCount = 0
			countInMin = 0
		}
	}
}
