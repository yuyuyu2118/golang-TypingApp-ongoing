package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	job1Button = pixel.Rect{}
	job2Button = pixel.Rect{}
	job3Button = pixel.Rect{}
	// job4Button = pixel.Rect{}
	// job5Button = pixel.Rect{}
	// job6Button = pixel.Rect{}
)

var (
	tempPosition = pixel.Vec{}
)

func initJob(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Select your job")
	tempPosition = topCenterPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. Warrior")
	tempPosition = centerLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	job1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. Priest")
	tempPosition = centerPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	job2Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "3. Wizard")
	tempPosition = centerRightPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	job3Button = Txt.Bounds().Moved(tempPosition)
}
