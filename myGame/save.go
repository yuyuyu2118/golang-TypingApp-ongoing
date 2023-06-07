package myGame

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/yuyuyu2118/typingGo/player"
)

var SaveFilePath = "player\\save\\save.csv"
var SaveFilePathItems = "player\\save\\saveItems.csv"

var loadContent = SaveFileLoad(SaveFilePath)

func SaveGame(saveFilePath string, saveNum int, player *player.PlayerStatus) {
	SaveFileCheck(saveFilePath)
	//saveContent := "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
	Name := player.Name
	MaxHP := strconv.FormatFloat(player.MaxHP, 'f', -1, 64)
	HP := strconv.FormatFloat(player.HP, 'f', -1, 64)
	OP := strconv.FormatFloat(player.OP, 'f', -1, 64)
	DP := strconv.FormatFloat(player.DP, 'f', -1, 64)
	MaxSP := strconv.FormatFloat(player.MaxSP, 'f', -1, 64)
	SP := strconv.FormatFloat(player.SP, 'f', -1, 64)
	BaseSP := strconv.FormatFloat(player.BaseSP, 'f', -1, 64)
	Gold := strconv.Itoa(player.Gold)
	Job := player.Job
	AP := strconv.Itoa(player.AP)
	Language := player.Language
	AttackTime := strconv.FormatFloat(player.AttackTimer, 'f', -1, 64)
	BaseOP := strconv.FormatFloat(3.0, 'f', -1, 64)
	BaseDP := strconv.FormatFloat(0.0, 'f', -1, 64)
	BaseAttackTime := strconv.FormatFloat(4.0, 'f', -1, 64)
	saveContent := Name + "," + MaxHP + "," + HP + "," + OP + "," + DP + "," + MaxSP + "," + SP + "," + BaseSP + "," + Gold + "," + Job + "," + AP + "," + Language + "," + AttackTime + "," + BaseOP + "," + BaseDP + "," + BaseAttackTime

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}

	// 改行文字で分割して行ごとのスライスに変換
	lines := strings.Split(string(content), "\n")

	// 行番号が有効な範囲かチェック
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}

	// 指定された行を上書き
	lines[saveNum] = saveContent

	// 更新後の内容を保存ファイルに書き込む
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)

}

func SaveFileLoad(saveFilePath string) [][]string {
	SaveFileCheck(saveFilePath)
	return CsvToSlice(saveFilePath)
}

func SaveFileCheck(saveFilePath string) {
	tempInitText := []string{"Name,MaxHP,HP,OP,DP,MaxSP,SP,BaseSP,Gold,Job,AP,language,AttackTimer,BaseOP,BaseDP,BaseAttackTimer",
		"NoName,30,30,3,0,50,0,2,0,No Job,0,Japanese,4.0,3.0,0.0,4.0",
		"0,0,0,0,0,0,0,0,0,0,,,,,,",
		"0,0,0,0,0,0,0,0,0,0,,,,,,",
		"0,0,0,0,0,0,0,0,0,0,,,,,,",
		"0,0,0,0,0,0,0,0,0,0,,,,,,",
		"WeaponName,,,,,,,,,,,,,,,",
		"ArmorName,,,,,,,,,,,,,,,",
		"AccessoryName,,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
		",,,,,,,,,,,,,,,",
	}
	initializeText := strings.Join(tempInitText, "\n")
	//initializeText := "Name,MaxHP,HP,OP,DP,MaxSP,SP,BaseSP,Gold,Job,AP,\nNoName,30,30,3,1,50,0,2,0,No Job,0,\nWeaponName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,\nArmorName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,\nAccessoryName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,"
	fileInfo, err := os.Stat(saveFilePath)
	if err != nil {
		// ファイルが存在しない場合、初回呼び出しとして初期化テキストを出力
		if os.IsNotExist(err) {
			err := ioutil.WriteFile(saveFilePath, []byte(initializeText), 0644)
			if err != nil {
				fmt.Println("保存ファイルの作成に失敗しました:", err)
				return
			}
			fmt.Println("保存ファイルを作成しました。初期化テキストを出力しました。")
			return
		}

		fmt.Println("保存ファイルの情報取得に失敗しました:", err)
		return
	}

	// 保存ファイルが空であるかをチェック
	if fileInfo.Size() == 0 {
		err := ioutil.WriteFile(saveFilePath, []byte(initializeText), 0644)
		if err != nil {
			fmt.Println("保存ファイルの初期化に失敗しました:", err)
			return
		}
		fmt.Println("保存ファイルを初期化しました。初期化テキストを出力しました。")
		return
	}
}

func LoadSliceToString(content []string) string {
	quotedValues := make([]string, len(content))
	for i, v := range content {
		quotedValues[i] = `"` + v + `"`
	}
	return strings.Join(quotedValues, ",")
}

// weaponが買えるようになったら有効化
func SaveDefeatedEnemyEvent(saveFilePath string, saveNum int, defeatEnemy string) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if defeatEnemy == "Slime" {
		tempInt, _ = strconv.Atoi(loadContent[2][0])
		loadContent[2][0] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Bird" {
		tempInt, _ = strconv.Atoi(loadContent[2][1])
		loadContent[2][1] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Plant" {
		tempInt, _ = strconv.Atoi(loadContent[2][2])
		loadContent[2][2] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Kobold" {
		tempInt, _ = strconv.Atoi(loadContent[2][3])
		loadContent[2][3] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Zombie" {
		tempInt, _ = strconv.Atoi(loadContent[2][4])
		loadContent[2][4] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Fairy" {
		tempInt, _ = strconv.Atoi(loadContent[2][5])
		loadContent[2][5] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Skull" {
		tempInt, _ = strconv.Atoi(loadContent[2][6])
		loadContent[2][6] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Wizard" {
		tempInt, _ = strconv.Atoi(loadContent[2][7])
		loadContent[2][7] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Solidier" {
		tempInt, _ = strconv.Atoi(loadContent[2][8])
		loadContent[2][8] = strconv.Itoa(tempInt + 1)
	} else if defeatEnemy == "Dragon" {
		tempInt, _ = strconv.Atoi(loadContent[2][9])
		loadContent[2][9] = strconv.Itoa(tempInt + 1)
	}
	saveContent := strings.Join(loadContent[saveNum], ",")

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}
	lines[saveNum] = saveContent

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveWeaponPurchaseEvent(saveFilePath string, saveNum int, purchaseWeapon string, player *player.PlayerStatus) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if purchaseWeapon == "weapon1" {
		tempInt, _ = strconv.Atoi(loadContent[3][0])
		loadContent[3][0] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[0] = loadContent[3][0]
	} else if purchaseWeapon == "weapon2" {
		tempInt, _ = strconv.Atoi(loadContent[3][1])
		loadContent[3][1] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[1] = loadContent[3][1]
	} else if purchaseWeapon == "weapon3" {
		tempInt, _ = strconv.Atoi(loadContent[3][2])
		loadContent[3][2] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[2] = loadContent[3][2]
	} else if purchaseWeapon == "weapon4" {
		tempInt, _ = strconv.Atoi(loadContent[3][3])
		loadContent[3][3] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[3] = loadContent[3][3]
	} else if purchaseWeapon == "weapon5" {
		tempInt, _ = strconv.Atoi(loadContent[3][4])
		loadContent[3][4] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[4] = loadContent[3][4]
	} else if purchaseWeapon == "weapon6" {
		tempInt, _ = strconv.Atoi(loadContent[3][5])
		loadContent[3][5] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[5] = loadContent[3][5]
	} else if purchaseWeapon == "weapon7" {
		tempInt, _ = strconv.Atoi(loadContent[3][6])
		loadContent[3][6] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[6] = loadContent[3][6]
	} else if purchaseWeapon == "weapon8" {
		tempInt, _ = strconv.Atoi(loadContent[3][7])
		loadContent[3][7] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[7] = loadContent[3][7]
	} else if purchaseWeapon == "weapon9" {
		tempInt, _ = strconv.Atoi(loadContent[3][8])
		loadContent[3][8] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[8] = loadContent[3][8]
	} else if purchaseWeapon == "weapon0" {
		tempInt, _ = strconv.Atoi(loadContent[3][9])
		loadContent[3][9] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[9] = loadContent[3][9]
	}
	saveContent := strings.Join(loadContent[saveNum], ",")

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}
	lines[saveNum] = saveContent

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveArmorPurchaseEvent(saveFilePath string, saveNum int, purchaseArmor string, player *player.PlayerStatus) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if purchaseArmor == "armor1" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][0])
		loadContent[saveNum][0] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[0] = loadContent[saveNum][0]
	} else if purchaseArmor == "armor2" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][1])
		loadContent[saveNum][1] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[1] = loadContent[saveNum][1]
	} else if purchaseArmor == "armor3" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][2])
		loadContent[saveNum][2] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[2] = loadContent[saveNum][2]
	} else if purchaseArmor == "armor4" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][3])
		loadContent[saveNum][3] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[3] = loadContent[saveNum][3]
	} else if purchaseArmor == "armor5" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][4])
		loadContent[saveNum][4] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[4] = loadContent[saveNum][4]
	} else if purchaseArmor == "armor6" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][5])
		loadContent[saveNum][5] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[5] = loadContent[saveNum][5]
	} else if purchaseArmor == "armor7" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][6])
		loadContent[saveNum][6] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[6] = loadContent[saveNum][6]
	} else if purchaseArmor == "armor8" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][7])
		loadContent[saveNum][7] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[7] = loadContent[saveNum][7]
	} else if purchaseArmor == "armor9" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][8])
		loadContent[saveNum][8] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[8] = loadContent[saveNum][8]
	} else if purchaseArmor == "armor0" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][9])
		loadContent[saveNum][9] = strconv.Itoa(tempInt + 1)
		player.PossessedArmor[9] = loadContent[saveNum][9]
	}
	saveContent := strings.Join(loadContent[saveNum], ",")

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}
	lines[saveNum] = saveContent

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveAccessoryPurchaseEvent(saveFilePath string, saveNum int, purchaseAccessory string, player *player.PlayerStatus) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if purchaseAccessory == "accessory1" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][0])
		loadContent[saveNum][0] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[0] = loadContent[saveNum][0]
	} else if purchaseAccessory == "accessory2" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][1])
		loadContent[saveNum][1] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[1] = loadContent[saveNum][1]
	} else if purchaseAccessory == "accessory3" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][2])
		loadContent[saveNum][2] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[2] = loadContent[saveNum][2]
	} else if purchaseAccessory == "accessory4" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][3])
		loadContent[saveNum][3] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[3] = loadContent[saveNum][3]
	} else if purchaseAccessory == "accessory5" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][4])
		loadContent[saveNum][4] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[4] = loadContent[saveNum][4]
	} else if purchaseAccessory == "accessory6" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][5])
		loadContent[saveNum][5] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[5] = loadContent[saveNum][5]
	} else if purchaseAccessory == "accessory7" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][6])
		loadContent[saveNum][6] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[6] = loadContent[saveNum][6]
	} else if purchaseAccessory == "accessory8" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][7])
		loadContent[saveNum][7] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[7] = loadContent[saveNum][7]
	} else if purchaseAccessory == "accessory9" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][8])
		loadContent[saveNum][8] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[8] = loadContent[saveNum][8]
	} else if purchaseAccessory == "accessory0" {
		tempInt, _ = strconv.Atoi(loadContent[saveNum][9])
		loadContent[saveNum][9] = strconv.Itoa(tempInt + 1)
		player.PossessedAccessory[9] = loadContent[saveNum][9]
	}
	saveContent := strings.Join(loadContent[saveNum], ",")

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}
	lines[saveNum] = saveContent

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveWeaponEnhancementEvent(saveFilePath string, saveNum int, enhancementWeapon string, player *player.PlayerStatus) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if enhancementWeapon == "weapon1" {
		tempInt, _ = strconv.Atoi(loadContent[9][0])
		loadContent[9][0] = strconv.Itoa(tempInt + 1)
		player.WeaponEnhancement[0] = loadContent[9][0]
	} else if enhancementWeapon == "weapon2" {
		tempInt, _ = strconv.Atoi(loadContent[3][1])
		loadContent[3][1] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[1] = loadContent[3][1]
	} else if enhancementWeapon == "weapon3" {
		tempInt, _ = strconv.Atoi(loadContent[3][2])
		loadContent[3][2] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[2] = loadContent[3][2]
	} else if enhancementWeapon == "weapon4" {
		tempInt, _ = strconv.Atoi(loadContent[3][3])
		loadContent[3][3] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[3] = loadContent[3][3]
	} else if enhancementWeapon == "weapon5" {
		tempInt, _ = strconv.Atoi(loadContent[3][4])
		loadContent[3][4] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[4] = loadContent[3][4]
	} else if enhancementWeapon == "weapon6" {
		tempInt, _ = strconv.Atoi(loadContent[3][5])
		loadContent[3][5] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[5] = loadContent[3][5]
	} else if enhancementWeapon == "weapon7" {
		tempInt, _ = strconv.Atoi(loadContent[3][6])
		loadContent[3][6] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[6] = loadContent[3][6]
	} else if enhancementWeapon == "weapon8" {
		tempInt, _ = strconv.Atoi(loadContent[3][7])
		loadContent[3][7] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[7] = loadContent[3][7]
	} else if enhancementWeapon == "weapon9" {
		tempInt, _ = strconv.Atoi(loadContent[3][8])
		loadContent[3][8] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[8] = loadContent[3][8]
	} else if enhancementWeapon == "weapon0" {
		tempInt, _ = strconv.Atoi(loadContent[3][9])
		loadContent[3][9] = strconv.Itoa(tempInt + 1)
		player.PossessedWeapon[9] = loadContent[3][9]
	}
	saveContent := strings.Join(loadContent[saveNum], ",")

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}
	lines[saveNum] = saveContent

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

// func SaveWeaponSellEvent(saveFilePath string, saveNum int, sellWeapon string, player *player.PlayerStatus) {
// 	var tempInt int
// 	loadContent := SaveFileLoad(saveFilePath)
// 	if sellWeapon == "weapon1" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][0])
// 		loadContent[3][0] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[0] = loadContent[3][0]
// 	} else if sellWeapon == "weapon2" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][1])
// 		loadContent[3][1] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[1] = loadContent[3][1]
// 	} else if sellWeapon == "weapon3" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][2])
// 		loadContent[3][2] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[2] = loadContent[3][2]
// 	} else if sellWeapon == "weapon4" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][3])
// 		loadContent[3][3] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[3] = loadContent[3][3]
// 	} else if sellWeapon == "weapon5" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][4])
// 		loadContent[3][4] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[4] = loadContent[3][4]
// 	} else if sellWeapon == "weapon6" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][5])
// 		loadContent[3][5] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[5] = loadContent[3][5]
// 	} else if sellWeapon == "weapon7" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][6])
// 		loadContent[3][6] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[6] = loadContent[3][6]
// 	} else if sellWeapon == "weapon8" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][7])
// 		loadContent[3][7] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[7] = loadContent[3][7]
// 	} else if sellWeapon == "weapon9" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][8])
// 		loadContent[3][8] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[8] = loadContent[3][8]
// 	} else if sellWeapon == "weapon0" {
// 		tempInt, _ = strconv.Atoi(loadContent[3][9])
// 		loadContent[3][9] = strconv.Itoa(tempInt - 1)
// 		player.PossessedWeapon[9] = loadContent[3][9]
// 	}
// 	saveContent := strings.Join(loadContent[saveNum], ",")

// 	content, err := ioutil.ReadFile(saveFilePath)
// 	if err != nil {
// 		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
// 		return
// 	}
// 	lines := strings.Split(string(content), "\n")
// 	if saveNum < 0 || saveNum >= len(lines) {
// 		fmt.Println("指定された行番号が範囲外です。")
// 		return
// 	}
// 	lines[saveNum] = saveContent

// 	output := strings.Join(lines, "\n")
// 	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
// 	if err != nil {
// 		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
// 		return
// 	}

// 	fmt.Println("保存ファイルを更新しました。")
// }

// TODO: アイテムのセーブ実装中
func SaveGameItems(SaveFilePathItems string, gainItem []string) error {
	// CSVファイルからデータを読み込む
	var records [][]string

	file, err := os.Open(SaveFilePathItems)
	if err != nil && os.IsNotExist(err) {
		// ファイルが存在しない場合は、新しいファイルを作成する
		file, err = os.Create(SaveFilePathItems)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		// ファイルが存在する場合は、データを読み込む
		defer file.Close()

		reader := csv.NewReader(file)
		records, err = reader.ReadAll()
		if err != nil {
			return err
		}
	}

	// アイテムをカウントするマップを作成する
	itemCountMap := make(map[string]int)
	for _, record := range records {
		if len(record) == 2 {
			count, err := strconv.Atoi(record[1])
			if err == nil {
				itemCountMap[record[0]] = count
			}
		}
	}

	// アイテムを追加または更新する
	for _, item := range gainItem {
		itemCountMap[item]++
	}

	// カウント数でアイテムを降順にソートする
	sortedItems := make([][2]string, 0, len(itemCountMap))
	for item, count := range itemCountMap {
		sortedItems = append(sortedItems, [2]string{item, strconv.Itoa(count)})
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		countI, _ := strconv.Atoi(sortedItems[i][1])
		countJ, _ := strconv.Atoi(sortedItems[j][1])
		return countI > countJ
	})

	// 新しいCSVフォーマットに変換する
	var newRecords [][]string
	for _, record := range sortedItems {
		newRecord := []string{record[0], record[1]}
		newRecords = append(newRecords, newRecord)
	}

	// 新しいCSVファイルを書き出す
	newFile, err := os.Create(SaveFilePathItems)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.WriteAll(newRecords)
	writer.Flush()

	return nil
}

func SaveFileItemsCheck(saveFilePathItems string) {
	initializeText := "木の枝,0"
	fileInfo, err := os.Stat(saveFilePathItems)
	if err != nil {
		// ファイルが存在しない場合、初回呼び出しとして初期化テキストを出力
		if os.IsNotExist(err) {
			err := ioutil.WriteFile(saveFilePathItems, []byte(initializeText), 0644)
			if err != nil {
				fmt.Println("保存ファイルの作成に失敗しました:", err)
				return
			}
			fmt.Println("保存ファイルを作成しました。初期化テキストを出力しました。")
			return
		}

		fmt.Println("保存ファイルの情報取得に失敗しました:", err)
		return
	}

	// 保存ファイルが空であるかをチェック
	if fileInfo.Size() == 0 {
		err := ioutil.WriteFile(saveFilePathItems, []byte(initializeText), 0644)
		if err != nil {
			fmt.Println("保存ファイルの初期化に失敗しました:", err)
			return
		}
		fmt.Println("保存ファイルを初期化しました。初期化テキストを出力しました。")
		return
	}

	SaveFileItemsCheckUpdateV101(saveFilePathItems)

}

func SaveFileItemsLoad(saveFilePathItems string) [][]string {
	SaveFileItemsCheck(saveFilePathItems)
	return CsvToSlice(saveFilePathItems)
}

func saveFileItemCount(records []string) [][]string {
	// アイテムをカウントするマップを作成する
	itemCountMap := make(map[string]int)
	for _, item := range records {
		itemCountMap[item]++
	}
	var newRecords [][]string
	for item, count := range itemCountMap {
		newRecord := []string{item, fmt.Sprintf("%d", count)}
		newRecords = append(newRecords, newRecord)
	}
	log.Println(newRecords)
	return newRecords
}

func GetMyItems(SaveFilePathItems string) ([]string, error) {
	// CSVファイルからデータを読み込む
	var records [][]string

	file, err := os.Open(SaveFilePathItems)
	if err != nil && os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		defer file.Close()
		reader := csv.NewReader(file)
		records, err = reader.ReadAll()
		if err != nil {
			return nil, err
		}
	}

	// アイテムをカウントするマップを作成する
	itemCountMap := make(map[string]int)
	for _, record := range records {
		if len(record) == 2 {
			count, err := strconv.Atoi(record[1])
			if err == nil {
				itemCountMap[record[0]] = count
			}
		}
	}
	sortedItems := make([][2]string, 0, len(itemCountMap))
	for item, count := range itemCountMap {
		sortedItems = append(sortedItems, [2]string{item, strconv.Itoa(count)})
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		countI, _ := strconv.Atoi(sortedItems[i][1])
		countJ, _ := strconv.Atoi(sortedItems[j][1])
		return countI > countJ
	})
	combinedItems := make([]string, 0, len(sortedItems))
	for _, value := range sortedItems {
		combined := fmt.Sprintf("%s: %s個", value[0], value[1])
		combinedItems = append(combinedItems, combined)
	}
	return combinedItems, err
}

func CountMyItems(SaveFilePathItems string) (map[string]int, error) {
	// CSVファイルからデータを読み込む
	var records [][]string

	file, err := os.Open(SaveFilePathItems)
	if err != nil && os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		defer file.Close()
		reader := csv.NewReader(file)
		records, err = reader.ReadAll()
		if err != nil {
			return nil, err
		}
	}

	// アイテムをカウントするマップを作成する
	itemCountMap := make(map[string]int)
	for _, record := range records {
		if len(record) == 2 {
			count, err := strconv.Atoi(record[1])
			if err == nil {
				itemCountMap[record[0]] = count
			}
		}
	}
	return itemCountMap, err
}

func SaveGameLostItems(SaveFilePathItems string, tempSlice map[string]int) error {
	// CSVファイルからデータを読み込む
	var records [][]string

	file, err := os.Open(SaveFilePathItems)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		return err
	}

	// アイテムをカウントするマップを作成する
	itemCountMap := make(map[string]int)
	for _, record := range records {
		if len(record) == 2 {
			count, err := strconv.Atoi(record[1])
			if err == nil {
				itemCountMap[record[0]] = count
			}
		}
	}

	// アイテムを追加または更新する
	for item, count := range tempSlice {
		for i := 0; i < 3; i++ {
			itemCountMap[item] = count
		}
	}

	// カウント数でアイテムを降順にソートする
	sortedItems := make([][2]string, 0, len(itemCountMap))
	for item, count := range itemCountMap {
		sortedItems = append(sortedItems, [2]string{item, strconv.Itoa(count)})
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		countI, _ := strconv.Atoi(sortedItems[i][1])
		countJ, _ := strconv.Atoi(sortedItems[j][1])
		return countI > countJ
	})

	// 新しいCSVフォーマットに変換する
	var newRecords [][]string
	for _, record := range sortedItems {
		newRecord := []string{record[0], record[1]}
		newRecords = append(newRecords, newRecord)
	}

	// 新しいCSVファイルを書き出す
	newFile, err := os.Create(SaveFilePathItems)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.WriteAll(newRecords)
	writer.Flush()

	return nil
}

func SaveGameWeapon(saveFilePath string, saveNum int, player *player.PlayerStatus) {
	SaveFileCheck(saveFilePath)
	//saveContent := "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
	Name := player.EquipmentWeapon[0]
	OP := player.EquipmentWeapon[1]
	DP := player.EquipmentWeapon[2]
	AttackTimer := player.EquipmentWeapon[3]
	saveContent := Name + "," + OP + "," + DP + "," + AttackTimer + ",,,,,,,,,,,,"

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}

	// 改行文字で分割して行ごとのスライスに変換
	lines := strings.Split(string(content), "\n")

	// 行番号が有効な範囲かチェック
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}

	// 指定された行を上書き
	lines[saveNum] = saveContent

	// 更新後の内容を保存ファイルに書き込む
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveGameArmor(saveFilePath string, saveNum int, player *player.PlayerStatus) {
	SaveFileCheck(saveFilePath)
	Name := player.EquipmentArmor[0]
	OP := player.EquipmentArmor[1]
	DP := player.EquipmentArmor[2]
	AttackTimer := player.EquipmentArmor[3]
	saveContent := Name + "," + OP + "," + DP + "," + AttackTimer + ",,,,,,,,,,,,"

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}

	// 改行文字で分割して行ごとのスライスに変換
	lines := strings.Split(string(content), "\n")

	// 行番号が有効な範囲かチェック
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}

	// 指定された行を上書き
	lines[saveNum] = saveContent

	// 更新後の内容を保存ファイルに書き込む
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveGameAccessory(saveFilePath string, saveNum int, player *player.PlayerStatus) {
	SaveFileCheck(saveFilePath)
	Name := player.EquipmentAccessory[0]
	OP := player.EquipmentAccessory[1]
	DP := player.EquipmentAccessory[2]
	AttackTimer := player.EquipmentAccessory[3]
	saveContent := Name + "," + OP + "," + DP + "," + AttackTimer + ",,,,,,,,,,,,"

	content, err := ioutil.ReadFile(saveFilePath)
	if err != nil {
		fmt.Println("保存ファイルの読み込みに失敗しました:", err)
		return
	}

	// 改行文字で分割して行ごとのスライスに変換
	lines := strings.Split(string(content), "\n")

	// 行番号が有効な範囲かチェック
	if saveNum < 0 || saveNum >= len(lines) {
		fmt.Println("指定された行番号が範囲外です。")
		return
	}

	// 指定された行を上書き
	lines[saveNum] = saveContent

	// 更新後の内容を保存ファイルに書き込む
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(saveFilePath, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
	time.Sleep(10 * time.Millisecond)
}

func SaveFileItemsCheckUpdateV101(saveFilePathItems string) {
	// Read CSV contents
	file, err := os.Open(saveFilePathItems)
	if err != nil {
		fmt.Println("ファイルを開くことができませんでした:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	itemMap := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		switch splitLine[0] {
		case "糸":
			splitLine[0] = "木の枝"
		case "針金":
			splitLine[0] = "ただの砥石"
		case "鋭い刃物", "新鮮な果物":
			splitLine[0] = "錆びた鉄鉱石"
		case "小さな砥石":
			splitLine[0] = "ただの砥石"
		case "皮革", "木工用接着剤":
			splitLine[0] = "木材"
		case "魔力の鉱石", "呪文の書":
			splitLine[0] = "魔鉱石"
		case "生命のエキス":
			splitLine[0] = "ただの砥石"
		case "堅固な木材", "鉄の塊":
			splitLine[0] = "鉄鉱石"
		case "衝撃吸収材":
			splitLine[0] = "鋭い砥石"
		case "鍛治のハンマー":
			splitLine[0] = "精錬された鉄"
		case "風の石":
			splitLine[0] = "軽量な合金"
		case "機敏な羽根":
			splitLine[0] = "鋭い砥石"
		case "堅固な鋼鉄", "竜の鱗":
			splitLine[0] = "伝説の鋼"
		case "魔力の宝珠":
			splitLine[0] = "精密砥石"
		case "神秘の石":
			splitLine[0] = "古代の鉄鉱石"
		case "優れた磨き粉":
			splitLine[0] = "精密砥石"
		case "鋭利な刃物", "魔力の結晶":
			splitLine[0] = "暗黒の合金"

		case "金の糸":
			splitLine[0] = "小さな宝石"
		case "守護の石":
			splitLine[0] = "宝石"

		case "草":
			splitLine[0] = "木の枝"
		case "リンゴ", "オレンジ":
			splitLine[0] = "ロープ"
		case "バナナ":
			splitLine[0] = "ロープ"
		case "魂の結晶", "鉄鉱石の欠片":
			splitLine[0] = "魔鉱石"
		case "ルーン石":
			splitLine[0] = "頑丈なロープ"
		case "スタンプカード":
			splitLine[0] = "頑丈なロープ"
		case "軽い鉄鉱石":
			splitLine[0] = "精錬された鉄"
		case "羽根":
			splitLine[0] = "軽量な合金"
		case "魔法の鉄鉱石", "巨大な竜の鱗":
			splitLine[0] = "伝説の鋼"
		case "勇者の証":
			splitLine[0] = "最高のロープ"
		case "幻影の絹糸":
			splitLine[0] = "古代の鉄鉱石"
		case "冥界の魂鉱石", "呪詛の織物":
			splitLine[0] = "暗黒の合金"
		case "魂魄":
			splitLine[0] = "最高のロープ"

		case "魔法の粉":
			splitLine[0] = "木の枝"
		case "緑の宝石":
			splitLine[0] = "小さな宝石"
		case "果物の種", "銀色の鎖":
			splitLine[0] = "錆びた鉄鉱石"
		case "太陽の水滴":
			splitLine[0] = "小さな宝石"
		case "癒しの草", "聖なる石":
			splitLine[0] = "木材"
		case "生命のエッセンス", "幻の鋼":
			splitLine[0] = "魔鉱石"
		case "魔鉱石":
			splitLine[0] = "小さな宝石"
		case "雷の結晶", "エネルギーの粉":
			splitLine[0] = "鉄鉱石"
		case "導電金属":
			splitLine[0] = "宝石"
		case "炎の粉":
			splitLine[0] = "精錬された鉄"
		case "鍛冶のハンマ-":
			splitLine[0] = "宝石"
		case "風の羽根", "弾力のある皮":
			splitLine[0] = "軽量な合金"
		case "速度の石":
			splitLine[0] = "宝石"
		case "勇気の石", "神秘の鉱石":
			splitLine[0] = "伝説の鋼"
		case "金属片":
			splitLine[0] = "古代の鉄鉱石"
		case "彫刻道具":
			splitLine[0] = "大きな宝石"
		case "幽霊の石", "鈍銀":
			splitLine[0] = "暗黒の合金"
		case "祈りの花":
			splitLine[0] = "大きな宝石"
		}
		itemName := splitLine[0]
		itemValue, _ := strconv.Atoi(splitLine[1])

		itemMap[itemName] += itemValue
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("ファイルの読み取り中にエラーが発生しました:", err)
		return
	}

	var outputLines []string
	for key, value := range itemMap {
		outputLines = append(outputLines, fmt.Sprintf("%s,%d", key, value))
	}

	output := strings.Join(outputLines, "\n")
	err = ioutil.WriteFile(saveFilePathItems, []byte(output), 0644)
	if err != nil {
		fmt.Println("CSVファイルの更新に失敗しました:", err)
		return
	}

	fmt.Println("CSVファイルの内容を更新しました。")
}
