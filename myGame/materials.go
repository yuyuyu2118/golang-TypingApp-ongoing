package myGame

import (
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

	InitMaterialsBelong(win, Txt)
}

var getItemBool bool
var gotoSlice []string

func InitMaterialsBelong(win *pixelgl.Window, Txt *text.Text) {

	materialsBelongMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.5)
	materialsBelongMessageBox.DrawMessageBox()
	var materialsBelongOptions string

	materialsBelongOptions = "\n持ち物/素材 Tabで切り替え BackSpace.戻る\n\n"

	if !getItemBool {
		gotoSlice, _ = GetMyItems(SaveFilePathItems)
		getItemBool = true
	}

	gotoSlice = append(gotoSlice)
	var itemCount int

	//45まで繰り返し
	for _, gotoName := range gotoSlice {
		if itemCount >= 0 && itemCount <= 4 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 5 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 6 && itemCount <= 9 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 10 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 11 && itemCount <= 14 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 15 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 16 && itemCount <= 19 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 20 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 21 && itemCount <= 24 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 25 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 26 && itemCount <= 29 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 30 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 31 && itemCount <= 34 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 35 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 36 && itemCount <= 39 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 40 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		} else if itemCount >= 41 && itemCount <= 44 {
			materialsBelongOptions += gotoName + " "
			itemCount++
		} else if itemCount == 45 {
			materialsBelongOptions += gotoName + "\n"
			itemCount++
		}
	}
	materialsBelongMessageBox.DrawMessageTxt(materialsBelongOptions)
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
