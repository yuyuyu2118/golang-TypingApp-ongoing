package myPos

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func CenPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	centerPos := pixel.V(
		win.Bounds().Center().X-txtPos.X,
		win.Bounds().Center().Y-txtPos.Y,
	)
	return centerPos
}

// 画面中央の右隅にテキストを描画
func CenRigPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	centerRightPos := pixel.V(
		win.Bounds().Center().X-txtPos.X+win.Bounds().W()/3,
		win.Bounds().Center().Y-txtPos.Y,
	)
	return centerRightPos
}

// 画面中央の左隅にテキストを描画
func CenLefPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	centerLeftPos := pixel.V(
		win.Bounds().Center().X-txtPos.X-win.Bounds().W()/3,
		win.Bounds().Center().Y-txtPos.Y,
	)
	return centerLeftPos
}

func TopCenPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	TopCenterPos := pixel.V(
		win.Bounds().Center().X-txtPos.X,
		win.Bounds().Center().Y-txtPos.Y+win.Bounds().H()/2-txt.LineHeight,
	)
	return TopCenterPos
}

func TopRigPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	TopRightPos := pixel.V(
		win.Bounds().Center().X-txtPos.X+win.Bounds().W()/3,
		win.Bounds().Center().Y-txtPos.Y+win.Bounds().H()/2-txt.LineHeight,
	)
	return TopRightPos
}

func TopLefPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	TopLeftPos := pixel.V(
		win.Bounds().Center().X-txtPos.X-win.Bounds().W()/3,
		win.Bounds().Center().Y-txtPos.Y+win.Bounds().H()/2-txt.LineHeight,
	)
	return TopLeftPos
}

func BotCenPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	bottleCenterPos := pixel.V(
		win.Bounds().Center().X-txtPos.X,
		win.Bounds().Center().Y-txtPos.Y-win.Bounds().H()/3-txt.LineHeight*2,
	)
	return bottleCenterPos
}

// TODO: これと同じように、端を超えないように調整したい
func BotRigPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W(), bounds.H()/2)
	bottleLeftPos := pixel.V(
		win.Bounds().Center().X-txtPos.X+win.Bounds().W()/2-20,
		win.Bounds().Center().Y-txtPos.Y-win.Bounds().H()/3-txt.LineHeight*2,
	)
	return bottleLeftPos
}

func BotLefPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bounds := txt.Bounds()
	txtPos := pixel.V(bounds.W()/2, bounds.H()/2)
	bottleLeftPos := pixel.V(
		win.Bounds().Center().X-txtPos.X-win.Bounds().W()/3+20,
		win.Bounds().Center().Y-txtPos.Y-win.Bounds().H()/3-txt.LineHeight*2,
	)
	return bottleLeftPos
}
