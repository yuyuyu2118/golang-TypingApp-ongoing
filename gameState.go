package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var currentGameState GameState

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

func initStartScreen(win *pixelgl.Window, Txt *text.Text) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	myPos.LineCenterAlign(win, startLines, Txt, "center")
}

func initGoToScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	initGoTo(win, Txt)
}

func initStageSlect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initStage(win, Txt)
}

func initTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initTown(win, Txt)
}

func initWeaponShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initWeapon(win, Txt)
}

func initArmorShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initArmor(win, Txt)
}

func initAccessoryShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initAccessory(win, Txt)
}

func initEquipmentScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initEquipment(win, Txt)
}

func initJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initJob(win, Txt)
}

func initSaveScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	initSave(win, Txt)
}

func initPlayingScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func initEndScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Grey)
	Txt.Clear()
}
