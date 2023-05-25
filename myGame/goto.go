package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
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

var (
	gotoButtonSlice = []pixel.Rect{}
)

func InitGoTo(win *pixelgl.Window, Txt *text.Text, topText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, topText)
	tempPosition = myPos.TopCenPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
	gotoSlice := []string{"1. ダンジョン", "2. 町", "3. 装備", "4. ジョブ", "BackSpace. 戻る"}

	for _, gotoName := range gotoSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, gotoName)
		yOffSet -= Txt.LineHeight + 40
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		gotoButtonSlice = append(gotoButtonSlice, Txt.Bounds().Moved(txtPos))
	}
}

func GoToClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	//TODO: 全部この形式にする　やばいバグ
	if CurrentGS == GoToScreen && (gotoButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		CurrentGS = StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if CurrentGS == GoToScreen && (gotoButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		CurrentGS = TownScreen
		log.Println("GoToScreen->Town")
	} else if CurrentGS == GoToScreen && (gotoButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		CurrentGS = EquipmentScreen
		log.Println("GoToScreen->Equipment")
	} else if CurrentGS == GoToScreen && (gotoButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		CurrentGS = JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if CurrentGS == GoToScreen && (gotoButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		CurrentGS = StartScreen
		log.Println("GoToScreen->StartScreen")
	}
	return CurrentGS
}
