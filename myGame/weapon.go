package myGame

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

type WeaponState int

const (
	weaponNil WeaponState = iota
	weapon1
	weapon2
	weapon3
	weapon4
	weapon5
	weapon6
	weapon7
	weapon8
	weapon9
	weapon0
)

var weaponSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var weaponNum = []string{"weapon0", "weapon1", "weapon2", "weapon3", "weapon4", "weapon5", "weapon6", "weapon7", "weapon8", "weapon9"}
var weaponName = []string{"木の棒", "果物ナイフ", "木刀", "ドレインソード", "スタンハンマー", "鉄の剣", "隼の剣", "勇者の剣", "名刀村正", "死神の大鎌"}

var (
	weaponPath = "assets/shop/weapon.csv"
	descWeapon = CsvToSlice(weaponPath)
)
var currentweaponState WeaponState

func InitWeapon(win *pixelgl.Window, Txt *text.Text, topText string) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)
	Txt.Color = colornames.White

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, topText)
	tempPosition = myPos.BotCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)

	if event.WeaponPurchaseEventInstance.Weapon1 {
		weaponSlice[0] = "1. 木の棒"
	}
	if event.WeaponPurchaseEventInstance.Weapon2 {
		weaponSlice[1] = "2. 果物ナイフ"
	}
	if event.WeaponPurchaseEventInstance.Weapon3 {
		weaponSlice[2] = "3. 木刀"
	}
	if event.WeaponPurchaseEventInstance.Weapon4 {
		weaponSlice[3] = "4. ドレインソード"
	}
	if event.WeaponPurchaseEventInstance.Weapon5 {
		weaponSlice[4] = "5. スタンハンマー"
	}
	if event.WeaponPurchaseEventInstance.Weapon6 {
		weaponSlice[5] = "6. 鉄の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon7 {
		weaponSlice[6] = "7. 隼の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon8 {
		weaponSlice[7] = "8. 勇者の剣"
	}
	if event.WeaponPurchaseEventInstance.Weapon9 {
		weaponSlice[8] = "9. 名刀村正"
	}
	if event.WeaponPurchaseEventInstance.Weapon0 {
		weaponSlice[9] = "0. 死神の大鎌"
	}

	for _, weaponName := range weaponSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, weaponName)
		yOffSet -= Txt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		buttonSlice = append(buttonSlice, Txt.Bounds().Moved(txtPos))
	}

	if win.Pressed(pixelgl.Key1) && event.WeaponPurchaseEventInstance.Weapon1 {
		currentweaponState = weapon1
	} else if win.Pressed(pixelgl.Key2) && event.WeaponPurchaseEventInstance.Weapon2 {
		currentweaponState = weapon2
	} else if win.Pressed(pixelgl.Key3) && event.WeaponPurchaseEventInstance.Weapon3 {
		currentweaponState = weapon3
	} else if win.Pressed(pixelgl.Key4) && event.WeaponPurchaseEventInstance.Weapon4 {
		currentweaponState = weapon4
	} else if win.Pressed(pixelgl.Key5) && event.WeaponPurchaseEventInstance.Weapon5 {
		currentweaponState = weapon5
	} else if win.Pressed(pixelgl.Key6) && event.WeaponPurchaseEventInstance.Weapon6 {
		currentweaponState = weapon6
	} else if win.Pressed(pixelgl.Key7) && event.WeaponPurchaseEventInstance.Weapon7 {
		currentweaponState = weapon7
	} else if win.Pressed(pixelgl.Key8) && event.WeaponPurchaseEventInstance.Weapon8 {
		currentweaponState = weapon8
	} else if win.Pressed(pixelgl.Key9) && event.WeaponPurchaseEventInstance.Weapon9 {
		currentweaponState = weapon9
	} else if win.Pressed(pixelgl.Key0) && event.WeaponPurchaseEventInstance.Weapon0 {
		currentweaponState = weapon0
	}

	switch currentweaponState {
	case weapon1:
		DescriptionWeapon(win, descWeapon, 0)
	case weapon2:
		DescriptionWeapon(win, descWeapon, 1)
	case weapon3:
		DescriptionWeapon(win, descWeapon, 2)
	case weapon4:
		DescriptionWeapon(win, descWeapon, 3)
	case weapon5:
		DescriptionWeapon(win, descWeapon, 4)
	case weapon6:
		DescriptionWeapon(win, descWeapon, 5)
	case weapon7:
		DescriptionWeapon(win, descWeapon, 6)
	case weapon8:
		DescriptionWeapon(win, descWeapon, 7)
	case weapon9:
		DescriptionWeapon(win, descWeapon, 8)
	case weapon0:
		DescriptionWeapon(win, descWeapon, 9)
	}
}

func WeaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	var tempWeapon string
	if (buttonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) && event.WeaponPurchaseEventInstance.Weapon1 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon1
		log.Println("WeaponShop->weapon1")
	} else if (buttonSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.Key2)) && event.WeaponPurchaseEventInstance.Weapon2 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon2
		CreateWeaponEvent(descWeapon, 0)
		log.Println("WeaponShop->weapon2")
	} else if (buttonSlice[2].Contains(mousePos) || win.JustPressed(pixelgl.Key3)) && event.WeaponPurchaseEventInstance.Weapon3 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon3
		log.Println("WeaponShop->weapon3")
	} else if (buttonSlice[3].Contains(mousePos) || win.JustPressed(pixelgl.Key4)) && event.WeaponPurchaseEventInstance.Weapon4 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon4
		log.Println("WeaponShop->weapon4")
	} else if (buttonSlice[4].Contains(mousePos) || win.JustPressed(pixelgl.Key5)) && event.WeaponPurchaseEventInstance.Weapon5 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon5
		log.Println("WeaponShop->weapon5")
	} else if (buttonSlice[5].Contains(mousePos) || win.JustPressed(pixelgl.Key6)) && event.WeaponPurchaseEventInstance.Weapon6 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon6
		log.Println("WeaponShop->weapon6")
	} else if (buttonSlice[6].Contains(mousePos) || win.JustPressed(pixelgl.Key7)) && event.WeaponPurchaseEventInstance.Weapon7 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon7
		log.Println("WeaponShop->weapon7")
	} else if (buttonSlice[7].Contains(mousePos) || win.JustPressed(pixelgl.Key8)) && event.WeaponPurchaseEventInstance.Weapon8 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon8
		log.Println("WeaponShop->weapon8")
	} else if (buttonSlice[8].Contains(mousePos) || win.JustPressed(pixelgl.Key9)) && event.WeaponPurchaseEventInstance.Weapon9 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon9
		log.Println("WeaponShop->weapon9")
	} else if (buttonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.WeaponPurchaseEventInstance.Weapon0 && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon0
		log.Println("WeaponShop->weapon0")
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.WeaponShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("WeaponShop->TownScreen")
	}
	if (buySellSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		//TODO: お金が足りないときの処理を記述
		if currentweaponState == weapon1 {
			//TODO: ファイルからweaponのインスタンス作成、weaponインスタンスからGoldを参照
			player.Gold -= 100
			tempWeapon = "weapon1"
		} else if currentweaponState == weapon2 {
			player.Gold -= 300
			tempWeapon = "weapon2"
		} else if currentweaponState == weapon3 {
			player.Gold -= 500
			tempWeapon = "weapon3"
		} else if currentweaponState == weapon4 {
			player.Gold -= 1500
			tempWeapon = "weapon4"
		} else if currentweaponState == weapon5 {
			player.Gold -= 3000
			tempWeapon = "weapon5"
		} else if currentweaponState == weapon6 {
			player.Gold -= 2000
			tempWeapon = "weapon6"
		} else if currentweaponState == weapon7 {
			player.Gold -= 4000
			tempWeapon = "weapon7"
		} else if currentweaponState == weapon8 {
			player.Gold -= 6000
			tempWeapon = "weapon8"
		} else if currentweaponState == weapon9 {
			player.Gold -= 8000
			tempWeapon = "weapon9"
		} else if currentweaponState == weapon0 {
			player.Gold -= 10000
			tempWeapon = "weapon0"
		}
		SaveWeaponPurchaseEvent(SaveFilePath, 3, tempWeapon, player)
		SaveGame(SaveFilePath, 1, player)
	}
	//  else if buySellSlice[1].Contains(mousePos) || win.JustPressed(pixelgl.KeyS) {
	// 	log.Println("Sell")
	// 	for _, value := range player.PossessedWeapon {
	// 		tempPossessedWeapon, _ := strconv.Atoi(value)
	// 		tempInt = append(tempInt, tempPossessedWeapon)
	// 	}
	// 	if currentweaponState == weapon1 && tempInt[0] >= 1 {
	// 		//TODO: ファイルからweaponのインスタンス作成、weaponインスタンスからGoldを参照
	// 		player.Gold += 50
	// 		tempWeapon = "weapon1"
	// 	} else if currentweaponState == weapon2 && tempInt[1] >= 1 {
	// 		player.Gold += 150
	// 		tempWeapon = "weapon2"
	// 	} else if currentweaponState == weapon3 && tempInt[2] >= 1 {
	// 		player.Gold += 250
	// 		tempWeapon = "weapon3"
	// 	} else if currentweaponState == weapon4 && tempInt[3] >= 1 {
	// 		player.Gold += 750
	// 		tempWeapon = "weapon4"
	// 	} else if currentweaponState == weapon5 && tempInt[4] >= 1 {
	// 		player.Gold += 1500
	// 		tempWeapon = "weapon5"
	// 	} else if currentweaponState == weapon6 && tempInt[5] >= 1 {
	// 		player.Gold += 1000
	// 		tempWeapon = "weapon6"
	// 	} else if currentweaponState == weapon7 && tempInt[6] >= 1 {
	// 		player.Gold += 2000
	// 		tempWeapon = "weapon7"
	// 	} else if currentweaponState == weapon8 && tempInt[7] >= 1 {
	// 		player.Gold += 3000
	// 		tempWeapon = "weapon8"
	// 	} else if currentweaponState == weapon9 && tempInt[8] >= 1 {
	// 		player.Gold += 4000
	// 		tempWeapon = "weapon9"
	// 	} else if currentweaponState == weapon0 && tempInt[9] >= 1 {
	// 		player.Gold += 5000
	// 		tempWeapon = "weapon0"
	// 	}
	// 	SaveWeaponSellEvent(SaveFilePath, 3, tempWeapon, player)
	// 	SaveGame(SaveFilePath, 1, player)
	// 	tempInt = tempInt[:0]
	// }
	return myState.CurrentGS
}

func InitWeaponBelongScreen(win *pixelgl.Window, Txt *text.Text) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/武器"
	InitWeaponBelong(win, Txt, botText)
}

func InitWeaponBelong(win *pixelgl.Window, Txt *text.Text, botText string) {
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

func WeaponBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec) myState.GameState {
	if myState.CurrentGS == myState.GoToScreen && (gotoButtonSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.Key1)) {
		myState.CurrentGS = myState.StageSelect
		log.Println("所持品/武器->ステージセレクト")
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
		log.Println("所持品/武器->GoTo")
	}
	return myState.CurrentGS
}

func DescriptionWeapon(win *pixelgl.Window, descWeapon [][]string, num int) {
	//TODO: Tabを押している間は強化素材等の情報を表示する
	num++
	xOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).X + 300
	yOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).Y - 50
	txtPos := pixel.V(0, 0)

	myUtil.DescriptionTxt.Color = colornames.White

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[0][1]+": "+descWeapon[num][1], "   カラー: "+descWeapon[num][17])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[0][2]+": "+descWeapon[num][2], descWeapon[0][3]+": "+descWeapon[num][3])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[0][4]+": "+descWeapon[num][4]+"S ")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "素材: "+descWeapon[num][5], descWeapon[num][6]+"個, ", descWeapon[num][7], descWeapon[num][8]+"個, ", descWeapon[num][9], descWeapon[num][10]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "説明: "+descWeapon[num][11])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descWeapon[num][12])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "特殊能力: "+descWeapon[num][14])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descWeapon[num][15])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[num][16])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	myUtil.DescriptionTxt.Color = colornames.White
	fmt.Fprintln(myUtil.DescriptionTxt, "B. 作ってもらう")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, myUtil.DescriptionTxt.Bounds().Moved(txtPos))
}

func CreateWeaponEvent(descWeapon [][]string, num int) {
	//TODO: 素材が足りるかどうかの判定実装中
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	for name, count := range tempSlice {
		if name == descWeapon[num][5] {
			tempCount, _ := strconv.Atoi(descWeapon[num][6])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
			}
		}
		if name == descWeapon[num][7] {
			tempCount, _ := strconv.Atoi(descWeapon[num][8])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
			}
		}
		if name == descWeapon[num][9] {
			tempCount, _ := strconv.Atoi(descWeapon[num][10])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
			}
		}
	}
}
