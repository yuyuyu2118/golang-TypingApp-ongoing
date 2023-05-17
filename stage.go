package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type stageInf struct {
	stageNum int
}

func newStageInf(stageNum int) *stageInf {
	return &stageInf{stageNum}
}

var (
	stage1Button = pixel.Rect{}
)

func initStage(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Select play Stage")
	tempPosition = topCenterPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. VS Knight")
	tempPosition = centerLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	stage1Button = Txt.Bounds().Moved(tempPosition)
}
