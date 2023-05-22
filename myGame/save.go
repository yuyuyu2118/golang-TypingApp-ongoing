package myGame

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/faiface/pixel"
)

var (
	save1Button = pixel.Rect{}
	save2Button = pixel.Rect{}
)

// func InitSave(win *pixelgl.Window, Txt *text.Text) {

// 	Txt.Clear()
// 	Txt.Color = colornames.White
// 	fmt.Fprintln(Txt, "Do you want to save?")
// 	tempPosition = myPos.TopCenterPos(win, Txt)
// 	myPos.DrawPos(win, Txt, tempPosition)

// 	Txt.Clear()
// 	Txt.Color = colornames.White
// 	fmt.Fprintln(Txt, "1. Yes")
// 	tempPosition = myPos.CenterLeftPos(win, Txt)
// 	myPos.DrawPos(win, Txt, tempPosition)
// 	save1Button = Txt.Bounds().Moved(tempPosition)

// 	Txt.Clear()
// 	Txt.Color = colornames.White
// 	fmt.Fprintln(Txt, "2. No")
// 	tempPosition = myPos.CenterPos(win, Txt)
// 	myPos.DrawPos(win, Txt, tempPosition)
// 	save2Button = Txt.Bounds().Moved(tempPosition)
// }

// func SaveClickEvent(win *pixelgl.Window, mousePos pixel.Vec, player *player.PlayerStatus) GameState {
// 	//TODO ページを作成したら追加
// 	if save1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
// 		SaveGame(player)
// 		CurrentGS = GoToScreen
// 		log.Println("Save Done!")
// 	} else if save1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
// 		CurrentGS = GoToScreen
// 		log.Println("saveScreen->GoToScreen")
// 	}
// 	return CurrentGS
// }

func SaveGame(saveFilePath string, saveNum int, saveContent string) {
	SaveFileCheck(saveFilePath)

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

func SaveFileLoad(saveFilePath string) [][]string {
	return CsvToSlice(saveFilePath)
}

func LoadSliceToString(content []string) string {
	quotedValues := make([]string, len(content))
	for i, v := range content {
		quotedValues[i] = `"` + v + `"`
	}
	return strings.Join(quotedValues, ",")
}
