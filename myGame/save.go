package myGame

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/yuyuyu2118/typingGo/player"
)

var SaveFilePath = "player\\save\\save.csv"
var SaveFilePathItems = "player\\save\\saveItems.csv"

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
	saveContent := Name + "," + MaxHP + "," + HP + "," + OP + "," + DP + "," + MaxSP + "," + SP + "," + BaseSP + "," + Gold + "," + Job + "," + AP + "," + Language + ","

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
}

func SaveFileLoad(saveFilePath string) [][]string {
	SaveFileCheck(saveFilePath)
	return CsvToSlice(saveFilePath)
}

func SaveFileCheck(saveFilePath string) {
	initializeText := "Name,MaxHP,HP,OP,DP,MaxSP,SP,BaseSP,Gold,Job,AP,\nNoName,30,30,3,1,50,0,2,0,No Job,0,\nWeaponName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,\nArmorName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,\nAccessoryName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,"
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
	} else if defeatEnemy == "Goblin" {
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
}

func SaveWeaponSellEvent(saveFilePath string, saveNum int, sellWeapon string, player *player.PlayerStatus) {
	var tempInt int
	loadContent := SaveFileLoad(saveFilePath)
	if sellWeapon == "weapon1" {
		tempInt, _ = strconv.Atoi(loadContent[3][0])
		loadContent[3][0] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[0] = loadContent[3][0]
	} else if sellWeapon == "weapon2" {
		tempInt, _ = strconv.Atoi(loadContent[3][1])
		loadContent[3][1] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[1] = loadContent[3][1]
	} else if sellWeapon == "weapon3" {
		tempInt, _ = strconv.Atoi(loadContent[3][2])
		loadContent[3][2] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[2] = loadContent[3][2]
	} else if sellWeapon == "weapon4" {
		tempInt, _ = strconv.Atoi(loadContent[3][3])
		loadContent[3][3] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[3] = loadContent[3][3]
	} else if sellWeapon == "weapon5" {
		tempInt, _ = strconv.Atoi(loadContent[3][4])
		loadContent[3][4] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[4] = loadContent[3][4]
	} else if sellWeapon == "weapon6" {
		tempInt, _ = strconv.Atoi(loadContent[3][5])
		loadContent[3][5] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[5] = loadContent[3][5]
	} else if sellWeapon == "weapon7" {
		tempInt, _ = strconv.Atoi(loadContent[3][6])
		loadContent[3][6] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[6] = loadContent[3][6]
	} else if sellWeapon == "weapon8" {
		tempInt, _ = strconv.Atoi(loadContent[3][7])
		loadContent[3][7] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[7] = loadContent[3][7]
	} else if sellWeapon == "weapon9" {
		tempInt, _ = strconv.Atoi(loadContent[3][8])
		loadContent[3][8] = strconv.Itoa(tempInt - 1)
		player.PossessedWeapon[8] = loadContent[3][8]
	} else if sellWeapon == "weapon0" {
		tempInt, _ = strconv.Atoi(loadContent[3][9])
		loadContent[3][9] = strconv.Itoa(tempInt - 1)
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
}

// TODO: アイテムのセーブ実装中
func SaveGameItems(SaveFilePathItems string, saveNum int, player *player.PlayerStatus, gainItem string) {
	SaveFileItemsCheck(SaveFilePathItems)
	tempContent := SaveFileItemsLoad(SaveFilePathItems)
	//saveContent := "NoName,30,30,3,1,50,0,2," + strconv.Itoa(player.Gold) + "," + player.Job + "," + strconv.Itoa(player.AP) + ",Japanese,"
	log.Println(tempContent[0])
	tempContent[0] = append(tempContent[0], gainItem)
	saveContent := tempContent[0]

	content, err := ioutil.ReadFile(SaveFilePathItems)
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
	lines[saveNum] = saveContent[0]

	// 更新後の内容を保存ファイルに書き込む
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(SaveFilePathItems, []byte(output), 0644)
	if err != nil {
		fmt.Println("保存ファイルの書き込みに失敗しました:", err)
		return
	}

	fmt.Println("保存ファイルを更新しました。")
}

func SaveFileItemsCheck(saveFilePathItems string) {
	initializeText := ""
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
}

func SaveFileItemsLoad(saveFilePathItems string) [][]string {
	SaveFileItemsCheck(saveFilePathItems)
	return CsvToSlice(saveFilePathItems)
}
