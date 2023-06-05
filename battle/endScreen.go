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
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
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

func BattleEndScreen(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus, enemy *enemy.EnemyStatus, loadContent [][]string) myState.GameState {
	xOffSet := 100.0
	xOffSet2 := 600.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	yOffSet2 := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 260
	txtPos := pixel.V(0, 0)
	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White
	fmt.Fprintln(myUtil.ScreenTxt, "リザルト  再戦 : Press Tab | 町に戻る : Press BackSpace")
	tempPosition = myPos.TopCenPos(win, myUtil.ScreenTxt)
	myPos.DrawPos(win, myUtil.ScreenTxt, tempPosition)
	//DropEvent
	if !dropEvent {
		for i := 0; i < 4; i++ {
			if rand.Float64() <= 0.45 { // 40%の確率でアイテムをドロップ
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

	ClearTxt := []string{"正解タイプボーナスゴールド: " + strconv.Itoa(collectType) + " * 0.3 = " + strconv.Itoa(gainGoldCollectType),
		"モンスタードロップゴールド: " + strconv.Itoa(gainGold) + "S",
		"",
		" 入力単語数:" + strconv.Itoa(wordsNum),
		"正解タイプ数:" + strconv.Itoa(collectType),
		"正解タイプ数:" + strconv.Itoa(collectType),
		"ミスタイプ数:" + strconv.Itoa(missType),
		"",
		"獲得AP:" + strconv.Itoa(enemy.DropAP),
	}
	var tempName = "ドロップ素材:"
	for _, item := range dropRandomItem {
		tempName += " " + item
	}
	ClearTxt = append(ClearTxt, tempName)

	//新武器や新ダンジョンを解放した旨のメッセージ
	for i, v := range loadContent[2] {
		if myGame.StageNum == i && v == strconv.Itoa(1) {
			for _, v := range unlockElements[i] {
				ClearTxt = append(ClearTxt, v)
			}
		}
	}

	if player.HP > 0 {
		for i, value := range ClearTxt {
			if i < 10 {
				myUtil.ScreenTxt.Clear()
				myUtil.ScreenTxt.Color = colornames.White
				fmt.Fprintln(myUtil.ScreenTxt, value)
				yOffSet -= myUtil.ScreenTxt.LineHeight + 20
				txtPos = pixel.V(xOffSet, yOffSet)
				tempPosition := pixel.IM.Moved(txtPos)
				myUtil.ScreenTxt.Draw(win, tempPosition)
			} else if i >= 10 {
				myUtil.ScreenTxt.Clear()
				myUtil.ScreenTxt.Color = colornames.White
				fmt.Fprintln(myUtil.ScreenTxt, value)
				yOffSet2 -= myUtil.ScreenTxt.LineHeight + 20
				txtPos = pixel.V(xOffSet2, yOffSet2)
				tempPosition := pixel.IM.Moved(txtPos)
				myUtil.ScreenTxt.Draw(win, tempPosition)
			}
		}

		//TODO: 最初の討伐時だけ武器の追加を教える.
		// tempInt, _ := strconv.Atoi(player.PossessedWeapon[0])
		// if tempInt >= 1 {
		// 	endLines = append(endLines, "武器屋に新しい武器が追加された<-New!!")
		// }
		// log.Println(endLines)
	} else {
		//平均キータイプ数 回/秒 Escでもう一度,Tabでタイトル
		endLines := []string{
			"あなたは負けてしまいました",
			"失ったゴールド" + strconv.Itoa(lostGold) + " gold",
			" 入力単語数:" + strconv.Itoa(wordsNum),
			"正解タイプ数:" + strconv.Itoa(collectType),
			"正解タイプ数:" + strconv.Itoa(collectType),
			"ミスタイプ数:" + strconv.Itoa(missType),
		}

		for i, value := range endLines {
			if i < 10 {
				myUtil.ScreenTxt.Clear()
				myUtil.ScreenTxt.Color = colornames.White
				fmt.Fprintln(myUtil.ScreenTxt, value)
				yOffSet -= myUtil.ScreenTxt.LineHeight + 20
				txtPos = pixel.V(xOffSet, yOffSet)
				tempPosition := pixel.IM.Moved(txtPos)
				myUtil.ScreenTxt.Draw(win, tempPosition)
			} else if i >= 10 {
				myUtil.ScreenTxt.Clear()
				myUtil.ScreenTxt.Color = colornames.White
				fmt.Fprintln(myUtil.ScreenTxt, value)
				yOffSet2 -= myUtil.ScreenTxt.LineHeight + 20
				txtPos = pixel.V(xOffSet2, yOffSet2)
				tempPosition := pixel.IM.Moved(txtPos)
				myUtil.ScreenTxt.Draw(win, tempPosition)
			}
		}
	}

	//画面遷移,いろいろリセット
	if win.JustPressed(pixelgl.KeyTab) {
		myState.CurrentGS = myState.PlayingScreen
		collectType, missType, index, wordsNum = 0, 0, 0, 0
		magicCollectType, magicMissType = 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		enemy.HP = enemy.MaxHP
		dropRandomItem = dropRandomItem[:0]
		dropEvent = false
		battleTimeBool = false
		Shuffle(words)
		myUtil.SetSaveReset(false)
		log.Println("Press:Enter -> GameState:Playing")
	} else if win.JustPressed(pixelgl.KeyBackspace) {
		myState.CurrentGS = myState.GoToScreen
		collectType, missType, index, wordsNum = 0, 0, 0, 0
		magicCollectType, magicMissType = 0, 0
		player.HP = player.MaxHP
		player.SP = 0
		enemy.HP = enemy.MaxHP
		dropRandomItem = dropRandomItem[:0]
		dropEvent = false
		battleTimeBool = false
		Shuffle(words)
		myUtil.SetSaveReset(false)
		log.Println("Press:Enter -> GameState:GoToScreen")
	}

	return myState.CurrentGS
}
