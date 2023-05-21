package myPos

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func DrawPos(win *pixelgl.Window, txt *text.Text, pos pixel.Vec) {
	txt.Draw(win, pixel.IM.Moved(pos))
}

func CenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerPos
}

// 画面中央の右隅にテキストを描画
func CenterRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X+win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

// 画面中央の左隅にテキストを描画
func CenterLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X-win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

func TopCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return TopCenterPos
}

func TopRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return TopCenterPos
}

func TopLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(WinHSize/6),
	)
	return centerLeftPos
}

func BottleCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-float64(WinHSize/3),
	)
	return bottleCenterPos
}

func BottleRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		-float64(WinHSize/3),
	)
	return bottleLeftPos
}

func BottleLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		-float64(WinHSize/3),
	)
	return bottleLeftPos
}

func LineCenterAlign(win *pixelgl.Window, lines []string, txt *text.Text, position string) {
	for _, line := range lines {
		centerX := win.Bounds().Center().Sub(txt.BoundsOf(line).Center()).X
		txt.Dot.X = centerX
		fmt.Fprintln(txt, line)
	}
	if position == "center" {
		DrawPos(win, txt, CenterPos(win, txt))
	}
}

func BottleRoundCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-win.Bounds().Center().Y/2,
	)
	return bottleCenterPos
}
