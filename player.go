package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type PlayerStatus struct {
	playerMaxHP  float64
	playerHP     float64
	playerOP     float64
	playerDP     float64
	playerMaxSP  float64
	playerSP     float64
	playerBaseSP float64
	playerGold   int
	playerJob    string
}

var (
	skillRect = pixel.R(0, 0, 0, 0)
)

func newPlayerStatus(MaxHP float64, HP float64, OP float64, DP float64, MaxSP float64, SP float64, BaseSP float64, Gold int, Job string) *PlayerStatus {
	return &PlayerStatus{MaxHP, HP, OP, DP, MaxSP, SP, BaseSP, Gold, Job}
}

func setPlayerSkillBar(win *pixelgl.Window, player *PlayerStatus) {
	// log.Println("win.Bounds().Max.X=", win.Bounds().Max.X)
	// log.Println("win.Bounds().Max.Y=", win.Bounds().Max.Y)
	rect := pixel.R(
		win.Bounds().Max.X-win.Bounds().Max.X/15,
		win.Bounds().Min.Y+win.Bounds().Max.Y/9,
		win.Bounds().Max.X-win.Bounds().Max.X/36,
		win.Bounds().Max.Y-win.Bounds().Max.Y/9,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Orange
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setPlayerSkillBarOut(win *pixelgl.Window, player *PlayerStatus) {
	//log.Println("maxsp", player.playerMaxSP, "SP=", player.playerSP, "Base=", player.playerBaseSP)
	rMinX := win.Bounds().Max.X - win.Bounds().Max.X/15
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/9
	rMaxX := win.Bounds().Max.X - win.Bounds().Max.X/36
	rMaxY := win.Bounds().Max.Y - win.Bounds().Max.Y/9
	if player.playerMaxSP <= player.playerSP {
		player.playerSP = 50
		skillRect = pixel.R(
			rMinX,
			rMinY+((win.Bounds().Max.Y-win.Bounds().Max.Y/9)-(win.Bounds().Min.Y+win.Bounds().Max.Y/9)),
			rMaxX,
			rMaxY,
		)
	} else if player.playerSP == 0 {
		skillRect = pixel.R(
			rMinX,
			rMinY,
			rMaxX,
			rMaxY,
		)
	} else if player.playerSP < player.playerMaxSP {
		skillRect = pixel.R(
			rMinX,
			rMinY+((win.Bounds().Max.Y-win.Bounds().Max.Y/9)-(win.Bounds().Min.Y+win.Bounds().Max.Y/9))*((player.playerMaxSP-(player.playerMaxSP-player.playerSP))/player.playerMaxSP),
			rMaxX,
			rMaxY,
		)
	}
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(skillRect.Min, skillRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setPlayerHPBar(win *pixelgl.Window, player *PlayerStatus) {
	rMinX := win.Bounds().Max.X - win.Bounds().Max.X/9
	rMinY := win.Bounds().Min.Y + win.Bounds().Max.Y/9
	rMaxX := win.Bounds().Max.X - win.Bounds().Max.X/15
	rMaxY := win.Bounds().Max.Y - win.Bounds().Max.Y/9

	hpRect := pixel.R(
		rMinX,
		rMinY,
		rMaxX,
		rMinY+(rMaxY-rMinY)*(player.playerHP/player.playerMaxHP), //-rMaxY*((player.playerHP)/player.playerMaxHP),
	)

	if player.playerHP <= 0 {
		hpRect = pixel.R(
			rMinX,
			rMinY,
			rMinX,
			rMinY,
		)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setPlayerHPBarOut(win *pixelgl.Window, player *PlayerStatus) {
	hpRect := pixel.R(
		win.Bounds().Max.X-win.Bounds().Max.X/9,
		win.Bounds().Min.Y+win.Bounds().Max.Y/9,
		win.Bounds().Max.X-win.Bounds().Max.X/15,
		win.Bounds().Max.Y-win.Bounds().Max.Y/9,
	)
	imd := imdraw.New(nil)
	imd.Color = colornames.Red
	imd.Push(hpRect.Min, hpRect.Max)
	imd.Rectangle(0)
	imd.Draw(win)
}

func setPlayerInf(win *pixelgl.Window, player *PlayerStatus) {
	setPlayerSkillBar(win, player)
	setPlayerSkillBarOut(win, player)
	setPlayerHPBarOut(win, player)
	setPlayerHPBar(win, player)
}
