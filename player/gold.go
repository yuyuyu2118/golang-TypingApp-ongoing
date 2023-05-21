package player

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

func (player *PlayerStatus) InitPlayerGold(win *pixelgl.Window, Txt *text.Text) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Gold:", player.Gold)
	xOffSet := 200.0
	yOffSet := win.Bounds().H() / 3
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	Txt.Draw(win, tempPosition)
}
