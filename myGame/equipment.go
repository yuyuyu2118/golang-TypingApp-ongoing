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
	equip1Button = pixel.Rect{}
	// equip2Button = pixel.Rect{}
	// equip3Button = pixel.Rect{}
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
	loadContent := SaveFileLoad("player\\save\\save.csv")
	counts := make(map[string]int)
	elements := loadContent[3]

	for i, val := range elements {
		num, err := strconv.Atoi(val)
		if err == nil {
			weaponKey := fmt.Sprintf("weapon%d", i)
			counts[weaponKey] = num
		}
	}

	log.Println(counts["weapon0"])

	if counts["weapon0"] != 0 {
		tempInt := counts["weapon0"]
		equipmentSlice = append(equipmentSlice, "木の棒: "+strconv.Itoa(tempInt))
	}
	if counts["weapon1"] != 0 {
		tempInt := counts["weapon1"]
		equipmentSlice = append(equipmentSlice, "果物ナイフ: "+strconv.Itoa(tempInt))
	}
	if counts["weapon2"] != 0 {
		tempInt := counts["weapon2"]
		equipmentSlice = append(equipmentSlice, "木刀: "+strconv.Itoa(tempInt))
	}
	if counts["weapon3"] != 0 {
		tempInt := counts["weapon3"]
		equipmentSlice = append(equipmentSlice, "ドレインソード: "+strconv.Itoa(tempInt))
	}
	if counts["weapon4"] != 0 {
		tempInt := counts["weapon4"]
		equipmentSlice = append(equipmentSlice, "スタンハンマー: "+strconv.Itoa(tempInt))
	}
	if counts["weapon5"] != 0 {
		tempInt := counts["weapon5"]
		equipmentSlice = append(equipmentSlice, "鉄の剣: "+strconv.Itoa(tempInt))
	}
	if counts["weapon6"] != 0 {
		tempInt := counts["weapon6"]
		equipmentSlice = append(equipmentSlice, "隼の件: "+strconv.Itoa(tempInt))
	}
	if counts["weapon7"] != 0 {
		tempInt := counts["weapon7"]
		equipmentSlice = append(equipmentSlice, "勇者の剣: "+strconv.Itoa(tempInt))
	}
	if counts["weapon8"] != 0 {
		tempInt := counts["weapon8"]
		equipmentSlice = append(equipmentSlice, "名刀村正: "+strconv.Itoa(tempInt))
	}
	if counts["weapon9"] != 0 {
		tempInt := counts["weapon9"]
		equipmentSlice = append(equipmentSlice, "死神の大鎌: "+strconv.Itoa(tempInt))
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
	if equipmentButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace) {
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
