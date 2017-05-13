package main

import (
	"math/rand"
	"time"
)

func random() float64 {
	rand.Seed(time.Now().Unix())
	val := rand.Float64() * 10
	return float64(val)
}

func fakeInputData() {
	for { //loop to send Register eaery X seconds
		r := random()
		sended := sendRegister(time.Now().Local(), r)
		if sended == false {
			date := time.Now().Local().Format("2006-01-02 15:04:05.000Z")
			c.Red("[ERROR] " + date + " Can not send register data")
		}
		//sleeps 10 seconds
		sleepSec(5)
	}
}
