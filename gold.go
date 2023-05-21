package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func initPlayerGold(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Gold:", player.PlayerGold)
	xOffSet := 200.0
	yOffSet := win.Bounds().H() / 3
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
}
