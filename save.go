package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var (
	save1Button = pixel.Rect{}
	save2Button = pixel.Rect{}
)

func initSave(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Do you want to save?")
	tempPosition = topCenterPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. Yes")
	tempPosition = centerLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	save1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. No")
	tempPosition = centerLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	save2Button = Txt.Bounds().Moved(tempPosition)
}

func saveClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState, player *PlayerStatus) GameState {
	//TODO ページを作成したら追加
	if equip1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {

		currentGameState = GoToScreen
		log.Println("Save Done!")
	} else if equip1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = GoToScreen
		log.Println("saveScreen->GoToScreen")
	}
	return currentGameState
}
