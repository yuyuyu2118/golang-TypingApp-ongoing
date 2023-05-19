package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	collectType = 0
	missType    = 0
	startTime   = time.Now()
)

func initBattleText(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) time.Duration {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "> ", words[score])
	drawPos(win, Txt, bottleRoundCenterPos(win, Txt, windowHeightSize))

	offset := Txt.Bounds().W()
	TxtOrigX := Txt.Dot.X
	spacing := 60.0
	if len(words)-score != 1 {
		Txt.Color = colornames.Darkgray
		offset := Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+1])
		drawPos(win, Txt, bottleRoundCenterPos(win, Txt, windowHeightSize).Add(pixel.V(offset+spacing, 0)))
		Txt.Dot.X = TxtOrigX
	}
	if !(len(words)-score == 2 || len(words)-score == 1) {
		Txt.Color = colornames.Gray
		offset += Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+2])
		drawPos(win, Txt, bottleRoundCenterPos(win, Txt, windowHeightSize).Add(pixel.V(offset+spacing*2, 0)))
	}
	//Txt.Dot.X = TxtOrigX

	Txt.Color = colornames.White
	Txt.Clear()
	fmt.Fprintln(Txt, "\n\n", "collectType = ", collectType, " missType = ", missType)
	drawPos(win, Txt, bottleRoundCenterPos(win, Txt, windowHeightSize))
	Txt.Dot.X = TxtOrigX

	//set Time+rule
	Txt.Clear()
	Txt.Color = colornames.White
	elapsed := time.Since(startTime)
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	drawPos(win, Txt, bottleLeftPos(win, Txt, windowHeightSize))
	return elapsed
}
