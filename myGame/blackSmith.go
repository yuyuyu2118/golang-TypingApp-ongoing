package myGame

import (
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	pg "github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
<<<<<<< HEAD
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
=======
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
>>>>>>> d236a68 (途中)
)

// var (
// 	japanFontPath = "assets/fonts/PixelMplus12-Regular.ttf"
// )

var (
	blackSmithButtonSlice = []pixel.Rect{}
	// weaponBelongButtonSlice = []pixel.Rect{}
)

var blackSmithSlice = []string{}

var tabCountBlackSmith int

<<<<<<< HEAD
func InitBlackSmith(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
=======
func InitBlackSmith(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus) {
>>>>>>> d236a68 (途中)
	if win.JustPressed(pixelgl.KeyTab) {
		if tabCountBlackSmith == 0 {
			myState.CurrentBlackSmith = myState.ArmorBlackSmith
			tabCountBlackSmith++
		} else if tabCountBlackSmith == 1 {
			myState.CurrentBlackSmith = myState.AccessoryBlackSmith
			tabCountBlackSmith++
		} else if tabCountBlackSmith == 2 {
			myState.CurrentBlackSmith = myState.MaterialsBlackSmith
			tabCountBlackSmith++
		} else if tabCountBlackSmith == 3 {
			myState.CurrentBlackSmith = myState.WeaponBlackSmith
			tabCountBlackSmith = 0
		}
	}
	switch myState.CurrentBlackSmith {
	case myState.WeaponBlackSmith:
		InitWeaponBlackSmithScreen(win, myUtil.DescriptionTxt, player)
		if myState.CurrentBlackSmith == myState.WeaponBlackSmith && myUtil.AnyKeyJustPressed(win, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
			WeaponBlackSmithClickEvent(win, win.MousePosition(), player)
		}
		// case myState.ArmorBlackSmith:
		// 	InitArmorBlackSmithScreen(win, myUtil.DescriptionTxt, player)
		// 	if myState.CurrentBlackSmith == myState.ArmorBlackSmith && myUtil.AnyKeyJustPressed(win, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
		// 		ArmorBlackSmithClickEvent(win, win.MousePosition(), player)
		// 	}
		// case myState.AccessoryBlackSmith:
		// 	InitAccessoryBlackSmithScreen(win, myUtil.DescriptionTxt, player)
		// 	if myState.CurrentBlackSmith == myState.AccessoryBlackSmith && myUtil.AnyKeyJustPressed(win, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
		// 		AccessoryBlackSmithClickEvent(win, win.MousePosition(), player)
		// 	}
		// case myState.MaterialsBlackSmith:
		// 	InitMaterialsBlackSmithScreen(win, myUtil.DescriptionTxt)
		// 	if myState.CurrentBlackSmith == myState.MaterialsBlackSmith && myUtil.AnyKeyJustPressed(win, pg.Key1, pg.Key2, pg.Key3, pg.Key4, pg.Key5, pg.Key6, pg.Key7, pg.Key8, pg.Key9, pg.Key0, pg.KeyBackspace) {
		// 		myState.CurrentBlackSmith = BlackSmithClickEvent(win, win.MousePosition())
		// 	}
	}
}

<<<<<<< HEAD
// func BlackSmithClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {
=======
// func BlackSmithClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
>>>>>>> d236a68 (途中)
// 	if len(equipmentButtonSlice) > 0 {
// 		if myState.CurrentGS == myState.EquipmentScreen && (win.JustPressed(pixelgl.KeyBackspace)) {
// 			myState.CurrentGS = myState.GoToScreen
// 			getItemBool = false
// 			log.Println("equipment->GoToScreen")
// 		}
// 	}
// 	return myState.CurrentGS
// }

type BlackSmithState int

const (
	blackSmithNil BlackSmithState = iota
	blackSmith1
	blackSmith2
	blackSmith3
	blackSmith4
	blackSmith5
	blackSmith6
	blackSmith7
	blackSmith8
	blackSmith9
	blackSmith10
)

var keyToBlackSmith = map[pixelgl.Button]BlackSmithState{
	pixelgl.Key1: blackSmith1,
	pixelgl.Key2: blackSmith2,
	pixelgl.Key3: blackSmith3,
	pixelgl.Key4: blackSmith4,
	pixelgl.Key5: blackSmith5,
	pixelgl.Key6: blackSmith6,
	pixelgl.Key7: blackSmith7,
	pixelgl.Key8: blackSmith8,
	pixelgl.Key9: blackSmith9,
	pixelgl.Key0: blackSmith10,
}

var currentBlackSmithState BlackSmithState

<<<<<<< HEAD
func BlackSmithClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {
=======
func BlackSmithClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
>>>>>>> d236a68 (途中)
	var tempBlackSmith = ""

	for i := 0; i < len(keyToBlackSmith)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.WeaponPurchaseEventInstance.Weapons[i] && myState.CurrentGS == myState.BlackSmithScreen {
			currentBlackSmithState = BlackSmithState(i + 1)
			//CreateWeaponEvent(descWeapon, 0)
			log.Println("鍛冶屋->鍛冶", i+1)
			tempMyMaterialBool = false
			tempMyMaterialName = []string{"", "", "", "", "", ""}
			tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
			break
		}
	}

	if (win.JustPressed(pixelgl.Key0)) && event.WeaponPurchaseEventInstance.Weapons[9] && myState.CurrentGS == myState.BlackSmithScreen {
		currentBlackSmithState = blackSmith10
		log.Println("鍛冶屋->鍛冶10")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", "", "", "", "", ""}
		tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.BlackSmithScreen {
		myState.CurrentGS = myState.GoToScreen
		log.Println("鍛冶屋->Goto")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", "", "", "", "", ""}
		tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
	}

	if (win.JustPressed(pixelgl.KeyS)) && player.Gold >= 0 {
		loadContent := SaveFileLoad(SaveFilePath)
		//TODO: お金が足りないときの処理を記述
		for i := 0; i < len(keyToBlackSmith)-1; i++ {
			if currentBlackSmithState == BlackSmithState(i+1) {
				//強化に必要なゴールド
				requiredGold, _ := strconv.Atoi(descWeapon[i+1][18])
				log.Println(requiredGold)
				belongWeapon, _ := strconv.Atoi(loadContent[3][i])
				if belongWeapon >= 1 {
					if player.Gold >= requiredGold {
						log.Println(descWeapon[i+1][4], "買える", player.Gold)
						enhancementOk := EnhancementWeaponEvent(win, descWeapon, i)
						if enhancementOk {
							player.Gold -= requiredGold
							tempBlackSmith = "weapon" + strconv.Itoa(i+1)
							tempMyMaterialBool = false
							tempMyMaterialName = []string{"", "", "", "", "", ""}
							tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
						}
					} else {
						log.Println(descWeapon[i+1][4], "お金が足りない", player.Gold)
					}
				} else {
					log.Println("すでに持っている")
					break
				}
			}
		}
		if currentBlackSmithState == blackSmith10 {
			requiredGold, _ := strconv.Atoi(descWeapon[10][18])
			log.Println(requiredGold)
			belongWeapon, _ := strconv.Atoi(loadContent[3][9])
			if belongWeapon >= 1 {
				if player.Gold >= requiredGold {
					log.Println(descWeapon[10][18], "買える", player.Gold)
					enhancementOk := EnhancementWeaponEvent(win, descWeapon, 9)
					if enhancementOk {
						player.Gold -= requiredGold
						tempBlackSmith = "weapon" + strconv.Itoa(0)
						tempMyMaterialBool = false
						tempMyMaterialName = []string{"", "", "", "", "", ""}
						tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
					}
				} else {
					log.Println(descWeapon[10][18], "お金が足りない", player.Gold)
				}
			} else {
				log.Println("すでに持っている")
			}
		}

		if tempBlackSmith != "" {
			SaveWeaponEnhancementEvent(SaveFilePath, 9, tempBlackSmith, player)
			SaveGame(SaveFilePath, 1, player)
		}
<<<<<<< HEAD

		//TODO: 関数化
		loadContent = SaveFileLoad(SaveFilePath)

		tempOP1, _ := strconv.ParseFloat(loadContent[1][13], 64)
		tempOP2, _ := strconv.ParseFloat(player.EquipmentWeapon[1], 64)
		tempOP3, _ := strconv.ParseFloat(player.EquipmentAccessory[1], 64)

		var tempOP4 float64
		tempName := player.EquipmentWeapon[0]

		for i, name := range weaponName {
			if tempName == name {
				coefficient, _ := strconv.ParseFloat(loadContent[9][i], 64)
				tempWeaponEnhancement, _ := strconv.ParseFloat(descWeapon[i+1][25], 64)
				tempOP4 = tempWeaponEnhancement * coefficient
				//log.Println("player", coefficient, tempWeaponEnhancement)
			}
		}

		log.Println(tempOP1, tempOP2, tempOP3, tempOP4)

		player.OP = tempOP1 + tempOP2 + tempOP3 + tempOP4

		SaveGame(SaveFilePath, 1, player)
		SaveGameWeapon(SaveFilePath, 6, player)
=======
>>>>>>> d236a68 (途中)
	}
	return myState.CurrentGS
}
