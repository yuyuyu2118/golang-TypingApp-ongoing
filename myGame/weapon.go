package myGame

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
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
	weapon10
)

var keyToWeapon = map[pixelgl.Button]WeaponState{
	pixelgl.Key1: weapon1,
	pixelgl.Key2: weapon2,
	pixelgl.Key3: weapon3,
	pixelgl.Key4: weapon4,
	pixelgl.Key5: weapon5,
	pixelgl.Key6: weapon6,
	pixelgl.Key7: weapon7,
	pixelgl.Key8: weapon8,
	pixelgl.Key9: weapon9,
	pixelgl.Key0: weapon10,
}

var weaponSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var weaponNum = []string{"weapon0", "weapon1", "weapon2", "weapon3", "weapon4", "weapon5", "weapon6", "weapon7", "weapon8", "weapon9"}
var weaponName = []string{"木の棒", "果物ナイフ", "木刀", "ドレインソード", "スタンハンマー", "鉄の剣", "隼の剣", "勇者の剣", "名刀村正", "死神の大鎌"}

var (
	weaponPath = "assets/shop/weapon.csv"
	descWeapon = CsvToSlice(weaponPath)
)
var currentweaponState WeaponState

func InitWeapon(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet, yOffSet, txtPos := myUtil.ShopInitAndText(win, myUtil.ScreenTxt, botText)

	for i, v := range weaponName {
		if event.WeaponPurchaseEventInstance.Weapons[i] {
			weaponSlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}
	if event.WeaponPurchaseEventInstance.Weapons[9] {
		weaponSlice[9] = "0. " + weaponName[9]
	}

	buttonSlice = myUtil.DisplayShopLineup(win, weaponSlice, buttonSlice, 30.0, colornames.White, myUtil.DescriptionTxt, xOffSet, yOffSet, txtPos)

	for i := 0; i < len(keyToWeapon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.WeaponPurchaseEventInstance.Weapons[i] {
			currentweaponState = WeaponState(i + 1)
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.WeaponPurchaseEventInstance.Weapons[9] {
		currentweaponState = weapon10
	}
	if currentweaponState >= weapon1 && currentweaponState <= weapon10 {
		if win.Pressed(pixelgl.KeyTab) {
			SubDescriptionWeapon(win, descWeapon, int(currentweaponState)-1)
		} else {
			DescriptionWeapon(win, descWeapon, int(currentweaponState)-1)
		}
	}
}

func WeaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {
	var tempWeapon = ""

	for i := 0; i < len(keyToWeapon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.WeaponPurchaseEventInstance.Weapons[i] && myState.CurrentGS == myState.WeaponShop {
			currentweaponState = WeaponState(i + 1)
			//CreateWeaponEvent(descWeapon, 0)
			log.Println("武器屋->武器", i+1)
			tempMyMaterialBool = false
			tempMyMaterialName = []string{"", "", "", "", "", ""}
			tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
			break
		}
	}

	if (win.JustPressed(pixelgl.Key0)) && event.WeaponPurchaseEventInstance.Weapons[9] && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon10
		log.Println("武器屋->武器10")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", "", "", "", "", ""}
		tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.WeaponShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("武器屋->町")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", "", "", "", "", ""}
		tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
	}

	if len(buySellSlice) > 0 {
		if (win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
			loadContent := SaveFileLoad(SaveFilePath)
			//TODO: お金が足りないときの処理を記述
			for i := 0; i < len(keyToWeapon)-1; i++ {
				if currentweaponState == WeaponState(i+1) {
					requiredGold, _ := strconv.Atoi(descWeapon[i+1][4])
					belongWeapon, _ := strconv.Atoi(loadContent[3][i])
					if belongWeapon == 0 {
						if player.Gold >= requiredGold {
							log.Println(descWeapon[i+1][4], "買える", player.Gold)
							createOk := CreateWeaponEvent(win, descWeapon, i)
							if createOk {
								player.Gold -= requiredGold
								tempWeapon = "weapon" + strconv.Itoa(i+1)
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
			if currentweaponState == weapon10 {
				requiredGold, _ := strconv.Atoi(descWeapon[10][4])
				if player.Gold >= requiredGold {
					log.Println(descWeapon[10][4], "買える", player.Gold)
				} else {
					log.Println(descWeapon[10][4], "お金が足りない", player.Gold)
				}
				log.Println(descWeapon[10][4])
				tempWeapon = "weapon" + strconv.Itoa(10)
			}

			if tempWeapon != "" {
				SaveWeaponPurchaseEvent(SaveFilePath, 3, tempWeapon, player)
				SaveGame(SaveFilePath, 1, player)
			}
		}
	}
	return myState.CurrentGS
}

var tempMyMaterialBool = false
var tempMyMaterialName = []string{"", "", "", "", "", ""}
var tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}

func DescriptionWeapon(win *pixelgl.Window, descWeapon [][]string, num int) {
	loadContent = SaveFileLoad(SaveFilePath)
	temp, _ := CountMyItems(SaveFilePathItems)

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
	fmt.Fprintln(myUtil.DescriptionTxt, "素材: "+descWeapon[num][5], descWeapon[num][6]+"個, ", descWeapon[num][7], descWeapon[num][8]+"個")
	//fmt.Fprintln(myUtil.DescriptionTxt, "所持: "+descWeapon[num][5], tempMaterials[0]+"個, ", descWeapon[num][7], tempMaterials[1]+"個, ", descWeapon[num][9], tempMaterials[2]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	if !tempMyMaterialBool {
		tempMyMaterialName[0] = descWeapon[num][5]
		tempMyMaterialName[1] = descWeapon[num][7]
		for name, count := range temp {
			if name == descWeapon[num][5] {
				tempMyMaterialName[0] = name
				tempMyMaterialCount[0] = count
			} else if name == descWeapon[num][7] {
				tempMyMaterialName[1] = name
				tempMyMaterialCount[1] = count
			} else if name == descWeapon[num][9] {
				tempMyMaterialName[2] = name
				tempMyMaterialCount[2] = count
			}
		}
		tempMyMaterialBool = true
	}

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "所持:", tempMyMaterialName[0], strconv.Itoa(tempMyMaterialCount[0])+"個,", tempMyMaterialName[1], strconv.Itoa(tempMyMaterialCount[1])+"個")
	//fmt.Fprintln(myUtil.DescriptionTxt, "所持: "+descWeapon[num][5], tempMaterials[0]+"個, ", descWeapon[num][7], tempMaterials[1]+"個, ", descWeapon[num][9], tempMaterials[2]+"個")
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
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, myUtil.DescriptionTxt.Bounds().Moved(txtPos))

	if loadContent[3][num-1] == strconv.Itoa(1) {
		myUtil.DescriptionTxt.Clear()
		myUtil.DescriptionTxt.Color = colornames.White
		fmt.Fprintln(myUtil.DescriptionTxt, "作成済み")
		xOffSet += 400
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition = pixel.IM.Moved(txtPos)
		myUtil.DescriptionTxt.Draw(win, tempPosition)
	} else {
		myUtil.DescriptionTxt.Clear()
		myUtil.DescriptionTxt.Color = colornames.White
		fmt.Fprintln(myUtil.DescriptionTxt)
		xOffSet += 400
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition = pixel.IM.Moved(txtPos)
		myUtil.DescriptionTxt.Draw(win, tempPosition)
	}
}

func SubDescriptionWeapon(win *pixelgl.Window, descWeapon [][]string, num int) {
	//TODO: Tabを押している間は強化素材等の情報を表示する
	num++
	xOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).X + 300
	yOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).Y - 50
	txtPos := pixel.V(0, 0)

	myUtil.DescriptionTxt.Color = colornames.White

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[0][18]+": "+descWeapon[num][18], "攻撃上昇値: ", descWeapon[num][25])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "強化素材: "+descWeapon[num][19], descWeapon[num][20]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "宝石アンロック(1回目):")
	fmt.Fprintln(myUtil.DescriptionTxt, "攻撃上昇値:", descWeapon[num][34], "アタックタイマー上昇値:", descWeapon[num][35])
	if descWeapon[num][26] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "奇跡の石: "+descWeapon[num][26]+"個 ")
	}
	if descWeapon[num][27] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "魔法のジェム: "+descWeapon[num][27]+"個 ")
	}
	if descWeapon[num][28] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "蒼き宝石: "+descWeapon[num][28]+"個 ")
	}
	if descWeapon[num][29] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "聖なる宝玉: "+descWeapon[num][29]+"個 ")
	}
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "宝石アンロック(2回目):")
	fmt.Fprintln(myUtil.DescriptionTxt, " 攻撃上昇値:", descWeapon[num][36], "アタックタイマー上昇値:", descWeapon[num][37])
	if descWeapon[num][30] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "奇跡の石: "+descWeapon[num][30]+"個 ")
	}
	if descWeapon[num][31] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "魔法のジェム: "+descWeapon[num][31]+"個 ")
	}
	if descWeapon[num][32] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "蒼き宝石: "+descWeapon[num][32]+"個 ")
	}
	if descWeapon[num][33] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "聖なる宝玉: "+descWeapon[num][33]+"個 ")
	}
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 80
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "モディファイア抽選費用: "+descWeapon[num][40]+"S")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 80
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

func CreateWeaponEvent(win *pixelgl.Window, descWeapon [][]string, num int) bool {
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if name == descWeapon[num][5] {
			tempCount, _ := strconv.Atoi(descWeapon[num][6])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		} else if (descWeapon)[num][5] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descWeapon[num][7] {
			tempCount, _ := strconv.Atoi(descWeapon[num][8])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descWeapon)[num][7] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descWeapon[num][9] {
			tempCount, _ := strconv.Atoi(descWeapon[num][10])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descWeapon)[num][9] == "" {
			//log.Println("なし")
			tempBool[2] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descWeapon[num][5] {
				tempCount, _ := strconv.Atoi(descWeapon[num][6])
				tempSlice[name] -= tempCount
			}
			if name == descWeapon[num][7] {
				tempCount, _ := strconv.Atoi(descWeapon[num][8])
				tempSlice[name] -= tempCount
			}
			if name == descWeapon[num][9] {
				tempCount, _ := strconv.Atoi(descWeapon[num][10])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して武器を作成しました。")

		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}

}

func InitWeaponBelongScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/武器"
	InitWeaponBelong(win, Txt, botText, player)
}

func InitWeaponBelong(win *pixelgl.Window, Txt *text.Text, botText string, player *myPlayer.PlayerStatus) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, botText, "Tabで切り替え", "BackSpace.戻る")
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
			//tempInt := counts["weapon"+strconv.Itoa(i)]
			equipmentSlice = append(equipmentSlice, strconv.Itoa(i+1)+". "+value /*+": "+strconv.Itoa(tempInt)*/)
		} else {
			equipmentSlice = append(equipmentSlice, "")
		}
	}
	for i, equipmentName := range equipmentSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, equipmentName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		equipmentButtonSlice = append(equipmentButtonSlice, Txt.Bounds().Moved(txtPos))

		if player.EquipmentWeapon[0] == weaponName[i] {
			Txt.Clear()
			fmt.Fprintln(Txt, "E. ")
			txtPos = pixel.V(xOffSet-40, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
	equipmentSlice = equipmentSlice[:0]
}

func WeaponBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) {
	loadContent := SaveFileLoad(SaveFilePath)

	if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key1)) && (player.PossessedWeapon[0] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[1][1])
		player.EquipmentWeapon[1] = descWeapon[1][2]
		player.EquipmentWeapon[3] = descWeapon[1][3]
		log.Println("装備1")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key2)) && (player.PossessedWeapon[1] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[2][1])
		player.EquipmentWeapon[1] = descWeapon[2][2]
		player.EquipmentWeapon[3] = descWeapon[2][3]
		log.Println("装備2")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key3)) && (player.PossessedWeapon[2] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[3][1])
		player.EquipmentWeapon[1] = descWeapon[3][2]
		player.EquipmentWeapon[3] = descWeapon[3][3]
		log.Println("装備3")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key4)) && (player.PossessedWeapon[3] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[4][1])
		player.EquipmentWeapon[1] = descWeapon[4][2]
		player.EquipmentWeapon[3] = descWeapon[4][3]
		log.Println("装備4")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key5)) && (player.PossessedWeapon[4] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[5][1])
		player.EquipmentWeapon[1] = descWeapon[5][2]
		player.EquipmentWeapon[3] = descWeapon[5][3]
		log.Println("装備5")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key6)) && (player.PossessedWeapon[5] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[6][1])
		player.EquipmentWeapon[1] = descWeapon[6][2]
		player.EquipmentWeapon[3] = descWeapon[6][3]
		log.Println("装備6")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key7)) && (player.PossessedWeapon[6] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[7][1])
		player.EquipmentWeapon[1] = descWeapon[7][2]
		player.EquipmentWeapon[3] = descWeapon[7][3]
		log.Println("装備7")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key8)) && (player.PossessedWeapon[7] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[8][1])
		player.EquipmentWeapon[1] = descWeapon[8][2]
		player.EquipmentWeapon[3] = descWeapon[8][3]
		log.Println("装備8")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key9)) && (player.PossessedWeapon[8] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[9][1])
		player.EquipmentWeapon[1] = descWeapon[9][2]
		player.EquipmentWeapon[3] = descWeapon[9][3]
		log.Println("装備9")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key0)) && (player.PossessedWeapon[9] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[10][1])
		player.EquipmentWeapon[1] = descWeapon[10][2]
		player.EquipmentWeapon[3] = descWeapon[10][3]
		log.Println("装備0")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.WeaponBelong
		myState.CurrentGS = myState.TownScreen
		log.Println("所持品/武器->GoTo")
	}
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

	tempAttackTimer1, _ := strconv.ParseFloat(loadContent[1][15], 64)
	tempAttackTimer2, _ := strconv.ParseFloat(player.EquipmentWeapon[3], 64)
	tempAttackTimer3, _ := strconv.ParseFloat(player.EquipmentArmor[3], 64)
	tempAttackTimer4, _ := strconv.ParseFloat(player.EquipmentAccessory[3], 64)
	player.AttackTimer = tempAttackTimer1 + tempAttackTimer2 + tempAttackTimer3 + tempAttackTimer4

	SaveGame(SaveFilePath, 1, player)
	SaveGameWeapon(SaveFilePath, 6, player)
	SaveWeaponEnhancementEvent(SaveFilePath, 9, tempName, player)
	log.Println("weapon", player.OP)
}

func InitWeaponBlackSmithScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "鍛冶/武器"
	InitWeaponBlackSmith(win, Txt, botText, player)
}

func InitWeaponBlackSmith(win *pixelgl.Window, Txt *text.Text, botText string, player *myPlayer.PlayerStatus) {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, Txt).Y - 100
	txtPos := pixel.V(0, 0)

	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, botText, "Tabで切り替え", "BackSpace.戻る")
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
			//tempInt := counts["weapon"+strconv.Itoa(i)]
			blackSmithSlice = append(blackSmithSlice, strconv.Itoa(i+1)+". "+value /*+": "+strconv.Itoa(tempInt)*/)
		} else {
			blackSmithSlice = append(blackSmithSlice, "")
		}
	}
	for i, blackSmithName := range blackSmithSlice {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, blackSmithName)
		yOffSet -= Txt.LineHeight + 25
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		Txt.Draw(win, tempPosition)
		blackSmithButtonSlice = append(blackSmithButtonSlice, Txt.Bounds().Moved(txtPos))

		if player.EquipmentWeapon[0] == weaponName[i] {
			Txt.Clear()
			fmt.Fprintln(Txt, "E. ")
			txtPos = pixel.V(xOffSet-40, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}

	for i := 0; i < len(keyToWeapon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.WeaponPurchaseEventInstance.Weapons[i] {
			currentweaponState = WeaponState(i + 1)
			tempMyMaterialBool = false
			tempMyMaterialName = []string{"", "", "", "", "", ""}
			tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.WeaponPurchaseEventInstance.Weapons[9] {
		currentweaponState = weapon10
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", "", "", "", "", ""}
		tempMyMaterialCount = []int{0, 0, 0, 0, 0, 0}
	}
	if currentweaponState >= weapon1 && currentweaponState <= weapon10 {
		if win.Pressed(pixelgl.KeyTab) {
			//SubDescriptionWeapon(win, descWeapon, int(currentweaponState)-1)
		} else {
			DescriptionWeaponBlackSmith(win, descWeapon, int(currentweaponState)-1)
		}
	}

	blackSmithSlice = blackSmithSlice[:0]
}

func WeaponBlackSmithClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) {
	//loadContent := SaveFileLoad(SaveFilePath)

	if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key1)) && (player.PossessedWeapon[0] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[1][1])
		player.EquipmentWeapon[1] = descWeapon[1][2]
		player.EquipmentWeapon[3] = descWeapon[1][3]

		log.Println("装備1", player.EquipmentWeapon[0], player.EquipmentWeapon[1], player.EquipmentWeapon[3])
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key2)) && (player.PossessedWeapon[1] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[2][1])
		player.EquipmentWeapon[1] = descWeapon[2][2]
		player.EquipmentWeapon[3] = descWeapon[2][3]
		log.Println("装備2")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key3)) && (player.PossessedWeapon[2] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[3][1])
		player.EquipmentWeapon[1] = descWeapon[3][2]
		player.EquipmentWeapon[3] = descWeapon[3][3]
		log.Println("装備3")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key4)) && (player.PossessedWeapon[3] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[4][1])
		player.EquipmentWeapon[1] = descWeapon[4][2]
		player.EquipmentWeapon[3] = descWeapon[4][3]
		log.Println("装備4")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key5)) && (player.PossessedWeapon[4] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[5][1])
		player.EquipmentWeapon[1] = descWeapon[5][2]
		player.EquipmentWeapon[3] = descWeapon[5][3]
		log.Println("装備5")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key6)) && (player.PossessedWeapon[5] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[6][1])
		player.EquipmentWeapon[1] = descWeapon[6][2]
		player.EquipmentWeapon[3] = descWeapon[6][3]
		log.Println("装備6")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key7)) && (player.PossessedWeapon[6] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[7][1])
		player.EquipmentWeapon[1] = descWeapon[7][2]
		player.EquipmentWeapon[3] = descWeapon[7][3]
		log.Println("装備7")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key8)) && (player.PossessedWeapon[7] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[8][1])
		player.EquipmentWeapon[1] = descWeapon[8][2]
		player.EquipmentWeapon[3] = descWeapon[8][3]
		log.Println("装備8")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key9)) && (player.PossessedWeapon[8] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[9][1])
		player.EquipmentWeapon[1] = descWeapon[9][2]
		player.EquipmentWeapon[3] = descWeapon[9][3]
		log.Println("装備9")
	} else if myState.CurrentBelong == myState.WeaponBelong && (win.JustPressed(pixelgl.Key0)) && (player.PossessedWeapon[9] == "1") {
		player.EquipmentWeapon[0] = strings.NewReplacer("【", "", "】", "").Replace(descWeapon[10][1])
		player.EquipmentWeapon[1] = descWeapon[10][2]
		player.EquipmentWeapon[3] = descWeapon[10][3]
		log.Println("装備0")
	} else if myState.CurrentBlackSmith == myState.WeaponBlackSmith && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBlackSmith = myState.WeaponBlackSmith
		myState.CurrentGS = myState.GoToScreen
		log.Println("鍛冶/武器->GoTo")
	}
	tempOP1, _ := strconv.ParseFloat(loadContent[1][13], 64)
	tempOP2, _ := strconv.ParseFloat(player.EquipmentWeapon[1], 64)
	tempOP3, _ := strconv.ParseFloat(player.EquipmentAccessory[1], 64)

	var tempOP4 float64
	tempName := player.EquipmentWeapon[0]
	log.Println("koko", tempName, player.WeaponEnhancement[0])
	if tempName == "木の棒" {
		coefficient, _ := strconv.ParseFloat(loadContent[9][0], 64)
		tempWeaponEnhancement, _ := strconv.ParseFloat(descWeapon[1][25], 64)
		tempOP4 = tempWeaponEnhancement * coefficient
		log.Println("kore", coefficient, tempWeaponEnhancement)
		//tempOP4, _ = strconv.ParseFloat(player.WeaponEnhancement[0], 64)
	}

	log.Println(tempOP1, tempOP2, tempOP3, tempOP4)

	player.OP = tempOP1 + tempOP2 + tempOP3 + tempOP4

	tempAttackTimer1, _ := strconv.ParseFloat(loadContent[1][15], 64)
	tempAttackTimer2, _ := strconv.ParseFloat(player.EquipmentWeapon[3], 64)
	tempAttackTimer3, _ := strconv.ParseFloat(player.EquipmentArmor[3], 64)
	tempAttackTimer4, _ := strconv.ParseFloat(player.EquipmentAccessory[3], 64)
	player.AttackTimer = tempAttackTimer1 + tempAttackTimer2 + tempAttackTimer3 + tempAttackTimer4

	SaveGame(SaveFilePath, 1, player)
	SaveGameWeapon(SaveFilePath, 6, player)
}

func DescriptionWeaponBlackSmith(win *pixelgl.Window, descWeapon [][]string, num int) {
	//loadContent = SaveFileLoad(SaveFilePath)
	temp, _ := CountMyItems(SaveFilePathItems)

	//TODO: Tabを押している間は強化素材等の情報を表示する
	num++
	xOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).X + 300
	yOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).Y - 50
	txtPos := pixel.V(0, 0)

	myUtil.DescriptionTxt.Color = colornames.White

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descWeapon[0][1]+": "+descWeapon[num][1], descWeapon[0][2]+": "+descWeapon[num][2], descWeapon[0][3]+": "+descWeapon[num][3])
	fmt.Fprintln(myUtil.DescriptionTxt, "--------------------------------------------------------------")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "強化: "+descWeapon[num][18]+"S", "攻撃+"+descWeapon[num][25], "\t", "素材: "+descWeapon[num][19], descWeapon[num][20]+"個", descWeapon[num][21], descWeapon[num][22]+"個")
	fmt.Fprintln(myUtil.DescriptionTxt, "--------------------------------------------------------------")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "宝石アンロック(1):", "攻撃+"+descWeapon[num][34], "アタックタイマー+"+descWeapon[num][35])
	if descWeapon[num][26] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "奇跡の石 "+descWeapon[num][26]+"個 ")
	}
	if descWeapon[num][27] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "魔法のジェム "+descWeapon[num][27]+"個 ")
	}
	if descWeapon[num][28] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "蒼き宝石 "+descWeapon[num][28]+"個 ")
	}
	if descWeapon[num][29] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "聖なる宝玉 "+descWeapon[num][29]+"個 ")
	}
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "宝石アンロック(2):", "攻撃+"+descWeapon[num][36], "アタックタイマー+"+descWeapon[num][37])
	if descWeapon[num][30] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "奇跡の石 "+descWeapon[num][30]+"個 ")
	}
	if descWeapon[num][31] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "魔法のジェム "+descWeapon[num][31]+"個 ")
	}
	if descWeapon[num][32] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "蒼き宝石 "+descWeapon[num][32]+"個 ")
	}
	if descWeapon[num][33] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "聖なる宝玉 "+descWeapon[num][33]+"個 ")
	}
	fmt.Fprintln(myUtil.DescriptionTxt, "\n--------------------------------------------------------------")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "モディファイア: "+descWeapon[num][40]+"S")
	fmt.Fprintln(myUtil.DescriptionTxt, "--------------------------------------------------------------")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 60
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	if !tempMyMaterialBool {
		tempMyMaterialName[0] = descWeapon[num][5]
		tempMyMaterialName[1] = descWeapon[num][7]
		for name, count := range temp {
			if name == descWeapon[num][5] {
				tempMyMaterialName[0] = name
				tempMyMaterialCount[0] = count
			} else if name == descWeapon[num][7] {
				tempMyMaterialName[1] = name
				tempMyMaterialCount[1] = count
			} else if name == "奇跡の石" {
				tempMyMaterialName[2] = name
				tempMyMaterialCount[2] = count
			} else if name == "魔法のジェム" {
				tempMyMaterialName[3] = name
				tempMyMaterialCount[3] = count
			} else if name == "蒼き宝石" {
				tempMyMaterialName[4] = name
				tempMyMaterialCount[4] = count
			} else if name == "聖なる宝玉" {
				tempMyMaterialName[5] = name
				tempMyMaterialCount[5] = count
			}
		}
		tempMyMaterialBool = true
	}

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "所持:", tempMyMaterialName[0], strconv.Itoa(tempMyMaterialCount[0])+"個,", tempMyMaterialName[1], strconv.Itoa(tempMyMaterialCount[1])+"個")
	if descWeapon[num][26] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "奇跡の石 "+strconv.Itoa(tempMyMaterialCount[2])+"個, ")
	}
	if descWeapon[num][27] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "魔法のジェム "+strconv.Itoa(tempMyMaterialCount[3])+"個, ")
	}
	if descWeapon[num][28] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "蒼き宝石 "+strconv.Itoa(tempMyMaterialCount[4])+"個, ")
	}
	if descWeapon[num][29] != strconv.Itoa(0) {
		fmt.Fprint(myUtil.DescriptionTxt, "聖なる宝玉 "+strconv.Itoa(tempMyMaterialCount[5])+"個 ")
	}
	//fmt.Fprintln(myUtil.DescriptionTxt, "所持: "+descWeapon[num][5], tempMaterials[0]+"個, ", descWeapon[num][7], tempMaterials[1]+"個, ", descWeapon[num][9], tempMaterials[2]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	xOffSet += 300

	myUtil.DescriptionTxt.Clear()
	myUtil.DescriptionTxt.Color = colornames.White
	fmt.Fprintln(myUtil.DescriptionTxt, "S. 強化")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 60
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, myUtil.DescriptionTxt.Bounds().Moved(txtPos))

	myUtil.DescriptionTxt.Clear()
	myUtil.DescriptionTxt.Color = colornames.White
	fmt.Fprintln(myUtil.DescriptionTxt, "U. ジェムアンロック")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, myUtil.DescriptionTxt.Bounds().Moved(txtPos))

	myUtil.DescriptionTxt.Clear()
	myUtil.DescriptionTxt.Color = colornames.White
	fmt.Fprintln(myUtil.DescriptionTxt, "M. モディファイア抽選")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)
	buySellSlice = append(buySellSlice, myUtil.DescriptionTxt.Bounds().Moved(txtPos))
}

func EnhancementWeaponEvent(win *pixelgl.Window, descWeapon [][]string, num int) bool {
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if name == descWeapon[num][19] {
			tempCount, _ := strconv.Atoi(descWeapon[num][20])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		} else if (descWeapon)[num][19] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descWeapon[num][21] {
			tempCount, _ := strconv.Atoi(descWeapon[num][22])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descWeapon)[num][21] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descWeapon[num][23] {
			tempCount, _ := strconv.Atoi(descWeapon[num][24])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descWeapon)[num][23] == "" {
			//log.Println("なし")
			tempBool[2] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descWeapon[num][19] {
				tempCount, _ := strconv.Atoi(descWeapon[num][20])
				tempSlice[name] -= tempCount
			}
			if name == descWeapon[num][21] {
				tempCount, _ := strconv.Atoi(descWeapon[num][22])
				tempSlice[name] -= tempCount
			}
			if name == descWeapon[num][23] {
				tempCount, _ := strconv.Atoi(descWeapon[num][24])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して武器を強化しました。")

		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}

}

func EnhancementArmorEvent(win *pixelgl.Window, descArmor [][]string, num int) bool {
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if name == descArmor[num][19] {
			tempCount, _ := strconv.Atoi(descArmor[num][20])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		} else if (descArmor)[num][19] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descArmor[num][21] {
			tempCount, _ := strconv.Atoi(descArmor[num][22])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descArmor)[num][21] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descArmor[num][23] {
			tempCount, _ := strconv.Atoi(descArmor[num][24])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descArmor)[num][23] == "" {
			//log.Println("なし")
			tempBool[2] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descArmor[num][19] {
				tempCount, _ := strconv.Atoi(descArmor[num][20])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][21] {
				tempCount, _ := strconv.Atoi(descArmor[num][22])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][23] {
				tempCount, _ := strconv.Atoi(descArmor[num][24])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して防具を強化しました。")

		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}

}

func EnhancementAccessoryEvent(win *pixelgl.Window, descAccessory [][]string, num int) bool {
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if name == descAccessory[num][20] {
			tempCount, _ := strconv.Atoi(descAccessory[num][21])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		} else if (descAccessory)[num][20] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descAccessory[num][22] {
			tempCount, _ := strconv.Atoi(descAccessory[num][23])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descAccessory)[num][22] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descAccessory[num][24] {
			tempCount, _ := strconv.Atoi(descAccessory[num][25])
			if count >= tempCount {
				//log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descAccessory)[num][24] == "" {
			//log.Println("なし")
			tempBool[2] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descAccessory[num][20] {
				tempCount, _ := strconv.Atoi(descAccessory[num][21])
				tempSlice[name] -= tempCount
			}
			if name == descAccessory[num][22] {
				tempCount, _ := strconv.Atoi(descAccessory[num][23])
				tempSlice[name] -= tempCount
			}
			if name == descAccessory[num][24] {
				tempCount, _ := strconv.Atoi(descAccessory[num][25])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費してアクセサリーを強化しました。")

		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}

}

func GemUnlockWeaponEvent(win *pixelgl.Window, descWeapon [][]string, num int, count int) bool {
	num++
	count--
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if descWeapon[num][26+4*count] != strconv.Itoa(0) {
			if name == "奇跡の石" {
				tempCount, _ := strconv.Atoi(descWeapon[num][26+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descWeapon[num][26+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descWeapon[num][27+4*count] != strconv.Itoa(0) {
			if name == "魔法のジェム" {
				tempCount, _ := strconv.Atoi(descWeapon[num][27+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descWeapon[num][27+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descWeapon[num][28+4*count] != strconv.Itoa(0) {
			if name == "蒼き宝石" {
				tempCount, _ := strconv.Atoi(descWeapon[num][28+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descWeapon[num][28+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descWeapon[num][29+4*count] != strconv.Itoa(0) {
			if name == "聖なる宝玉" {
				tempCount, _ := strconv.Atoi(descWeapon[num][29+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descWeapon[num][29+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] && tempBool[3] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if descWeapon[num][26+4*count] != strconv.Itoa(0) {
				if name == "奇跡の石" {
					tempCount, _ := strconv.Atoi(descWeapon[num][26+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
					log.Println(name, count, tempCount, tempSlice[name])
				}
			}

			if descWeapon[num][27+4*count] != strconv.Itoa(0) {
				if name == "魔法のジェム" {
					tempCount, _ := strconv.Atoi(descWeapon[num][27+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descWeapon[num][28+4*count] != strconv.Itoa(0) {
				if name == "蒼き宝石" {
					tempCount, _ := strconv.Atoi(descWeapon[num][28+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descWeapon[num][29+4*count] != strconv.Itoa(0) {
				if name == "聖なる宝玉" {
					tempCount, _ := strconv.Atoi(descWeapon[num][29+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して武器を宝石アンロックしました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}

func GemUnlockArmorEvent(win *pixelgl.Window, descArmor [][]string, num int, count int) bool {
	num++
	count--
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if descArmor[num][26+4*count] != strconv.Itoa(0) {
			if name == "奇跡の石" {
				tempCount, _ := strconv.Atoi(descArmor[num][26+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descArmor[num][26+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descArmor[num][27+4*count] != strconv.Itoa(0) {
			if name == "魔法のジェム" {
				tempCount, _ := strconv.Atoi(descArmor[num][27+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descArmor[num][27+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descArmor[num][28+4*count] != strconv.Itoa(0) {
			if name == "蒼き宝石" {
				tempCount, _ := strconv.Atoi(descArmor[num][28+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descArmor[num][28+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descArmor[num][29+4*count] != strconv.Itoa(0) {
			if name == "聖なる宝玉" {
				tempCount, _ := strconv.Atoi(descArmor[num][29+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descArmor[num][29+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] && tempBool[3] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if descArmor[num][26+4*count] != strconv.Itoa(0) {
				if name == "奇跡の石" {
					tempCount, _ := strconv.Atoi(descArmor[num][26+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
					log.Println(name, count, tempCount, tempSlice[name])
				}
			}

			if descArmor[num][27+4*count] != strconv.Itoa(0) {
				if name == "魔法のジェム" {
					tempCount, _ := strconv.Atoi(descArmor[num][27+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descArmor[num][28+4*count] != strconv.Itoa(0) {
				if name == "蒼き宝石" {
					tempCount, _ := strconv.Atoi(descArmor[num][28+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descArmor[num][29+4*count] != strconv.Itoa(0) {
				if name == "聖なる宝玉" {
					tempCount, _ := strconv.Atoi(descArmor[num][29+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して防具を宝石アンロックしました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}

func GemUnlockAccessoryEvent(win *pixelgl.Window, descAccessory [][]string, num int, count int) bool {
	num++
	count--
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false, false}

	for name, count := range tempSlice {
		//log.Println(name, count)
		if descAccessory[num][28+4*count] != strconv.Itoa(0) {
			if name == "奇跡の石" {
				tempCount, _ := strconv.Atoi(descAccessory[num][28+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descAccessory[num][28+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descAccessory[num][29+4*count] != strconv.Itoa(0) {
			if name == "魔法のジェム" {
				tempCount, _ := strconv.Atoi(descAccessory[num][29+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descAccessory[num][29+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descAccessory[num][30+4*count] != strconv.Itoa(0) {
			if name == "蒼き宝石" {
				tempCount, _ := strconv.Atoi(descAccessory[num][30+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descAccessory[num][30+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}

		if descAccessory[num][31+4*count] != strconv.Itoa(0) {
			if name == "聖なる宝玉" {
				tempCount, _ := strconv.Atoi(descAccessory[num][31+4*count])
				if count >= tempCount {
					tempBool[0] = true
				}
			}
		} else if descAccessory[num][31+4*count] == strconv.Itoa(0) {
			tempBool[0] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] && tempBool[3] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if descAccessory[num][28+4*count] != strconv.Itoa(0) {
				if name == "奇跡の石" {
					tempCount, _ := strconv.Atoi(descAccessory[num][28+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
					log.Println(name, count, tempCount, tempSlice[name])
				}
			}

			if descAccessory[num][29+4*count] != strconv.Itoa(0) {
				if name == "魔法のジェム" {
					tempCount, _ := strconv.Atoi(descAccessory[num][29+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descAccessory[num][30+4*count] != strconv.Itoa(0) {
				if name == "蒼き宝石" {
					tempCount, _ := strconv.Atoi(descAccessory[num][30+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}

			if descAccessory[num][31+4*count] != strconv.Itoa(0) {
				if name == "聖なる宝玉" {
					tempCount, _ := strconv.Atoi(descAccessory[num][31+4*count])
					if count >= tempCount {
						tempBool[0] = true
					}
					tempSlice[name] -= tempCount
				}
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費してアクセサリーを宝石アンロックしました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}
