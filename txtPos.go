package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func drawPos(win *pixelgl.Window, txt *text.Text, pos pixel.Vec) {
	txt.Draw(win, pixel.IM.Moved(pos))
}

func centerPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerPos
}

// 画面中央の右隅にテキストを描画
func centerRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X+win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

// 画面中央の左隅にテキストを描画
func centerLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X-win.Bounds().Max.X/3,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerLeftPos
}

func topCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		win.Bounds().Min.Y+float64(winHSize/6),
	)
	return TopCenterPos
}

func topRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(winHSize/6),
	)
	return TopCenterPos
}

func topLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	centerLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(winHSize/6),
	)
	return centerLeftPos
}

func bottleCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-float64(winHSize/3),
	)
	return bottleCenterPos
}

func bottleRightPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X+win.Bounds().Max.X/3,
		-float64(winHSize/3),
	)
	return bottleLeftPos
}

func bottleLeftPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X-win.Bounds().Max.X/3,
		-float64(winHSize/3),
	)
	return bottleLeftPos
}

func lineCenterAlign(win *pixelgl.Window, lines []string, txt *text.Text, position string) {
	for _, line := range lines {
		centerX := win.Bounds().Center().Sub(txt.BoundsOf(line).Center()).X
		txt.Dot.X = centerX
		fmt.Fprintln(txt, line)
	}
	if position == "center" {
		drawPos(win, txt, centerPos(win, txt))
	}
}

func bottleRoundCenterPos(win *pixelgl.Window, txt *text.Text) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-win.Bounds().Center().Y/2,
	)
	return bottleCenterPos
}
