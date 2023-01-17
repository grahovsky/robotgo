package main

import (
	"fmt"
	"math/rand"

	"github.com/go-vgo/robotgo"
)

type resolution struct {
	width int
	high  int
}

type coord struct {
	x int
	y int
}

type image struct {
	size int
	path string
}

func bitmap() {

	res := resolution{width: 1920, high: 1080}
	center := coord{x: res.width / 2, y: res.high / 2}

	exampleImage := image{size: 600, path: "pictures/example.png"}
	areaImage := image{size: 610}

	robotgo.ActiveName("code")

	robotgo.MilliSleep(100)

	// bit := robotgo.CaptureScreen(420, 200, 200, 200)
	// defer robotgo.FreeBitmap(bit)
	// robotgo.SaveBitmap(bit, "example.png")

	// other way
	robotgo.SaveCapture(exampleImage.path, center.x-exampleImage.size/2, center.y-exampleImage.size/2, exampleImage.size, exampleImage.size)

	// ?not working?
	// img := robotgo.ToImage(bit)
	// robotgo.SavePng(img, "test_1.png")

	exampleBitmap := robotgo.OpenBitmap(exampleImage.path)

	for {

		robotgo.MilliSleep(50 + rand.Intn(20))
		area := robotgo.CaptureScreen(center.x-areaImage.size/2, center.y-areaImage.size/2, areaImage.size, areaImage.size)

		x, y := robotgo.FindBitmap(exampleBitmap, area, 0.1)
		//x, y := robotgo.FindBitmap(exampleBitmap) // find on full screen

		if x >= 0 && y >= 0 {
			fmt.Println(x, y)
		} else {
			fmt.Println("not found")
		}

	}

}
