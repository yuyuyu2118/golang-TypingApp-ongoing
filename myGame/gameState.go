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
	BattleEnemyScreen
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

	topText := "Where are you going?"
	InitGoTo(win, Txt, topText)
}

func InitStageSlect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()

	topText := "Select play Stage"
	InitGoTo(win, Txt, topText)
	InitStage(win, Txt)
}

func InitTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitTown(win, Txt)
}

func InitWeaponShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitWeapon(win, Txt)
}

func InitArmorShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitArmor(win, Txt)
}

func InitAccessoryShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitAccessory(win, Txt)
}

func InitEquipmentScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitEquipment(win, Txt)
}

func InitJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)

	InitJob(win, Txt)
}

// func InitSaveScreen(win *pixelgl.Window, Txt *text.Text) {
// 	win.Clear(colornames.Black)

// 	InitSave(win, Txt)
// }

func InitPlayingScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitBattleEnemyScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitEndScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Grey)
	Txt.Clear()
}
