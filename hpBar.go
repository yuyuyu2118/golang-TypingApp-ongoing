package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type PlayerStatus struct {
	playerMaxHP float64
	playerHP    float64
	playerOP    float64
	playerDP    float64
	playerGold  int
	playerJob   string
}

func newPlayerStatus(MaxHP float64, HP float64, OP float64, DP float64, Gold int, Job string) *PlayerStatus {
	return &PlayerStatus{MaxHP, HP, OP, DP, Gold, Job}
}

type enemyStatus struct {
	enemyMaxHP float64
	enemyHP    float64
	enemyOP    float64
	enemyDP    float64
	enemyGold  int
}

func setEnemyHPBar(win *pixelgl.Window, scaledSize pixel.Vec, HP float64, MaxHP float64) {
	// Define rect parameters
	rectWidth := scaledSize.X * ((MaxHP - (MaxHP - HP)) * 0.01)
	rectHeight := 50.0 // Change this value to set the height of the rectangle
	rectPosY := win.Bounds().Center().Y - (scaledSize.Y / 2) - rectHeight

	// Create rect and draw it using a filled shape
	rect := pixel.R(
		win.Bounds().Center().X-(rectWidth/2),
		rectPosY,
		win.Bounds().Center().X+(rectWidth/2),
		rectPosY+rectHeight,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setEnemyHPBarOut(win *pixelgl.Window, scaledSize pixel.Vec) {
	// Define rect parameters
	rectWidth := scaledSize.X
	rectHeight := 50.0 // Change this value to set the height of the rectangle
	rectPosY := win.Bounds().Center().Y - (scaledSize.Y / 2) - rectHeight

	// Create rect and draw it using a filled shape
	rect := pixel.R(
		win.Bounds().Center().X-(rectWidth/2),
		rectPosY,
		win.Bounds().Center().X+(rectWidth/2),
		rectPosY+rectHeight,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}
