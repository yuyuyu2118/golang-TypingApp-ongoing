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

type ArmorState int

const (
	armorNil ArmorState = iota
	armor1
	armor2
	armor3
	armor4
	armor5
	armor6
	armor7
	armor8
	armor9
	armor10
)

var keyToArmor = map[pixelgl.Button]ArmorState{
	pixelgl.Key1: armor1,
	pixelgl.Key2: armor2,
	pixelgl.Key3: armor3,
	pixelgl.Key4: armor4,
	pixelgl.Key5: armor5,
	pixelgl.Key6: armor6,
	pixelgl.Key7: armor7,
	pixelgl.Key8: armor8,
	pixelgl.Key9: armor9,
	pixelgl.Key0: armor10,
}

var armorSlice = []string{"1. ???", "2. ???", "3. ???", "4. ???", "5. ???", "6. ???", "7. ???", "8. ???", "9. ???", "0. ???"}
var armorNum = []string{"armor0", "armor1", "armor2", "armor3", "armor4", "armor5", "armor6", "armor7", "armor8", "armor9"}
var armorName = []string{"草織りのローブ", "フルーツアーマー", "木の鎧", "ソウルバインドプレート", "スタンプレート", "鉄の鎧", "飛翔のマント", "勇者の鎧", "刃舞の衣", "冥界の鎧"}

var (
	armorPath = "assets/shop/armor.csv"
	descArmor = CsvToSlice(armorPath)
)
var currentarmorState ArmorState

func InitArmor(win *pixelgl.Window, Txt *text.Text, botText string) {
	// メッセージボックスのインスタンス生成
	armorMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 0.4, 0.5)
	// メッセージボックスの表示
	armorMessageBox.DrawMessageBox()

	// 防具用メッセージ作成
	var armorOptions string
	for i, armor := range armorName {
		// arrowIconをループの各イテレーションでリセット
		arrowIcon := ""
		if currentarmorState == ArmorState(i+1) {
			arrowIcon = " ←"
		}
		if event.ArmorPurchaseEventInstance.Armors[i] {
			armorOptions += strconv.Itoa(i+1) + ". " + armor + arrowIcon + "\n"
		} else {
			armorOptions += strconv.Itoa(i+1) + ". ???\n"
		}
	}
	// 10番目の防具の表示を修正
	if event.ArmorPurchaseEventInstance.Armors[9] {
		arrowIcon := "" // ここでもリセット
		if currentarmorState == armor10 {
			arrowIcon = " ←"
		}
		armorOptions = strings.TrimSuffix(armorOptions, "10. ???\n") // 10番目の"????"を削除
		armorOptions += "0. " + armorName[9] + arrowIcon + "\n"
	}

	// メッセージボックスにテキストを表示
	armorMessageBox.DrawMessageTxt("どの防具を購入しますか？\nキーボードに対応する数字を入力してください。\n" + armorOptions + "\n\nBackSpaceキーでタイトルに戻る")

	// キー入力による防具選択処理
	for i := 0; i < len(keyToArmor)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if win.Pressed(key) && event.ArmorPurchaseEventInstance.Armors[i] {
			currentarmorState = ArmorState(i + 1)
			break
		}
	}
	if win.Pressed(pixelgl.Key0) && event.ArmorPurchaseEventInstance.Armors[9] {
		currentarmorState = armor10
	}

	armorDescriptionMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0.4, 0, 1, 0.5)

	armorDescriptionMessageBox.DrawMessageBox()

	// 武器の説明表示処理
	if currentarmorState >= armor1 && currentarmorState <= armor10 {
		if win.Pressed(pixelgl.KeyTab) {
			// SubDescriptionArmor(win, descArmor, int(currentarmorState)-1)
		} else {
			DescriptionArmor(win, descArmor, int(currentarmorState)-1, armorDescriptionMessageBox)
		}
	}
}

func ArmorClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) myState.GameState {
	var tempArmor = ""

	for i := 0; i < len(keyToArmor)-1; i++ {
		key := pixelgl.Button(i + int(pixelgl.Key1))
		if (win.Pressed(key)) && event.ArmorPurchaseEventInstance.Armors[i] && myState.CurrentGS == myState.ArmorShop {
			currentarmorState = ArmorState(i + 1)
			//CreateArmorEvent(descArmor, 0)
			log.Println("防具屋->防具", i+1)
			tempMyMaterialBool = false
			tempMyMaterialName = []string{"", ""}
			tempMyMaterialCount = []int{0, 0}
			break
		}
	}

	if (win.JustPressed(pixelgl.Key0)) && event.ArmorPurchaseEventInstance.Armors[9] && myState.CurrentGS == myState.ArmorShop {
		currentarmorState = armor10
		log.Println("防具屋->防具10")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", ""}
		tempMyMaterialCount = []int{0, 0}
	} else if win.JustPressed(pixelgl.KeyBackspace) && myState.CurrentGS == myState.ArmorShop {
		myState.CurrentGS = myState.TownScreen
		log.Println("防具屋->町")
		tempMyMaterialBool = false
		tempMyMaterialName = []string{"", ""}
		tempMyMaterialCount = []int{0, 0}
	}

	if (win.JustPressed(pixelgl.KeyB)) && player.Gold >= 100 {
		loadContent := SaveFileLoad(SaveFilePath)
		//TODO: お金が足りないときの処理を記述
		for i := 0; i < len(keyToArmor)-1; i++ {
			if currentarmorState == ArmorState(i+1) {
				requiredGold, _ := strconv.Atoi(descArmor[i+1][4])
				belongArmor, _ := strconv.Atoi(loadContent[4][i])
				//log.Println(loadContent)
				log.Println(belongArmor)
				if belongArmor == 0 {
					if player.Gold >= requiredGold {
						log.Println(descArmor[i+1][4], "買える", player.Gold)
						createOk := CreateArmorEvent(descArmor, i)
						if createOk {
							player.Gold -= requiredGold
							tempArmor = "armor" + strconv.Itoa(i+1)
							tempMyMaterialBool = false
							tempMyMaterialName = []string{"", ""}
							tempMyMaterialCount = []int{0, 0}
						}
					} else {
						log.Println(descArmor[i+1][4], "お金が足りない", player.Gold)
					}
				} else {
					log.Println("すでに持っている")
					break
				}
			}
		}
		if currentarmorState == armor10 {
			requiredGold, _ := strconv.Atoi(descArmor[10][4])
			if player.Gold >= requiredGold {
				log.Println(descArmor[10][4], "買える", player.Gold)
			} else {
				log.Println(descArmor[10][4], "お金が足りない", player.Gold)
			}
			log.Println(descArmor[10][4])
			tempArmor = "armor" + strconv.Itoa(10)
		}

		if tempArmor != "" {
			SaveArmorPurchaseEvent(SaveFilePath, 4, tempArmor, player)
			SaveGame(SaveFilePath, 1, player)
		}
	}
	return myState.CurrentGS
}

func DescriptionArmor(win *pixelgl.Window, descArmor [][]string, num int, msgBox *myPos.MessageBox) {
	loadContent = SaveFileLoad(SaveFilePath)
	temp, _ := CountMyItems(SaveFilePathItems)

	//TODO: Tabを押している間は強化素材等の情報を表示する
	num++

	var armorDescriptionOptions string
	// 防具名
	armorDescriptionOptions += descArmor[0][1] + ": " + descArmor[num][1] + "\n"

	armorDescriptionOptions += "カラー: " + descArmor[num][17] + "\n"

	armorDescriptionOptions += descArmor[0][2] + ": " + descArmor[num][2] + "\n"

	armorDescriptionOptions += descArmor[0][3] + ": " + descArmor[num][3] + "\n"

	armorDescriptionOptions += descArmor[0][4] + ": " + descArmor[num][4] + "S\n"

	armorDescriptionOptions += "素材: " + descArmor[num][5] + descArmor[num][6] + "個, " + descArmor[num][7] + descArmor[num][8] + "個\n"

	if !tempMyMaterialBool {
		tempMyMaterialName[0] = descArmor[num][5]
		tempMyMaterialName[1] = descArmor[num][7]
		for name, count := range temp {
			if name == descArmor[num][5] {
				tempMyMaterialName[0] = name
				tempMyMaterialCount[0] = count
			} else if name == descArmor[num][7] {
				tempMyMaterialName[1] = name
				tempMyMaterialCount[1] = count
			} else if name == descArmor[num][9] {
				tempMyMaterialName[2] = name
				tempMyMaterialCount[2] = count
			}
		}
		tempMyMaterialBool = true
	}

	armorDescriptionOptions += "所持: " + tempMyMaterialName[0] + strconv.Itoa(tempMyMaterialCount[0]) + "個, " + tempMyMaterialName[1] + strconv.Itoa(tempMyMaterialCount[1]) + "個\n"

	armorDescriptionOptions += "説明: " + descArmor[num][11] + "\n"

	armorDescriptionOptions += "　　  " + descArmor[num][12] + "\n"

	armorDescriptionOptions += "特殊能力: " + descArmor[num][14] + "\n"

	armorDescriptionOptions += "　　　　  " + descArmor[num][15] + "\n"

	armorDescriptionOptions += "　　　　  " + descArmor[num][16] + "\n\n"

	armorDescriptionOptions += "B. 作ってもらう\n"

	if loadContent[3][num-1] == strconv.Itoa(1) {
		armorDescriptionOptions += "作成済み\n"
	} else {
		armorDescriptionOptions += "\n"
	}

	msgBox.DrawMessageTxt(armorDescriptionOptions)
}

func CreateArmorEvent(descArmor [][]string, num int) bool {
	//TODO: 素材が足りるかどうかの判定実装中
	num++
	tempSlice, _ := CountMyItems(SaveFilePathItems)
	var tempBool = []bool{false, false, false}

	for name, count := range tempSlice {
		if name == descArmor[num][5] {
			tempCount, _ := strconv.Atoi(descArmor[num][6])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[0] = true
			}
		} else if (descArmor)[num][5] == "" {
			//log.Println("なし")
			tempBool[0] = true
		}
		if name == descArmor[num][7] {
			tempCount, _ := strconv.Atoi(descArmor[num][8])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[1] = true
			}
		} else if (descArmor)[num][7] == "" {
			//log.Println("なし")
			tempBool[1] = true
		}
		if name == descArmor[num][9] {
			tempCount, _ := strconv.Atoi(descArmor[num][10])
			if count >= tempCount {
				log.Println(name, count, tempCount, "足りてます")
				tempBool[2] = true
			}
		} else if (descArmor)[num][9] == "" {
			//log.Println("なし")
			tempBool[2] = true
		}
	}
	if tempBool[0] && tempBool[1] && tempBool[2] {
		log.Println("素材が全部あります")
		for name, _ := range tempSlice {
			if name == descArmor[num][5] {
				tempCount, _ := strconv.Atoi(descArmor[num][6])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][7] {
				tempCount, _ := strconv.Atoi(descArmor[num][8])
				tempSlice[name] -= tempCount
			}
			if name == descArmor[num][9] {
				tempCount, _ := strconv.Atoi(descArmor[num][10])
				tempSlice[name] -= tempCount
			}
		}
		log.Println(tempSlice)
		SaveGameLostItems(SaveFilePathItems, tempSlice)
		log.Println("素材を消費して防具を作成しました。")
		return true
	} else {
		log.Println("素材が一部足りません")
		return false
	}
}

func InitArmorBelongScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus) {
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	botText := "持ち物/防具"
	InitArmorBelong(win, Txt, botText, player)
}

func InitArmorBelong(win *pixelgl.Window, Txt *text.Text, botText string, player *myPlayer.PlayerStatus) {
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
	elements := loadContent[4]

	for i, val := range elements {
		num, err := strconv.Atoi(val)
		if err == nil {
			armorKey := fmt.Sprintf("armor%d", i)
			counts[armorKey] = num
		}
	}

	for i, value := range armorName {
		if counts["armor"+strconv.Itoa(i)] != 0 {
			//tempInt := counts["armor"+strconv.Itoa(i)]
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

		if player.EquipmentArmor[0] == armorName[i] {
			Txt.Clear()
			fmt.Fprintln(Txt, "E. ")
			txtPos = pixel.V(xOffSet-40, yOffSet)
			tempPosition = pixel.IM.Moved(txtPos)
			Txt.Draw(win, tempPosition)
		}
	}
	equipmentSlice = equipmentSlice[:0]
}

func ArmorBelongClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *myPlayer.PlayerStatus) {
	loadContent := SaveFileLoad(SaveFilePath)

	if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key1)) && (player.PossessedArmor[0] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[1][1])
		player.EquipmentArmor[2] = descArmor[1][2]
		player.EquipmentArmor[3] = descArmor[1][3]
		log.Println("装備1")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key2)) && (player.PossessedArmor[1] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[2][1])
		player.EquipmentArmor[2] = descArmor[2][2]
		player.EquipmentArmor[3] = descArmor[2][3]
		log.Println("装備2")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key3)) && (player.PossessedArmor[2] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[3][1])
		player.EquipmentArmor[2] = descArmor[3][2]
		player.EquipmentArmor[3] = descArmor[3][3]
		log.Println("装備3")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key4)) && (player.PossessedArmor[3] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[4][1])
		player.EquipmentArmor[2] = descArmor[4][2]
		player.EquipmentArmor[3] = descArmor[4][3]
		log.Println("装備4")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key5)) && (player.PossessedArmor[4] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[5][1])
		player.EquipmentArmor[2] = descArmor[5][2]
		player.EquipmentArmor[3] = descArmor[5][3]
		log.Println("装備5")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key6)) && (player.PossessedArmor[5] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[6][1])
		player.EquipmentArmor[2] = descArmor[6][2]
		player.EquipmentArmor[3] = descArmor[6][3]
		log.Println("装備6")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key7)) && (player.PossessedArmor[6] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[7][1])
		player.EquipmentArmor[2] = descArmor[7][2]
		player.EquipmentArmor[3] = descArmor[7][3]
		log.Println("装備7")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key8)) && (player.PossessedArmor[7] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[8][1])
		player.EquipmentArmor[2] = descArmor[8][2]
		player.EquipmentArmor[3] = descArmor[8][3]
		log.Println("装備8")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key9)) && (player.PossessedArmor[8] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[9][1])
		player.EquipmentArmor[2] = descArmor[9][2]
		player.EquipmentArmor[3] = descArmor[9][3]
		log.Println("装備9")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.Key0)) && (player.PossessedArmor[9] == "1") {
		player.EquipmentArmor[0] = strings.NewReplacer("【", "", "】", "").Replace(descArmor[10][1])
		player.EquipmentArmor[2] = descArmor[10][2]
		player.EquipmentArmor[3] = descArmor[10][3]
		log.Println("装備0")
	} else if myState.CurrentBelong == myState.ArmorBelong && (win.JustPressed(pixelgl.KeyBackspace)) {
		myState.CurrentBelong = myState.ArmorBelong
		myState.CurrentGS = myState.TownScreen
		log.Println("所持品/武器->GoTo")
	}
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
	SaveGameArmor(SaveFilePath, 7, player)
}
