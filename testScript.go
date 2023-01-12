package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/go-vgo/robotgo"
)

func testScript() {

	// go func() {
	// 	select {
	// 	case <-robotgo.EventProcess(s):
	// 	}
	// }()

	robotgo.MilliSleep(2000)

	for {

		robotgo.MouseSleep = 70 - rand.Intn(5)
		robotgo.KeySleep = 70 - rand.Intn(5)

		robotgo.MoveSmooth(600+rand.Intn(500), 380+rand.Intn(5))
		robotgo.MoveSmoothRelative(0+rand.Intn(3), -10-rand.Intn(3))

		robotgo.Click("left")
		robotgo.MoveRelative(20+rand.Intn(5), 0+rand.Intn(3))
		robotgo.Click("right")
		robotgo.MoveRelative(30+rand.Intn(5), 0+rand.Intn(3))
		robotgo.Click("left")

		x, y := robotgo.GetMousePos()
		fmt.Println("pos: ", x, y)

		color := robotgo.GetPixelColor(x, y)
		fmt.Println("color: ", color)

		// robotgo.KeyTap("enter")
		// robotgo.KeyTap("escape")

		if color == "417ac9" {
			os.Exit(0)
		}

		robotgo.MilliSleep(1000 - rand.Intn(900))

	}

}
