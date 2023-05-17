package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	tempString = ""
)

func testMode(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	win.Clear(colornames.Mediumblue)
	//picMonster.Draw(win, pixel.IM)

	Txt.Clear()
	tempString = "RightPosition"
	fmt.Fprintln(Txt, tempString)
	drawPos(win, Txt, topRightPos(win, Txt, windowHeightSize))

	Txt.Clear()
	tempString = "LeftPosition"
	fmt.Fprintln(Txt, tempString)
	drawPos(win, Txt, topLeftPos(win, Txt, windowHeightSize))

	Txt.Clear()
	tempString = "bottleCenterPosition"
	fmt.Fprintln(Txt, tempString)
	drawPos(win, Txt, bottleCenterPos(win, Txt, windowHeightSize))

	Txt.Clear()
	tempString = "bottleRightPosition"
	fmt.Fprintln(Txt, tempString)
	drawPos(win, Txt, bottleRightPos(win, Txt, windowHeightSize))

	Txt.Clear()
	tempString = "bottleLeftPosition"
	fmt.Fprintln(Txt, tempString)
	drawPos(win, Txt, bottleLeftPos(win, Txt, windowHeightSize))
}
