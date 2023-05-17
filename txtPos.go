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

func centerPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	centerPos := pixel.V(
		win.Bounds().Center().Sub(txt.Bounds().Center()).X,
		win.Bounds().Center().Sub(txt.Bounds().Center()).Y,
	)
	return centerPos
}

func topCenterPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		win.Bounds().Min.Y+float64(size/6),
	)
	return TopCenterPos
}

func topRightPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Max.X-win.Bounds().Max.X/3,
		win.Bounds().Min.Y+float64(size/6),
	)
	return TopCenterPos
}

func topLeftPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	TopCenterPos := pixel.V(
		win.Bounds().Min.X,
		win.Bounds().Min.Y+float64(size/6),
	)
	return TopCenterPos
}

func bottleCenterPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	bottleCenterPos := pixel.V(
		win.Bounds().Max.X/2-txt.Bounds().Center().X,
		-float64(size/4),
	)
	return bottleCenterPos
}

func bottleLeftPos(win *pixelgl.Window, txt *text.Text, size int) pixel.Vec {
	bottleLeftPos := pixel.V(
		win.Bounds().Min.X,
		-float64(size/3),
	)
	return bottleLeftPos
}

func lineCenterAlign(win *pixelgl.Window, winSize int, lines []string, txt *text.Text, position string) {
	for _, line := range lines {
		centerX := win.Bounds().Center().Sub(txt.BoundsOf(line).Center()).X
		txt.Dot.X = centerX
		fmt.Fprintln(txt, line)
	}
	if position == "center" {
		drawPos(win, txt, centerPos(win, txt, winSize))
	}
}
