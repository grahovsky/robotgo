package main

import (
	"fmt"
	"math/rand"

	"github.com/go-vgo/robotgo"
)

func explain() {

	robotgo.MilliSleep(1000)
	robotgo.ActiveName("nautilus")

	for {

		x, y := robotgo.GetMousePos()
		fmt.Println("pos: ", x, y)

		color := robotgo.GetPixelColor(x, y)
		fmt.Println("color: ", color)

		robotgo.MilliSleep(1043 - rand.Intn(100))

	}

}
