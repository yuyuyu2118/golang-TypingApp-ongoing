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

func InitMaterialsBelongScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	topText := "持ち物/素材"
	InitMaterialsBelong(win, Txt, topText)
}

func InitMaterialsBelong(win *pixelgl.Window, Txt *text.Text, topText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, topText, "1.武器", "2.防具", "3.アクセサリー", "4.素材", "BackSpace.戻る")
	tempPosition = myPos.BotCenPos(win, Txt)
	myPos.DrawPos(win, Txt, tempPosition)

	gotoSlice, _ := GetMyItems(SaveFilePathItems)
	gotoSlice = append(gotoSlice)

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

func MaterialsBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	//TODO: 全部この形式にする　やばいバグ
	if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.TownScreen
		log.Println("GoToScreen->Town")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.EquipmentScreen
		log.Println("GoToScreen->Equipment")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.StartScreen
		log.Println("GoToScreen->StartScreen")
	}
	return myState.CurrentGS
}
