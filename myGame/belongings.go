package myGame

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	pg "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
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
	// xOffSet := myPos.TopLefPos(win, Txt).X + 400
	// yOffSet := myPos.TopLefPos(win, Txt).Y - 50
	// txtPos := pixel.V(0, 0)

	switch myState.CurrentBelong {
	case myState.WeaponBelong:
		InitWeaponBelongScreen(win, myUtil.ScreenTxt)
		if myState.CurrentBelong == myState.WeaponBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			myState.CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case myState.ArmorBelong:
		if myState.CurrentBelong == myState.ArmorBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			myState.CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case myState.AccessoryBelong:
		if myState.CurrentBelong == myState.AccessoryBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			myState.CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	case myState.MaterialsBelong:
		InitMaterialsBelongScreen(win, myUtil.ScreenTxt)
		if myState.CurrentBelong == myState.MaterialsBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.KeyBackspace) {
			myState.CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	}
}

func EquipmentClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	//TODO ページを作成したら追加
	if myState.CurrentGS == myState.EquipmentScreen && (equipmentButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		log.Println("equipment->GoToScreen")
	}
	return myState.CurrentGS
}

func BelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.BelongState {
	//TODO: 循環参照を避けて、マウスpositionを追加
	if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key1)) {
		myState.CurrentBelong = myState.WeaponBelong
		log.Println("WeaponBelong->WeaponBelong")
	} else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key2)) {
		myState.CurrentBelong = myState.ArmorBelong
		log.Println("WeaponBelong->ArmorBelong")
	} else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key3)) {
		myState.CurrentBelong = myState.AccessoryBelong
		log.Println("WeaponBelong->AccessoryBelong")
	} else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key4)) {
		myState.CurrentBelong = myState.MaterialsBelong
		log.Println("WeaponBelong->MaterialsBelong")
	}
	return myState.CurrentBelong
}
