package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font"
)

func initializeSound(filePath string) (beep.StreamSeekCloser, func()) {
	soundFile := filePath
	f, err := os.Open(soundFile)
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		panic(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	return streamer, func() { _ = f.Close() }
}

func initializeWindow(windowHeightSize int) (*pixelgl.Window, pixelgl.WindowConfig) {
	windowWidthSize := (windowHeightSize / 16) * 9

	cfg := pixelgl.WindowConfig{Title: "MyPlatformer",
		Bounds:    pixel.R(0, 0, float64(windowHeightSize), float64(windowWidthSize)), //960,720
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	checkErrorPanic(err)
	win.SetTitle(fmt.Sprintf("%s | FPS: ", cfg.Title))

	return win, cfg
}

func initializeCanvas(sheet pixel.Picture, canvasSize int) (*pixelgl.Canvas, *imdraw.IMDraw) {
	bottomLCX := float64(-canvasSize / 2)
	bottomLCY := (-(float64(canvasSize/16) * 9) / 2)
	TopRCX := float64(+canvasSize / 2)
	TopRCY := (+(float64(canvasSize/16) * 9) / 2)

	canvas := pixelgl.NewCanvas(pixel.R(bottomLCX, bottomLCY, TopRCX, TopRCY))
	imd := imdraw.New(sheet)
	imd.Precision = 128

	return canvas, imd
}

func initializeText(face font.Face, color color.Color) *text.Text {
	basicAtlas := text.NewAtlas(face, text.ASCII)
	basicTxt := text.New(pixel.V(50, 500), basicAtlas)
	basicTxt.Color = color
	return basicTxt
}
