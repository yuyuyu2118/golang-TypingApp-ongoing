package player

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

// TODO: 不要
func (player *PlayerStatus) InitPlayerGold(win *pixelgl.Window, Txt *text.Text) {
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, player.Job, " Gold:", player.Gold)
	tempPosition := myPos.TopLefPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)
}
