package main

import (
	"os"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {

	register_exit_keydown()

	s := robotgo.EventStart()
	robotgo.EventProcess(s)

	//explain()
	// testScript()
	//shootScript()
	bitmap()

}

func register_exit_keydown() {
	robotgo.EventHook(hook.KeyDown, []string{"shift", "z"}, func(e hook.Event) {
		robotgo.EventEnd()
		os.Exit(0)
	})
}
