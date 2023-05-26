package myGame

import (
	"log"
	"os"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"golang.org/x/image/colornames"
)

func InitStartScreen(win *pixelgl.Window, Txt *text.Text) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	startLines := []string{
		"タイピングバトルRPG",
		"\n",
		"Enterキーを押してスタート",
		"Escapeキーを押して終了",
		"(※オートセーブです)",
	}

	myPos.LineCenterAlign(win, startLines, Txt, "center")

	//GoToScreenに行く
	if win.JustPressed(pixelgl.KeyEnter) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("Press:Enter -> GameState:GoToScreen")
	}
	//testModeを開く
	if win.JustPressed(pixelgl.KeyT) {
		myState.CurrentGS = myState.TestState
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
	botText := "どこに行く? BackSpace.戻る"
	InitGoTo(win, Txt, botText)
}

func InitStageSlect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "だれと戦う? BackSpace.戻る"
	InitGoTo(win, Txt, botText)
	InitStage(win, Txt)
}

func InitTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "どこに行く? BackSpace.戻る"
	InitGoTo(win, Txt, botText)
	InitTown(win, Txt)
}

func InitWeaponShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "なにを買う? BackSpace.戻る"
	InitWeapon(win, Txt, botText)
}

func InitArmorShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)

	InitArmor(win, Txt)
}

func InitAccessoryShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)

	InitAccessory(win, Txt)
}

func InitEquipmentScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	InitEquipment(win, Txt)
}

func InitJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "どれに転職する? BackSpace.戻る"
	InitGoTo(win, Txt, botText)
	InitJob(win, Txt)
}

func InitPlayingScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitBattleEnemyScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitSkillScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Black)
	Txt.Clear()
}

func InitEndScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Grey)
	Txt.Clear()
}
