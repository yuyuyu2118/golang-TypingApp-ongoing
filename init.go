package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	pg "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

func initializeWindow() (*pixelgl.Window, pixelgl.WindowConfig) {
	windowWidthSize := (winHSize / 16) * 9

	cfg := pixelgl.WindowConfig{Title: "MyPlatformer",
		Bounds:    pixel.R(0, 0, float64(winHSize), float64(windowWidthSize)), //960,720
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	myUtil.CheckErrorPanic(err)
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

func initScreenInformation(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {

	switch myGame.CurrentGS {
	case myGame.GoToScreen:
		myGame.InitGoToScreen(win, Txt)
		if myGame.CurrentGS == myGame.GoToScreen && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.GoToClickEvent(win, win.MousePosition())
		}
	case myGame.StageSelect:
		myGame.InitStageSlect(win, Txt)
		if myGame.CurrentGS == myGame.StageSelect && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.StageClickEvent(win, win.MousePosition())
		}
	case myGame.TownScreen:
		myGame.InitTownScreen(win, Txt)
		if myGame.CurrentGS == myGame.TownScreen && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.TownClickEvent(win, win.MousePosition())
		}
	case myGame.WeaponShop:
		myGame.InitWeaponShop(win, Txt)
		if myGame.CurrentGS == myGame.WeaponShop && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyS, pg.KeyB, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.WeaponClickEvent(win, win.MousePosition(), player)
		}
	case myGame.ArmorShop:
		myGame.InitArmorShop(win, Txt)
		if myGame.CurrentGS == myGame.ArmorShop && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.ArmorClickEvent(win, win.MousePosition())
		}
	case myGame.AccessoryShop:
		myGame.InitAccessoryShop(win, Txt)
		if myGame.CurrentGS == myGame.AccessoryShop && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.AccessoryClickEvent(win, win.MousePosition())
		}
	case myGame.EquipmentScreen:
		myGame.InitEquipmentScreen(win, Txt)
		if myGame.CurrentGS == myGame.EquipmentScreen && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.EquipmentClickEvent(win, win.MousePosition())
		}
	case myGame.JobSelect:
		myGame.InitJobSelect(win, Txt)
		if myGame.CurrentGS == myGame.JobSelect && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.KeyBackspace) {
			myGame.CurrentGS = myGame.JobClickEvent(win, win.MousePosition(), player)
			myGame.SaveGame(myGame.SaveFilePath, 1, player)
		}
	case myGame.PlayingScreen:
		myGame.InitPlayingScreen(win, Txt)
	case myGame.BattleEnemyScreen:
		myGame.InitBattleEnemyScreen(win, Txt)
	}
	player.InitPlayerStatus(win, Txt)

}
