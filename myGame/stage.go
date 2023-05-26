package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"golang.org/x/image/colornames"
)

var StageNum int

var (
	stage1Button = pixel.Rect{}
)

var (
	stageButtonSlice = []pixel.Rect{}
)

func InitStage(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//stageSlice := []string{"1. Slime", "2. Bird", "3. Plant", "4. Goblin", "5. Zombie", "6. Fairy", "7. Skull", "8. Wizard", "9. Solidier", "10. Dragon", "BackSpace. EXIT"}
	stageSlice := []string{"1. スライム", "2. バード", "3. プラント", "4. ゴブリン", "5. ゾンビ", "6. フェアリー", "7. スカル", "8. ウィザード", "9. ソルジャー", "BackSpace. 戻る"}

	for _, stageName := range stageSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, stageName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		stageButtonSlice = append(stageButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func StageClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {

	if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Slime")
		StageNum = 0
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Bird")
		StageNum = 1
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Plant")
		StageNum = 2
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Goblin")
		StageNum = 3
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Zombie")
		StageNum = 4
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.Key6)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Fairy")
		StageNum = 5
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[6].Contains(mousePos) || win.JustPressed(pixelgl.Key7)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Skull")
		StageNum = 6
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[7].Contains(mousePos) || win.JustPressed(pixelgl.Key8)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Wizard")
		StageNum = 7
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[8].Contains(mousePos) || win.JustPressed(pixelgl.Key9)) {
		myState.CurrentGS = myState.PlayingScreen
		log.Println("PlayStage is VS Solidier")
		StageNum = 8
	} else if myState.CurrentGS == myState.StageSelect && (stageButtonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("StageScreen -> GoToScreen")
	}
	log.Println("PlayStage is", StageNum)
	return myState.CurrentGS
}
