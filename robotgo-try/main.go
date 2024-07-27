package main

import (
	"github.com/go-vgo/robotgo"
)

// see https://github.com/go-vgo/robotgo?tab=readme-ov-file#examples
func main() {
	robotgo.MouseSleep = 100

	robotgo.ScrollDir(10, "up")
	robotgo.ScrollDir(20, "right")

	robotgo.Scroll(0, -10)
	robotgo.Scroll(100, 0)

	robotgo.MilliSleep(100)
	robotgo.ScrollSmooth(-10, 6)

	robotgo.Move(10, 20)
	robotgo.MoveRelative(0, -10)
	robotgo.DragSmooth(10, 10)
}
