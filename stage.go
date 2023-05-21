package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
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

func initStage(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Select play Stage")
	tempPosition = myPos.TopCenterPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. VS Knight")
	tempPosition = myPos.CenterLeftPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
	stage1Button = Txt.Bounds().Moved(tempPosition)
}

func stageClickEvent(win *pixelgl.Window, mousePos pixel.Vec, stage *stageInf) myGame.GameState {

	if stage1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		myGame.CurrentGS = myGame.PlayingScreen
		log.Println("PlayStage is VS knight")
		stage.stageNum = 1
	}
	log.Println("YourJob is", stage.stageNum)
	return myGame.CurrentGS
}
