package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/battle"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
)

const (
	winHSize = 1440
)

var startTime time.Time
var Ticker *time.Ticker
var Frame = 0
var Last = time.Now()

func run() {
	rand.Seed(time.Now().UnixNano())                        // 乱数のシードを設定します。
	win, _ := initializeWindow()                            // ウィンドウを初期化します。
	myPos.SetCfg(winHSize)                                  // ウィンドウの高さを設定します。
	myUtil.InitTxtFontLoading()                             // テキストフォントを初期化します。
	loadContent := myGame.SaveFileLoad(myGame.SaveFilePath) // セーブファイルをロードします。
	myGame.SaveFileItemsCheck(myGame.SaveFilePathItems)     // アイテムセーブファイルをロードします。
	player := myPlayer.NewPlayerStatus(loadContent)         // プレイヤーのステータスを初期化します。
	event.InitializeEventInstance(loadContent)              // イベントのインスタンスを初期化します。
	enemy.CreateEnemySettings()                             // 敵のステータスを初期化します。

	imd := imdraw.New(nil) // フェードイン・アウト用の描画オブジェクトを初期化します。
	setTime := time.Now()  // フェードイン・アウトの開始時間を設定します。
	fadeDuration := 1.0    // フェードイン・アウトの時間を設定します。

	startTime = time.Now() // ゲーム開始時間を設定します。

	for !win.Closed() {
		switch myState.CurrentGS {
		case myState.FadeScreen:
			myState.CurrentGS = myState.StartScreen
		case myState.StartScreen: //スタート画面
			alpha, elapsedTime := myUtil.FadeScreen(win, imd, setTime, fadeDuration)
			myGame.InitStartScreen(win, myUtil.StartTxt, alpha, 1.0)
			if elapsedTime < fadeDuration {
				myUtil.DrawFadingRectangleInOut(win, imd, alpha, true) // フェードインを追加
				imd.Draw(win)
			}
		case myState.GoToScreen: //GoTo画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.StageSelect: //ダンジョンセレクト画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.TownScreen: //ショップ選択画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.WeaponShop: //武器店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myState.ArmorShop: //防具店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myState.AccessoryShop: //アクセサリー店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myState.EquipmentScreen: //装備画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.JobSelect: //職業選択画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.BlackSmithScreen: //職業選択画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myState.PlayingScreen: //戦闘画面
			if !myUtil.GetPlayerReset() {
				loadContent = myGame.SaveFileLoad(myGame.SaveFilePath)
				player = myPlayer.NewPlayerStatus(loadContent)
				myGame.SaveGame(myGame.SaveFilePath, 1, player)
				myUtil.SetPlayerReset(true)
			}

			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt)
			battle.InitPlayingBattle(win, player, time.Since(startTime))
		case myState.BattleEnemyScreen: //敵行動画面
			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt)
			battle.InitEnemyBattle(win, player, time.Since(startTime))
			myUtil.UpdateEnemyTimer(myState.CurrentGS, &startTime)
		case myState.SkillScreen: //スキル画面
			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt)
			battle.InitSkillBattle(win, player, time.Since(startTime))
		case myState.EndScreen: //リザルト画面
			loadContent := myGame.SaveFileLoad(myGame.SaveFilePath)
			event.InitializeEventInstance(loadContent)

			myGame.InitEndScreen(win, myUtil.ScreenTxt)
			myState.CurrentGS = battle.BattleEndScreen(win, myUtil.ScreenTxt, player, &enemy.EnemySettings[myGame.StageNum], loadContent)

			if !myUtil.GetSaveReset() {
				myGame.SaveGame(myGame.SaveFilePath, 1, player)
				myUtil.SetSaveReset(true)
			}
		case myState.TestState:
			myGame.TestMode(win, myUtil.ScreenTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
