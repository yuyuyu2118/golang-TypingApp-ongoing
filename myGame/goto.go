package myGame

import (
	_ "image/png"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

var tempPosition pixel.Vec

var (
	goTo1Button = pixel.Rect{}
	goTo2Button = pixel.Rect{}
	goTo3Button = pixel.Rect{}
	goTo4Button = pixel.Rect{}
	goTo5Button = pixel.Rect{}
	goTo6Button = pixel.Rect{}
)

// GoToScreenの初期化
func InitGoTo(win *pixelgl.Window, Txt *text.Text, bottleText string) {
	gotoMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.4)
	gotoMessageBox.DrawMessageBox()
	gotoMessageBox.DrawMessageTxt("どこへ行きますか？キーボードに対応する数字を入力してください。\n1. ダンジョン\n2. 町\n3. ジョブ\n4. 鍛冶屋\n\nBackSpaceキーでタイトルに戻る")
}

func GoToClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.TownScreen
		log.Println("GoToScreen->Town")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.BlackSmithScreen
		log.Println("GoToScreen->BlackSmithScreen")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.StartScreen
		log.Println("GoToScreen->StartScreen")
	}
	return myState.CurrentGS
}
