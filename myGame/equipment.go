package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"golang.org/x/image/colornames"
)

var (
	equipmentButtonSlice = []pixel.Rect{}
)

var equipmentSlice = []string{}

func InitEquipment(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	//gotoSlice := []string{"1. Dungeon", "2. Town", "3. Equipment", "4. Job", "5. Save", "6. EXIT"}
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
	equipmentSlice = append(equipmentSlice, "BackSpace. 戻る")

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

func EquipmentClickEvent(win *pixelgl.Window, mousePos pixel.Vec) GameState {
	//TODO ページを作成したら追加
	if CurrentGS == EquipmentScreen && (equipmentButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		CurrentGS = GoToScreen
		log.Println("equipment->GoToScreen")
	}
	return CurrentGS
}

// func initPlayerEquipment(win *pixelgl.Window, Txt *text.Text, player *PlayerStatus) {
// 	Txt.Clear()
// 	Txt.Color = colornames.White
// 	fmt.Fprintln(Txt, "Weapon: ", "\nArmor: ", "\nAccessory: ")
// 	xOffSet := 0.0
// 	yOffSet := win.Bounds().H()/3 - Txt.LineHeight*3
// 	txtPos := pixel.V(xOffSet, yOffSet)
// 	tempPosition := pixel.IM.Moved(txtPos)
// 	Txt.Draw(win, tempPosition)
// }
