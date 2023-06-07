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

	startLines := []string{
		"タイピングバトルRPG",
		"\n",
		"Enterキーを押してスタート",
	}

	// startPos := pixel.V(win.Bounds().Center().X, win.Bounds().Max.Y-StartTxt.LineHeight*float64(len(startLines)))

	// // 各行の終了位置を計算
	// lineEndPositions := make([]pixel.Vec, len(startLines))
	// for idx, line := range startLines {
	// 	centerX := win.Bounds().Center().Sub(StartTxt.BoundsOf(line).Center()).X
	// 	lineEndPositions[idx] = pixel.V(centerX, win.Bounds().Center().Y-StartTxt.LineHeight*float64(len(startLines)/2-idx))
	// }

	// for idx, line := range startLines {
	// 	endPos := lineEndPositions[idx]

	// 	if animationStartTime.IsZero() {
	// 		animationStartTime = time.Now()
	// 	}

	// 	//animationStartTime.Add(time.Duration(animationDuration*float64(idx))*time.Second)
	// 	myUtil.AnimateText(win, StartTxt, myUtil.CompletedText, []string{line}, animationStartTime, startPos, endPos, animationDuration)
	// }

	myPos.LineCenterAlign(win, startLines, StartTxt, "center")

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

	botText := "どれに転職する? BackSpace.戻る"
	InitGoTo(win, Txt, botText)
	InitJob(win, Txt)
}

<<<<<<< HEAD
func InitBlackSmithScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
=======
func InitBlackSmithScreen(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {
>>>>>>> d236a68 (途中)
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
