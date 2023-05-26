package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

func InitAccessoryBelongScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/アクセサリー"
	InitAccessoryBelong(win, Txt, botText)
}

func InitAccessoryBelong(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, botText, "1.武器", "2.防具", "3.アクセサリー", "4.素材", "BackSpace.戻る")
	tempPosition = myPos.BotCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)

	loadContent := SaveFileLoad(SaveFilePath)
	counts := make(map[string]int)
	elements := loadContent[3]

	for i, val := range elements {
		num, err := strconv.Atoi(val)
		if err == nil {
			weaponKey := fmt.Sprintf("weapon%d", i)
			counts[weaponKey] = num
		}
	}

	for i, value := range weaponName {
		if counts["weapon"+strconv.Itoa(i)] != 0 {
			tempInt := counts["weapon"+strconv.Itoa(i)]
			equipmentSlice = append(equipmentSlice, value+": "+strconv.Itoa(tempInt))
		}
	}

	for _, equipmentName := range equipmentSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, equipmentName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		equipmentButtonSlice = append(equipmentButtonSlice, Txt.Bounds().Moved(txtPos))
	}
	equipmentSlice = equipmentSlice[:0]
}

func AccessoryBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
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
	} else if myState.CurrentGS == myState.GoToScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.WeaponBelong
		myState.CurrentGS = myState.StartScreen
		log.Println("所持品/アクセサリー->GoTo")
	}
	return myState.CurrentGS
}
