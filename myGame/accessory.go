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
	// メッセージボックスのインスタンス生成
	accessoryMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 0.4, 0.5)
	// メッセージボックスの表示
	accessoryMessageBox.DrawMessageBox()

	// アクセサリー用メッセージ作成
	var accessoryOptions string
	for i, accessory := range accessoryName {
		// arrowIconをループの各イテレーションでリセット
		arrowIcon := ""
		if currentaccessoryState == AccessoryState(i+1) {
			arrowIcon = " ←"
		}
		if event.AccessoryPurchaseEventInstance.Accessorys[i] {
			accessoryOptions += strconv.Itoa(i+1) + ". " + accessory + arrowIcon + "\n"
		} else {
			accessoryOptions += strconv.Itoa(i+1) + ". ???\n"
		}
	}
	// 10番目のアクセサリーの表示を修正
	if event.AccessoryPurchaseEventInstance.Accessorys[9] {
		arrowIcon := "" // ここでもリセット
		if currentaccessoryState == accessory10 {
			arrowIcon = " ←"
		}
		accessoryOptions = strings.TrimSuffix(accessoryOptions, "10. ???\n") // 10番目の"????"を削除
		accessoryOptions += "0. " + accessoryName[9] + arrowIcon + "\n"
	}

	// メッセージボックスにテキストを表示
	accessoryMessageBox.DrawMessageTxt("どのアクセサリーを購入しますか？\nキーボードに対応する数字を入力してください。\n" + accessoryOptions + "\n\nBackSpaceキーでタイトルに戻る")

	// キー入力によるアクセサリー選択処理
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

	accessoryDescriptionMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0.4, 0, 1, 0.5)

	accessoryDescriptionMessageBox.DrawMessageBox()

	// アクセサリーの説明表示処理
	if currentaccessoryState >= accessory1 && currentaccessoryState <= accessory10 {
		if win.Pressed(pixelgl.KeyTab) {
			// SubDescriptionAccessory(win, descAccessory, int(currentaccessoryState)-1)
		} else {
			DescriptionAccessory(win, descAccessory, int(currentaccessoryState)-1, accessoryDescriptionMessageBox)
		}
	}
}

func AccessoryClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {
	var tempAccessory = ""

	for i := 0; i < len(keyToAccessory)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.AccessoryPurchaseEventInstance.Accessorys[i] && myState.CurrentGS == myState.AccessoryShop {
			currentaccessoryState = AccessoryState(i + 1)
			//CreateAccessoryEvent(descAccessory, 0)
			log.Println("アクセサリー屋->アクセサリー", i+1)
			tempMyMaterialBool = false
			tempMyMaterialName = []string{"", ""}
			tempMyMaterialCount = []int{0, 0}
			break
		}
	}

	if (win.JustPressed(pixelgl.Key0)) && event.AccessoryPurchaseEventInstance.Accessorys[9] && myState.CurrentGS == myState.AccessoryShop {
		currentaccessoryState = accessory10
		log.Println("アクセサリー屋->アクセサリー10")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", ""}
		tempMyMaterialCount = []int{0, 0}
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.AccessoryShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("アクセサリー屋->町")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", ""}
		tempMyMaterialCount = []int{0, 0}
	}

	if (win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
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
							tempMyMaterialBool = false
							tempMyMaterialName = []string{"", ""}
							tempMyMaterialCount = []int{0, 0}
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

func DescriptionAccessory(win *pixelgl.Window, descAccessory [][]string, num int, msgBox *myPos.MessageBox) {
	loadContent = SaveFileLoad(SaveFilePath)
	temp, _ := CountMyItems(SaveFilePathItems)
	//TODO: Tabを押している間は強化素材等の情報を表示する
	//TODO: 行数削減したい
	num++

	myUtil.DescriptionTxt.Color = colornames.White

	var accessoryDescriptionOptions string

	accessoryDescriptionOptions += descAccessory[0][1] + ": " + descAccessory[num][1] + "\n"
	accessoryDescriptionOptions += "カラー: " + descAccessory[num][18] + "\n"

	accessoryDescriptionOptions += descAccessory[0][2] + ": " + descAccessory[num][2] + "\n"
	accessoryDescriptionOptions += descAccessory[0][3] + ": " + descAccessory[num][3] + "\n"
	accessoryDescriptionOptions += descAccessory[0][4] + ": " + descAccessory[num][4] + "\n"

	accessoryDescriptionOptions += descAccessory[0][5] + ": " + descAccessory[num][5] + "S\n"

	accessoryDescriptionOptions += "素材: " + descAccessory[num][6] + " " + descAccessory[num][7] + "個, " + descAccessory[num][8] + " " + descAccessory[num][9] + "個\n"

	if !tempMyMaterialBool {
		tempMyMaterialName[0] = descAccessory[num][6]
		tempMyMaterialName[1] = descAccessory[num][8]
		for name, count := range temp {
			if name == descAccessory[num][6] {
				tempMyMaterialName[0] = name
				tempMyMaterialCount[0] = count
			} else if name == descAccessory[num][8] {
				tempMyMaterialName[1] = name
				tempMyMaterialCount[1] = count
			} else if name == descAccessory[num][10] {
				tempMyMaterialName[2] = name
				tempMyMaterialCount[2] = count
			}
		}
		tempMyMaterialBool = true
	}

	accessoryDescriptionOptions += "所持: " + tempMyMaterialName[0] + " " + strconv.Itoa(tempMyMaterialCount[0]) + "個, " + tempMyMaterialName[1] + " " + strconv.Itoa(tempMyMaterialCount[1]) + "個\n"

	accessoryDescriptionOptions += "説明: " + descAccessory[num][12] + "\n"

	accessoryDescriptionOptions += "　　  " + descAccessory[num][13] + "\n"

	accessoryDescriptionOptions += "特殊能力: " + descAccessory[num][15] + "\n"

	accessoryDescriptionOptions += "　　　　  " + descAccessory[num][16] + "\n"

	accessoryDescriptionOptions += "　　　　  " + descAccessory[num][17] + "\n\n"

	accessoryDescriptionOptions += "B. 作ってもらう\n"

	loadContent = SaveFileLoad(SaveFilePath)

	if loadContent[5][num-1] == strconv.Itoa(1) {
		accessoryDescriptionOptions += "作成済み\n"
	}

	msgBox.DrawMessageTxt(accessoryDescriptionOptions)
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
		} else if (descAccessory)[num][6] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descAccessory[num][8] {
			tempCount, _ := strconv.Atoi(descAccessory[num][9])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descAccessory)[num][8] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descAccessory[num][10] {
			tempCount, _ := strconv.Atoi(descAccessory[num][11])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descAccessory)[num][10] == "" {
			//log.Println("なし")
			tempBool[2] = true
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

func InitAccessoryBelongScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/アクセサリー"
	InitAccessoryBelong(win, Txt, botText, player)
}

func InitAccessoryBelong(win *pixelgl.Window, Txt *text.Text, botText string, player *myPlayer.PlayerStatus) {
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
			//tempInt := counts["accessory"+strconv.Itoa(i)]
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

		if player.EquipmentAccessory[0] == accessoryName[i] {
			Txt.Clear()
			fmt.Fprintln(Txt, "E. ")
			txtPos = pixel.V(xOffSet-40, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
	equipmentSlice = equipmentSlice[:0]
}

func AccessoryBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) {
	loadContent := SaveFileLoad(SaveFilePath)

	if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key1)) && (player.PossessedAccessory[0] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[1][1])
		player.EquipmentAccessory[1] = descAccessory[1][2]
		player.EquipmentAccessory[2] = descAccessory[1][3]
		player.EquipmentAccessory[3] = descAccessory[1][4]
		log.Println("装備1")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key2)) && (player.PossessedAccessory[1] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[2][1])
		player.EquipmentAccessory[1] = descAccessory[2][2]
		player.EquipmentAccessory[2] = descAccessory[2][3]
		player.EquipmentAccessory[3] = descAccessory[2][4]
		log.Println("装備2")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key3)) && (player.PossessedAccessory[2] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[3][1])
		player.EquipmentAccessory[1] = descAccessory[3][2]
		player.EquipmentAccessory[2] = descAccessory[3][3]
		player.EquipmentAccessory[3] = descAccessory[3][4]
		log.Println("装備3")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key4)) && (player.PossessedAccessory[3] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[4][1])
		player.EquipmentAccessory[1] = descAccessory[4][2]
		player.EquipmentAccessory[2] = descAccessory[4][3]
		player.EquipmentAccessory[3] = descAccessory[4][4]
		log.Println("装備4")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key5)) && (player.PossessedAccessory[4] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[5][1])
		player.EquipmentAccessory[1] = descAccessory[5][2]
		player.EquipmentAccessory[2] = descAccessory[5][3]
		player.EquipmentAccessory[3] = descAccessory[5][4]
		log.Println("装備5")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key6)) && (player.PossessedAccessory[5] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[6][1])
		player.EquipmentAccessory[1] = descAccessory[6][2]
		player.EquipmentAccessory[2] = descAccessory[6][3]
		player.EquipmentAccessory[3] = descAccessory[6][4]
		log.Println("装備6")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key7)) && (player.PossessedAccessory[6] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[7][1])
		player.EquipmentAccessory[1] = descAccessory[7][2]
		player.EquipmentAccessory[2] = descAccessory[7][3]
		player.EquipmentAccessory[3] = descAccessory[7][4]
		log.Println("装備7")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key8)) && (player.PossessedAccessory[7] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[8][1])
		player.EquipmentAccessory[1] = descAccessory[8][2]
		player.EquipmentAccessory[2] = descAccessory[8][3]
		player.EquipmentAccessory[3] = descAccessory[8][4]
		log.Println("装備8")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key9)) && (player.PossessedAccessory[8] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[9][1])
		player.EquipmentAccessory[1] = descAccessory[9][2]
		player.EquipmentAccessory[2] = descAccessory[9][3]
		player.EquipmentAccessory[3] = descAccessory[9][4]
		log.Println("装備9")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.Key0)) && (player.PossessedAccessory[9] == "1") {
		player.EquipmentAccessory[0] = strings.NewReplacer("【", "", "】", "").Replace(descAccessory[10][1])
		player.EquipmentAccessory[1] = descAccessory[10][2]
		player.EquipmentAccessory[2] = descAccessory[10][3]
		player.EquipmentAccessory[3] = descAccessory[10][4]
		log.Println("装備0")
	} else if myState.CurrentBelong == myState.AccessoryBelong && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.AccessoryBelong
		myState.CurrentGS = myState.TownScreen
		log.Println("所持品/武器->GoTo")
	}
	tempOP1, _ := strconv.ParseFloat(loadContent[1][13], 64)
	tempOP2, _ := strconv.ParseFloat(player.EquipmentWeapon[1], 64)
	tempOP3, _ := strconv.ParseFloat(player.EquipmentAccessory[1], 64)
	player.OP = tempOP1 + tempOP2 + tempOP3
	log.Println(&player.OP)

	tempDP1, _ := strconv.ParseFloat(loadContent[1][14], 64)
	tempDP2, _ := strconv.ParseFloat(player.EquipmentArmor[2], 64)
	tempDP3, _ := strconv.ParseFloat(player.EquipmentAccessory[2], 64)
	player.DP = tempDP1 + tempDP2 + tempDP3

	tempAttackTimer1, _ := strconv.ParseFloat(loadContent[1][15], 64)
	tempAttackTimer2, _ := strconv.ParseFloat(player.EquipmentWeapon[3], 64)
	tempAttackTimer3, _ := strconv.ParseFloat(player.EquipmentArmor[3], 64)
	tempAttackTimer4, _ := strconv.ParseFloat(player.EquipmentAccessory[3], 64)
	player.AttackTimer = tempAttackTimer1 + tempAttackTimer2 + tempAttackTimer3 + tempAttackTimer4

	SaveGame(SaveFilePath, 1, player)
	SaveGameAccessory(SaveFilePath, 8, player)
}
