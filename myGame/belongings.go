package myGame

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	pg "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

var (
	japanFontPath = "assets/fonts/PixelMplus12-Regular.ttf"
)

var (
	equipmentButtonSlice    = []pixel.Rect{}
	weaponBelongButtonSlice = []pixel.Rect{}
)

var equipmentSlice = []string{}

var tabCount int

func InitEquipment(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {
	if win.JustPressed(pixelgl.KeyTab) {
		if tabCount == 0 {
			myState.CurrentBelong = myState.ArmorBelong
			tabCount++
		} else if tabCount == 1 {
			myState.CurrentBelong = myState.AccessoryBelong
			tabCount++
		} else if tabCount == 2 {
			myState.CurrentBelong = myState.MaterialsBelong
			tabCount++
		} else if tabCount == 3 {
			myState.CurrentBelong = myState.WeaponBelong
			tabCount = 0
		}
	}
	switch myState.CurrentBelong {
	case myState.WeaponBelong:
		InitWeaponBelongScreen(win, myUtil.DescriptionTxt, player)
		if myState.CurrentBelong == myState.WeaponBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
			WeaponBelongClickEvent(win, win.MousePosition(), player)
		}
	case myState.ArmorBelong:
		InitArmorBelongScreen(win, myUtil.DescriptionTxt, player)
		if myState.CurrentBelong == myState.ArmorBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
			ArmorBelongClickEvent(win, win.MousePosition(), player)
		}
	case myState.AccessoryBelong:
		InitAccessoryBelongScreen(win, myUtil.DescriptionTxt, player)
		if myState.CurrentBelong == myState.AccessoryBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
			AccessoryBelongClickEvent(win, win.MousePosition(), player)
		}
	case myState.MaterialsBelong:
		InitMaterialsBelongScreen(win, myUtil.DescriptionTxt)
		if myState.CurrentBelong == myState.MaterialsBelong && myUtil.AnyKeyJustPressed(win, pg.MouseButtonLeft, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
			myState.CurrentBelong = BelongClickEvent(win, win.MousePosition())
		}
	}
}

func EquipmentClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	if myState.CurrentGS == myState.EquipmentScreen && (equipmentButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		getItemBool = false
		log.Println("equipment->GoToScreen")
	}
	return myState.CurrentGS
}

func BelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.BelongState {
	// if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key1)) {
	// 	myState.CurrentBelong = myState.WeaponBelong
	// 	log.Println("WeaponBelong->WeaponBelong")
	// } else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key2)) {
	// 	myState.CurrentBelong = myState.ArmorBelong
	// 	log.Println("WeaponBelong->ArmorBelong")
	// } else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key3)) {
	// 	myState.CurrentBelong = myState.AccessoryBelong
	// 	log.Println("WeaponBelong->AccessoryBelong")
	// } else if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.Key4)) {
	// 	myState.CurrentBelong = myState.MaterialsBelong
	// 	log.Println("WeaponBelong->MaterialsBelong")
	// }
	if (myState.CurrentBelong == myState.WeaponBelong || myState.CurrentBelong == myState.ArmorBelong || myState.CurrentBelong == myState.AccessoryBelong || myState.CurrentBelong == myState.MaterialsBelong) && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentGS = myState.GoToScreen
		myState.CurrentBelong = myState.WeaponBelong
		getItemBool = false
		tabCount = 0
		log.Println("shop->Back")
	}
	return myState.CurrentBelong
}
