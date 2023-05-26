package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	pg "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"golang.org/x/image/colornames"
)

var CurrentBelong BelongState

type BelongState int

const (
	WeaponBelong BelongState = iota
	ArmorBelong
	AccessoryBelong
	MaterialsBelong
)

var (
	japanFontPath = "assets/fonts/PixelMplus12-Regular.ttf"
)

var (
	equipmentButtonSlice    = []pixel.Rect{}
	weaponBelongButtonSlice = []pixel.Rect{}
)

var equipmentSlice = []string{}

func InitEquipment(win *pixelgl.Window, Txt *text.Text) {
	xOffSet := myPos.TopLefPos(win, Txt).X + 400
	yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	txtPos := pixel.V(0, 0)

	switch CurrentBelong {
	case WeaponBelong:
		if CurrentBelong == WeaponBelong && belongAnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case ArmorBelong:
		if CurrentBelong == ArmorBelong && belongAnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case AccessoryBelong:
		if CurrentBelong == AccessoryBelong && belongAnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case MaterialsBelong:
		if CurrentBelong == MaterialsBelong && belongAnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	}

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

func EquipmentClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	if myState.CurrentGS == myState.EquipmentScreen && (equipmentButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("equipment->GoToScreen")
	}
	return myState.CurrentGS
}

func BelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) BelongState {
	//TODO: 循環参照を避けて、マウスpositionを追加
	if (CurrentBelong == WeaponBelong || CurrentBelong == ArmorBelong || CurrentBelong == AccessoryBelong || CurrentBelong == MaterialsBelong) && (win.JustPressed(pixelgl.Key1)) {
		CurrentBelong = WeaponBelong
		log.Println("WeaponBelong->WeaponBelong")
	} else if (CurrentBelong == WeaponBelong || CurrentBelong == ArmorBelong || CurrentBelong == AccessoryBelong || CurrentBelong == MaterialsBelong) && (win.JustPressed(pixelgl.Key2)) {
		CurrentBelong = ArmorBelong
		log.Println("WeaponBelong->ArmorBelong")
	} else if (CurrentBelong == WeaponBelong || CurrentBelong == ArmorBelong || CurrentBelong == AccessoryBelong || CurrentBelong == MaterialsBelong) && (win.JustPressed(pixelgl.Key3)) {
		CurrentBelong = AccessoryBelong
		log.Println("WeaponBelong->AccessoryBelong")
	} else if (CurrentBelong == WeaponBelong || CurrentBelong == ArmorBelong || CurrentBelong == AccessoryBelong || CurrentBelong == MaterialsBelong) && (win.JustPressed(pixelgl.Key4)) {
		CurrentBelong = MaterialsBelong
		log.Println("WeaponBelong->MaterialsBelong")
	}
	return CurrentBelong
}

func belongAnyKeyJustPressed(win *pixelgl.Window, keys ...pixelgl.Button) bool {
	for _, key := range keys {
		if win.JustPressed(key) {
			return true
		}
	}
	return false
}
