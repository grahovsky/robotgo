package main

import (
	"math/rand"
	"reflect"

	"github.com/go-vgo/robotgo"
)

func shootScript() {

	targetsCoord := [][]int{
		{400, 200},
		{20, 720},
		{800, 400},
		{1200, 600},
		{935, 1070},
		{1600, 800},
	}

	robotgo.MilliSleep(2000)
	findedCoord := []int{0, 0}

	for {

		robotgo.MouseSleep = 70 - rand.Intn(10)

		for _, targetCoord := range targetsCoord {

			color := robotgo.GetPixelColor(targetCoord[0], targetCoord[1])
			if color == "417ac9" || color == "333333" {
				if !reflect.DeepEqual(targetCoord, findedCoord) {
					findedCoord = targetCoord
					robotgo.MoveSmooth(findedCoord[0]+rand.Intn(3), findedCoord[1]+rand.Intn(10), 0.5, 0.2)
					break
				}
			}

		}

		robotgo.MilliSleep(100 - rand.Intn(40))

	}

}
