package myGame

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var (
	tempString = ""
)

func TestMode(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Mediumblue)
	//picMonster.Draw(win, pixel.IM)

	Txt.Clear()
	tempString = "RightPosition"
	fmt.Fprintln(Txt, tempString)
	myPos.DrawPos(win, Txt, myPos.TopRightPos(win, Txt))

	Txt.Clear()
	tempString = "LeftPosition"
	fmt.Fprintln(Txt, tempString)
	myPos.DrawPos(win, Txt, myPos.TopLeftPos(win, Txt))

	Txt.Clear()
	tempString = "bottleCenterPosition"
	fmt.Fprintln(Txt, tempString)
	myPos.DrawPos(win, Txt, myPos.BottleCenterPos(win, Txt))

	Txt.Clear()
	tempString = "bottleRightPosition"
	fmt.Fprintln(Txt, tempString)
	myPos.DrawPos(win, Txt, myPos.BottleRightPos(win, Txt))

	Txt.Clear()
	tempString = "bottleLeftPosition"
	fmt.Fprintln(Txt, tempString)
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))
}
