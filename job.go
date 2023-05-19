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

func initJob(win *pixelgl.Window, Txt *text.Text) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Select your job")
	tempPosition = topCenterPos(win, Txt)
	drawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. Warrior")
	tempPosition = centerLeftPos(win, Txt)
	drawPos(win, Txt, tempPosition)
	job1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. Priest")
	tempPosition = centerPos(win, Txt)
	drawPos(win, Txt, tempPosition)
	job2Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "3. Wizard")
	tempPosition = centerRightPos(win, Txt)
	drawPos(win, Txt, tempPosition)
	job3Button = Txt.Bounds().Moved(tempPosition)
}

func jobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *PlayerStatus) GameState {

	if job1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = GoToScreen
		player.playerJob = "Warrior"
	} else if job2Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = GoToScreen
		player.playerJob = "Priest"
	} else if job3Button.Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentGameState = GoToScreen
		player.playerJob = "Wizard"
	}
	log.Println("YourJob is", player.playerJob)
	return currentGameState
}

func initPlayerJob(win *pixelgl.Window, Txt *text.Text, player *PlayerStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, player.playerJob)
	xOffSet := 0.0
	yOffSet := win.Bounds().H() / 3
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
}
