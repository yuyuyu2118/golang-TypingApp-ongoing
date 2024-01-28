package battle

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

var tempPosition pixel.Vec
var dropRandomItem = []string{}
var dropEvent bool

var unlockElements = [][]string{{"新職業: 見習い剣士", "新武器: 木の棒", "新防具: 草織りのローブ", "新アクセサリー: 樹木のペンダント", "新ダンジョン解放: バード"},
	{"新職業: 狩人", "新武器: 果物ナイフ", "新防具: フルーツアーマー", "新アクセサリー: フルーツブレスレット", "新ダンジョン解放: プラント"},
	{"新職業: モンク", "新武器: 木刀", "新防具: 木の鎧", "新アクセサリー: 平和のバンド", "新ダンジョン解放: コボルト"},
	{"新職業: 魔法使い", "新武器: ドレインソード", "新防具: ソウルバインドプレート", "新アクセサリー: ライフリンクのリング", "新ダンジョン解放: ゾンビ"},
	{"新職業: 化け物", "新武器: スタンハンマー", "新防具: スタンプレート", "新アクセサリー: ショックウェーブリング", "新ダンジョン解放: フェアリー"},
	{"新武器: 鉄の剣", "新防具: 鉄の鎧", "新アクセサリー: 鉄のブレスレット", "新ダンジョン解放: スカル"},
	{"新武器: 隼の剣", "新防具: 飛翔のマント", "新アクセサリー: 疾走のリング", "新ダンジョン解放: ウィザード"},
	{"新武器: 勇者の剣", "新防具: 勇者の鎧", "新アクセサリー: 勇者のペンダント", "新ダンジョン解放: ソルジャー"},
	{"新武器: 名刀村正", "新防具: 刃舞の衣", "新アクセサリー: 刀匠の指輪", "新ダンジョン解放: ドラゴン"},
	{"新武器: 死神の大鎌", "新防具: 冥界の鎧", "新アクセサリー: 霊魂のイヤリング", "新モード解放: 裏"},
}

func BattleEndScreen(win *pixelgl.Window, Txt *text.Text, player *myPlayer.PlayerStatus, enemy *enemy.EnemyStatus, loadContent [][]string) myState.GameState {
	//DropEvent
	if !dropEvent {
		if rand.Float64() <= 0.7 { // 5%の確率でアイテムをドロップ
			dropRandomItem = append(dropRandomItem, enemy.DropItems[0])
		}
		for i := 1; i < 4; i++ {
			if rand.Float64() <= 0.40 { // 40%の確率でアイテムをドロップ
				dropRandomItem = append(dropRandomItem, enemy.DropItems[i])
			}
		}
		if rand.Float64() <= 0.05 { // 5%の確率でアイテムをドロップ
			dropRandomItem = append(dropRandomItem, enemy.DropItems[4])
		}
		myGame.SaveFileItemsLoad(myGame.SaveFilePathItems)
		myGame.SaveGameItems(myGame.SaveFilePathItems, dropRandomItem)
		dropEvent = true
	}

	// 結果表示用のMessageBoxを作成
	resultMessageBox := myPos.NewMessageBox(win, myUtil.MessageTxt, colornames.White, colornames.White, 5, 0, 0, 1, 0.5)
	resultMessageBox.DrawMessageBox()

	// 結果テキストの構築
	var resultText string
	if player.HP > 0 {
		// クリア時のメッセージ
		resultText += fmt.Sprintf("正解タイプボーナスゴールド: %d x 0.5 = %dS\n", collectType, gainGoldCollectType)
		resultText += fmt.Sprintf("モンスタードロップゴールド: %dS\n\n", gainGold)
		resultText += fmt.Sprintf("入力単語数: %d\n", wordsNum)
		resultText += fmt.Sprintf("正解タイプ数: %d\n", collectType)
		resultText += fmt.Sprintf("ミスタイプ数: %d\n\n", missType)
		resultText += fmt.Sprintf("獲得AP: %d\n", enemy.DropAP)

		var tempName = "ドロップ素材:"
		for _, item := range dropRandomItem {
			tempName += " " + item
		}
		resultText += tempName + "\n"

		// 新武器や新ダンジョンを解放した旨のメッセージを追加
		for i, v := range loadContent[2] {
			if myGame.StageNum == i && v == strconv.Itoa(1) {
				for _, unlockInfo := range unlockElements[i] {
					resultText += unlockInfo + "\n"
				}
			}
		}

		resultText += "リザルト  再戦 : Press Tab | 町に戻る : Press BackSpace\n"
	} else {
		// 敗北時のメッセージ
		resultText += "あなたは負けてしまいました\n"
		resultText += fmt.Sprintf("失ったゴールド: %dS\n", lostGold)
		resultText += fmt.Sprintf("入力単語数: %d\n", wordsNum)
		resultText += fmt.Sprintf("正解タイプ数: %d\n", collectType)
		resultText += fmt.Sprintf("ミスタイプ数: %d\n", missType)
		resultText += "リザルト  再戦 : Press Tab | 町に戻る : Press BackSpace\n"
	}

	// MessageBoxに結果テキストを描画
	resultMessageBox.DrawMessageTxt(resultText)

	//画面遷移,いろいろリセット
	if win.JustPressed(pixelgl.KeyTab) {
		myState.CurrentGS = myState.PlayingScreen
		collectType, missType, index, wordsNum = 0, 0, 0, 0
		magicCollectType, magicMissType = 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		player.OP = player.MaxOP
		player.DP = player.MaxDP
		enemy.HP = enemy.MaxHP
		dropRandomItem = dropRandomItem[:0]
		dropEvent = false
		battleTimeBool = false
		Shuffle(words)
		myUtil.SetSaveReset(false)
		myUtil.SetPlayerReset(false)
		log.Println("Press:Enter -> GameState:Playing")
	} else if win.JustPressed(pixelgl.KeyBackspace) {
		myState.CurrentGS = myState.GoToScreen
		collectType, missType, index, wordsNum = 0, 0, 0, 0
		magicCollectType, magicMissType = 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		player.OP = player.MaxOP
		player.DP = player.MaxDP
		enemy.HP = enemy.MaxHP
		dropRandomItem = dropRandomItem[:0]
		dropEvent = false
		battleTimeBool = false
		Shuffle(words)
		myUtil.SetSaveReset(false)
		myUtil.SetPlayerReset(false)
		log.Println("Press:Enter -> GameState:GoToScreen")
	}

	return myState.CurrentGS
}
