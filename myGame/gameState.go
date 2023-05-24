package myGame

import (
	"log"
	"os"

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

func InitStartScreen(win *pixelgl.Window, Txt *text.Text) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	startLines := []string{
		"This is a TypingBattleGame",
		"\n",
		"START : Press Enter",
	}

	myPos.LineCenterAlign(win, startLines, Txt, "center")

	//GoToScreenに行く
	if win.JustPressed(pixelgl.KeyEnter) {
		CurrentGS = GoToScreen
		log.Println("Press:Enter -> GameState:GoToScreen")
	}
	//testModeを開く
	if win.JustPressed(pixelgl.KeyT) {
		CurrentGS = TestState
		log.Println("TestMode")
	}
	//TODO: ゲーム終了、あとで削除?
	if win.JustPressed(pixelgl.KeyEscape) {
		win.Destroy()
		os.Exit(0)
	}
}

func InitGoToScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	//TODO: languageの設定
	topText := "どこに行く?"
	InitGoTo(win, Txt, topText)
}

func InitStageSlect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()

	topText := "だれと戦う?"
	InitGoTo(win, Txt, topText)
	InitStage(win, Txt)
}

func InitTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()

	topText := "どこに行く？"
	InitGoTo(win, Txt, topText)
	InitTown(win, Txt)
}

func InitWeaponShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()

	topText := "なにを買う?"
	InitWeapon(win, Txt, topText)
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
	Txt.Clear()

	topText := "装備画面"
	InitGoTo(win, Txt, topText)
	InitEquipment(win, Txt)
}

func InitJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()

	topText := "どれに転職する?"
	InitGoTo(win, Txt, topText)
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
