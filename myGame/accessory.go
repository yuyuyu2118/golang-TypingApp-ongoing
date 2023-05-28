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

type AccessoryState int

const (
	accessoryNil AccessoryState = iota
	accessory1
	accessory2
	accessory3
	accessory4
	accessory5
	accessory6
	accessory7
	accessory8
	accessory9
	accessory10
)

var keyToAccessory = map[pixelgl.Button]AccessoryState{
	pixelgl.Key1: accessory1,
	pixelgl.Key2: accessory2,
	pixelgl.Key3: accessory3,
	pixelgl.Key4: accessory4,
	pixelgl.Key5: accessory5,
	pixelgl.Key6: accessory6,
	pixelgl.Key7: accessory7,
	pixelgl.Key8: accessory8,
	pixelgl.Key9: accessory9,
	pixelgl.Key0: accessory10,
}

var accessorySlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var accessoryNum = []string{"accessory0", "accessory1", "accessory2", "accessory3", "accessory4", "accessory5", "accessory6", "accessory7", "accessory8", "accessory9"}
var accessoryName = []string{"樹木のペンダント", "フルーツブレスレット", "平和のバンド", "ライフリンクのリング", "ショックウェーブリング", "鉄のブレスレット", "疾走のリング", "勇者のペンダント", "刀匠の指輪", "霊魂のイヤリング"}

var (
	accessoryPath = "assets/shop/accessory.csv"
	descAccessory = CsvToSlice(accessoryPath)
)
var currentaccessoryState AccessoryState

func InitAccessory(win *pixelgl.Window, Txt *text.Text, botText string) {
	xOffSet, yOffSet, txtPos := myUtil.ShopInitAndText(win, myUtil.ScreenTxt, botText)

	for i, v := range accessoryName {
		if event.AccessoryPurchaseEventInstance.Accessorys[i] {
			accessorySlice[i] = strconv.Itoa(i+1) + ". " + v
		}
	}
	if event.AccessoryPurchaseEventInstance.Accessorys[9] {
		accessorySlice[9] = "0. " + accessoryName[9]
	}

	buttonSliceAccessory = myUtil.DisplayShopLineup(win, accessorySlice, buttonSliceAccessory, 30.0, colornames.White, myUtil.DescriptionTxt, xOffSet, yOffSet, txtPos)

	for i := 0; i < len(keyToAccessory)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.AccessoryPurchaseEventInstance.Accessorys[i] {
			currentaccessoryState = AccessoryState(i + 1)
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.AccessoryPurchaseEventInstance.Accessorys[9] {
		currentaccessoryState = accessory10
	}
	if currentaccessoryState >= accessory1 && currentaccessoryState <= accessory10 {
		DescriptionAccessory(win, descAccessory, int(currentaccessoryState)-1)
	}
}

func AccessoryClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) myState.GameState {
	var tempAccessory = ""

	for i := 0; i < len(keyToAccessory)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (buttonSliceAccessory[i].Contains(mousePos) || win.Pressed(key)) && event.AccessoryPurchaseEventInstance.Accessorys[i] && myState.CurrentGS == myState.AccessoryShop {
			currentaccessoryState = AccessoryState(i + 1)
			//CreateAccessoryEvent(descAccessory, 0)
			log.Println("アクセサリー屋->アクセサリー", i+1)
			break
		}
	}

	if (buttonSliceAccessory[9].Contains(mousePos) || win.JustPressed(pixelgl.Key0)) && event.AccessoryPurchaseEventInstance.Accessorys[9] && myState.CurrentGS == myState.AccessoryShop {
		currentaccessoryState = accessory10
		log.Println("アクセサリー屋->アクセサリー10")
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.AccessoryShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("アクセサリー屋->町")
	}

	if (buySellSliceAccessory[0].Contains(mousePos) || win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		loadContent := SaveFileLoad(SaveFilePath)
		//TODO: お金が足りないときの処理を記述
		for i := 0; i < len(keyToAccessory)-1; i++ {
			if currentaccessoryState == AccessoryState(i+1) {
				requiredGold, _ := strconv.Atoi(descAccessory[i+1][5])
				belongAccessory, _ := strconv.Atoi(loadContent[5][i])
				//log.Println(loadContent)
				log.Println(belongAccessory)
				if belongAccessory == 0 {
					if player.Gold >= requiredGold {
						log.Println(descAccessory[i+1][5], "買える", player.Gold)
						createOk := CreateAccessoryEvent(descAccessory, i)
						if createOk {
							player.Gold -= requiredGold
							tempAccessory = "accessory" + strconv.Itoa(i+1)
						}
					} else {
						log.Println(descAccessory[i+1][5], "お金が足りない", player.Gold)
					}
				} else {
					log.Println("すでに持っている")
					break
				}
			}
		}
		if currentaccessoryState == accessory10 {
			requiredGold, _ := strconv.Atoi(descAccessory[10][5])
			if player.Gold >= requiredGold {
				log.Println(descAccessory[10][5], "買える", player.Gold)
			} else {
				log.Println(descAccessory[10][5], "お金が足りない", player.Gold)
			}
			log.Println(descAccessory[10][5])
			tempAccessory = "accessory" + strconv.Itoa(10)
		}

		if tempAccessory != "" {
			SaveAccessoryPurchaseEvent(SaveFilePath, 5, tempAccessory, player)
			SaveGame(SaveFilePath, 1, player)
		}
	}

	return myState.CurrentGS
}

func DescriptionAccessory(win *pixelgl.Window, descAccessory [][]string, num int) {
	//TODO: Tabを押している間は強化素材等の情報を表示する
	//TODO: 行数削減したい
	num++
	xOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).X + 300
	yOffSet := myPos.TopLefPos(win, myUtil.DescriptionTxt).Y - 50
	txtPos := pixel.V(0, 0)

	myUtil.DescriptionTxt.Color = colornames.White

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descAccessory[0][1]+": "+descAccessory[num][1], "   カラー: "+descAccessory[num][18])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descAccessory[0][2]+": "+descAccessory[num][2], descAccessory[0][3]+": "+descAccessory[num][3], descAccessory[0][4]+": "+descAccessory[num][4])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descAccessory[0][5]+": "+descAccessory[num][5]+"S ")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "素材: "+descAccessory[num][6], descAccessory[num][7]+"個, ", descAccessory[num][8], descAccessory[num][9]+"個, ", descAccessory[num][10], descAccessory[num][11]+"個")
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 30
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "説明: "+descAccessory[num][12])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descAccessory[num][13])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, "特殊能力: "+descAccessory[num][15])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 50
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, " "+descAccessory[num][16])
	yOffSet -= myUtil.DescriptionTxt.LineHeight + 10
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition = pixel.IM.Moved(txtPos)
	myUtil.DescriptionTxt.Draw(win, tempPosition)

	myUtil.DescriptionTxt.Clear()
	fmt.Fprintln(myUtil.DescriptionTxt, descAccessory[num][17])
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
	buySellSliceAccessory = append(buySellSliceAccessory, myUtil.DescriptionTxt.Bounds().Moved(txtPos))
}

func CreateAccessoryEvent(descAccessory [][]string, num int) bool {
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		if name == descAccessory[num][6] {
			tempCount, _ := strconv.Atoi(descAccessory[num][7])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		}
		if name == descAccessory[num][8] {
			tempCount, _ := strconv.Atoi(descAccessory[num][9])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		}
		if name == descAccessory[num][10] {
			tempCount, _ := strconv.Atoi(descAccessory[num][11])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descAccessory[num][6] {
				tempCount, _ := strconv.Atoi(descAccessory[num][7])
				tempSlice[name] -= tempCount
			}
			if name == descAccessory[num][8] {
				tempCount, _ := strconv.Atoi(descAccessory[num][9])
				tempSlice[name] -= tempCount
			}
			if name == descAccessory[num][10] {
				tempCount, _ := strconv.Atoi(descAccessory[num][11])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費してアクセサリーを作成しました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}

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
	fmt.Fprintln(myUtil.ScreenTxt, botText, "Tabで切り替え", "BackSpace.戻る")
	tempPosition = myPos.BotCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)

	loadContent := SaveFileLoad(SaveFilePath)
	counts := make(map[string]int)
	elements := loadContent[5]

	for i, val := range elements {
		num, err := strconv.Atoi(val)
		if err == nil {
			accessoryKey := fmt.Sprintf("accessory%d", i)
			counts[accessoryKey] = num
		}
	}

	for i, value := range accessoryName {
		if counts["accessory"+strconv.Itoa(i)] != 0 {
			tempInt := counts["accessory"+strconv.Itoa(i)]
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
