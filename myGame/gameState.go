package myGame

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var CurrentGS GameState

type GameState int

const (
	StartScreen GameState = iota
	GoToScreen

	StageSelect
	TownScreen
	EquipmentScreen
	JobSelect
	SaveScreen

	PlayingScreen
	EndScreen
	TestState

	WeaponShop
	ArmorShop
	AccessoryShop
	BlackSmith
)

var (
	startLines = []string{
		"This is a TypingBattleGame",
		"\n",
		"START : Press Enter",
	}
)

func InitStartScreen(win *pixelgl.Window, Txt *text.Text) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	myPos.LineCenterAlign(win, startLines, Txt, "center")
}

func InitGoToScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	initGoTo(win, Txt)
}

func InitStageSlect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initStage(win, Txt)
}

func InitTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initTown(win, Txt)
}

func InitWeaponShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initWeapon(win, Txt)
}

func InitArmorShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initArmor(win, Txt)
}

func InitAccessoryShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initAccessory(win, Txt)
}

func InitEquipmentScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initEquipment(win, Txt)
}

func InitJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initJob(win, Txt)
}

func InitSaveScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initSave(win, Txt)
}

func InitPlayingScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitEndScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Grey)
	Txt.Clear()
}
