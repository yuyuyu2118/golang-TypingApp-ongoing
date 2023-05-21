package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var (
	collectType = 0
	missType    = 0
	startTime   = time.Now()
)

func initBattleText(win *pixelgl.Window, Txt *text.Text) time.Duration {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "> ", words[score])
	myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

	offset := Txt.Bounds().W()
	TxtOrigX := Txt.Dot.X
	spacing := 60.0
	if len(words)-score != 1 {
		Txt.Color = colornames.Darkgray
		offset := Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+1])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
		Txt.Dot.X = TxtOrigX
	}
	if !(len(words)-score == 2 || len(words)-score == 1) {
		Txt.Color = colornames.Gray
		offset += Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+2])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
	}
	//Txt.Dot.X = TxtOrigX

	Txt.Color = colornames.White
	Txt.Clear()
	fmt.Fprintln(Txt, "\n\n", "collectType = ", collectType, " missType = ", missType)
	myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))
	Txt.Dot.X = TxtOrigX

	//set Time+rule
	Txt.Clear()
	Txt.Color = colornames.White
	elapsed := time.Since(startTime)
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))
	return elapsed
}
