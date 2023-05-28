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
		DescriptionWeapon(win, descWeapon, int(currentweaponState)-1)
	}
}

func WeaponClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	var tempWeapon = ""

	for i := 0; i < len(keyToWeapon)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (buttonSlice[i].Contains(mousePos) || win.Pressed(key)) && event.WeaponPurchaseEventInstance.Weapons[i] && myState.CurrentGS == myState.WeaponShop {
			currentweaponState = WeaponState(i + 1)
			//CreateWeaponEvent(descWeapon, 0)
			log.Println("武器屋->武器", i+1)
			break
		}
	}

	if (buttonSlice[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.WeaponPurchaseEventInstance.Weapons[9] && myState.CurrentGS == myState.WeaponShop {
		currentweaponState = weapon10
		log.Println("武器屋->武器10")
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.WeaponShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("武器屋->町")
	}

	if (buySellSlice[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		loadContent := SaveFileLoad(SaveFilePath)
		//TODO: お金が足りないときの処理を記述
		for i := 0; i < len(keyToWeapon)-1; i++ {
			if currentweaponState == WeaponState(i+1) {
				requiredGold, _ := strconv.Atoi(descWeapon[i+1][4])
				belongWeapon, _ := strconv.Atoi(loadContent[3][i])
				if belongWeapon == 0 {
					if player.Gold >= requiredGold {
						log.Println(descWeapon[i+1][4], "買える", player.Gold)
						createOk := CreateWeaponEvent(descWeapon, i)
						if createOk {
							player.Gold -= requiredGold
							tempWeapon = "weapon" + strconv.Itoa(i+1)
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

func CreateWeaponEvent(descWeapon [][]string, num int) bool {
	//TODO: 素材が足りるかどうかの判定実装中
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		if name == descWeapon[num][5] {
			tempCount, _ := strconv.Atoi(descWeapon[num][6])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		}
		if name == descWeapon[num][7] {
			tempCount, _ := strconv.Atoi(descWeapon[num][8])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		}
		if name == descWeapon[num][9] {
			tempCount, _ := strconv.Atoi(descWeapon[num][10])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
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
