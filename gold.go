package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func initPlayerGold(win *pixelgl.Window, Txt *text.Text, windowHeightSize int, player *PlayerStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Gold:", player.playerGold)
	tempPosition = topLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
}
