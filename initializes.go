package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font"
)

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

// func initializeCanvas(sheet pixel.Picture, canvasSize int) (*pixelgl.Canvas, *imdraw.IMDraw) {
// 	bottomLCX := float64(-canvasSize / 2)
// 	bottomLCY := (-(float64(canvasSize/16) * 9) / 2)
// 	TopRCX := float64(+canvasSize / 2)
// 	TopRCY := (+(float64(canvasSize/16) * 9) / 2)

// 	canvas := pixelgl.NewCanvas(pixel.R(bottomLCX, bottomLCY, TopRCX, TopRCY))
// 	imd := imdraw.New(sheet)
// 	imd.Precision = 128

// 	return canvas, imd
// }

func initializeText(face font.Face, color color.Color) *text.Text {
	basicAtlas := text.NewAtlas(face, text.ASCII)
	basicTxt := text.New(pixel.V(50, 500), basicAtlas)
	basicTxt.Color = color
	return basicTxt
}

func initializeAnyText(fontPath string, size int, color color.Color) *text.Text {
	face, _ := loadTTF(fontPath, float64(size))
	return initializeText(face, color)
}

// func initializeSound(filePath string) (beep.StreamSeekCloser, func()) {
// 	soundFile := filePath
// 	f, err := os.Open(soundFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	streamer, format, err := mp3.Decode(f)
// 	if err != nil {
// 		f.Close()
// 		panic(err)
// 	}
// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

// 	return streamer, func() { _ = f.Close() }
// }

func initScreenInformation(win *pixelgl.Window, Txt *text.Text, windowHeightSize int, player *PlayerStatus, currentGameState GameState) {

	switch currentGameState {
	case GoToScreen:
		initGoToScreen(win, Txt, windowHeightSize)
	case StageSelect:
		initStageSlect(win, Txt, windowHeightSize)
	case TownScreen:
		initTownScreen(win, Txt, windowHeightSize)
	case WeaponShop:
		initWeaponShop(win, Txt, windowHeightSize)
	case ArmorShop:
		initArmorShop(win, Txt, windowHeightSize)
	case AccessoryShop:
		initAccessoryShop(win, Txt, windowHeightSize)
	case EquipmentScreen:
		initEquipmentScreen(win, Txt, windowHeightSize)
	case JobSelect:
		initJobSelect(win, Txt, windowHeightSize)
	case SaveScreen:
		initSaveScreen(win, Txt, windowHeightSize)
	case PlayingScreen:
		initPlayingScreen(win, Txt, windowHeightSize)
	}

	initPlayerGold(win, Txt, windowHeightSize, player)
	initPlayerJob(win, Txt, windowHeightSize, player)
	initPlayerStatus(win, Txt, windowHeightSize, player)
	//initPlayerEquipment(win, Txt, windowHeightSize, player)
}
