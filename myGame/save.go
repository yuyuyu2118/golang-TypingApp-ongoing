package myGame

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	tempInitText := []string{"Name,MaxHP,HP,OP,DP,MaxSP,SP,BaseSP,Gold,Job,AP,language,",
		"NoName,30,30,5,1,50,0,2,0,No Job,0,Japanese,",
		"0,0,0,0,0,0,0,0,0,0,,,",
		"0,0,0,0,0,0,0,0,0,0,,,",
		"WeaponName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,,",
		"ArmorName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,,",
		"AccessoryName,Buy,Sell,Required Materials,Materials1,Materials2,Materials3,Attack Power,Unique Abilities,,,,",
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
	initializeText := " "
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
