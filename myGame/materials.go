package myGame

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

func InitMaterialsBelongScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/素材"
	InitMaterialsBelong(win, Txt, botText)
}

var getItemBool bool
var gotoSlice []string

func InitMaterialsBelong(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet := 100.0
	yOffSet1Line := myPos.TopLefPos(win, Txt).Y - 100
	yOffSet2Line := myPos.TopCenPos(win, Txt).Y - 100
	yOffSet3Line := myPos.TopCenPos(win, Txt).Y - 100

	txtPos := pixel.V(0, 0)

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, botText, "Tabで切り替え", "BackSpace.戻る")
	tempPosition = myPos.BotCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)

	if !getItemBool {
		gotoSlice, _ = GetMyItems(SaveFilePathItems)
		getItemBool = true
	}

	gotoSlice = append(gotoSlice)
	var itemCount int

	//TODO: 折り返し
	for _, gotoName := range gotoSlice {
		if itemCount >= 0 && itemCount <= 14 {
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, gotoName)
			yOffSet1Line -= Txt.LineHeight + 10
			txtPos = pixel.V(xOffSet, yOffSet1Line)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
			gotoButtonSlice = append(gotoButtonSlice, Txt.Bounds().Moved(txtPos))
			itemCount++
		} else if itemCount >= 15 && itemCount <= 29 {
			xOffSet = 400
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, gotoName)
			yOffSet2Line -= Txt.LineHeight + 10
			txtPos = pixel.V(xOffSet, yOffSet2Line)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
			gotoButtonSlice = append(gotoButtonSlice, Txt.Bounds().Moved(txtPos))
			itemCount++
		} else if itemCount > 30 && itemCount <= 45 {
			xOffSet += 400
			Txt.Clear()
			Txt.Color = colornames.White
			fmt.Fprintln(Txt, gotoName)
			yOffSet3Line -= Txt.LineHeight + 10
			txtPos = pixel.V(xOffSet, yOffSet3Line)
			tempPosition := pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
			gotoButtonSlice = append(gotoButtonSlice, Txt.Bounds().Moved(txtPos))
			itemCount++
		}
	}
}

func MaterialsBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("GoToScreen->Dungeon")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key2)) {
		myState.CurrentGS = myState.TownScreen
		log.Println("GoToScreen->Town")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key3)) {
		myState.CurrentGS = myState.EquipmentScreen
		log.Println("GoToScreen->Equipment")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.Key4)) {
		myState.CurrentGS = myState.JobSelect
		log.Println("GoToScreen->JobSelect")
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.WeaponBelong
		myState.CurrentGS = myState.StartScreen
		log.Println("所持品/素材->GoTo")
	}
	return myState.CurrentGS
}
