package myPos

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func DrawPos(win *pixelgl.Window, txt *text.Text, pos pixel.Vec) {
	txt.Draw(win, pixel.IM.Moved(pos))
}

func TopCenterPos(win *pixelgl.Window, txt *text.Text, winHSize int) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		win.Bounds().Min.Y+float64(winHSize/6),
	)
	return TopCenterPos
}
