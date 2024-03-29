package myGame

import (
	"image/color"
	"log"
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

var (
	animationStartTime time.Time
)

func InitStartScreen(win *pixelgl.Window, StartTxt *text.Text, alpha float64, animationDuration float64) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	StartTxt.Clear()

	txtColor := color.RGBA{R: 255, G: 255, B: 255, A: uint8(alpha * 255)} // テキストのアルファ値を更新
	StartTxt.Color = txtColor

	myUtil.CompletedTxt.Color = txtColor

	// startLines := []string{
	// 	"タイピングバトルRPG",
	// 	"\n",
	// 	"Enterキーを押してスタート",
	// }
	// DrawCenteredText関数を使用してテキストを中央に描画
	myPos.RelativeDraw(win, StartTxt, "タイピングバトルRPG", 0.5, 0.55)
	myPos.RelativeDraw(win, StartTxt, "Enterキーを押してスタート", 0.5, 0.45)

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

	InitStage(win, Txt)
}

func InitTownScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

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
	Txt.Clear()

	botText := "なにを買う? BackSpace.戻る"
	InitArmor(win, Txt, botText)
}

func InitAccessoryShop(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "なにを買う? BackSpace.戻る"
	InitAccessory(win, Txt, botText)
}

func InitEquipmentScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	InitEquipment(win, Txt, player)
}

func InitJobSelect(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	InitJob(win, Txt)
}

func InitBlackSmithScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	InitBlackSmith(win, Txt, player)
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
