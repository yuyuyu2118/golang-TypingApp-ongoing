package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	event "github.com/yuyuyu2118/typingGo/Event"
	"github.com/yuyuyu2118/typingGo/battle"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

const (
	winHSize = 1440
)

var startTime time.Time
var Ticker *time.Ticker
var Frame = 0
var Last = time.Now()

func run() {
	rand.Seed(time.Now().UnixNano())
	win, _ := initializeWindow()
	myPos.SetCfg(winHSize)
	myUtil.InitTxtFontLoading()
	loadContent := myGame.SaveFileLoad(myGame.SaveFilePath)
	player := player.NewPlayerStatus(loadContent[1], loadContent[3])
	event.CreateWeaponPurchaseEvent(loadContent[2])
	enemy.CreateEnemySettings()

	for !win.Closed() {
		switch myGame.CurrentGS {
		case myGame.StartScreen: //スタート画面
			myGame.InitStartScreen(win, myUtil.ScreenTxt)
		case myGame.GoToScreen: //GoTo画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myGame.StageSelect: //ダンジョンセレクト画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myGame.TownScreen: //ショップ選択画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myGame.WeaponShop: //武器店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myGame.ArmorShop: //防具店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myGame.AccessoryShop: //アクセサリー店
			initScreenInformation(win, myUtil.DescriptionTxt, player)
		case myGame.EquipmentScreen: //装備画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myGame.JobSelect: //職業選択画面
			initScreenInformation(win, myUtil.ScreenTxt, player)
		case myGame.PlayingScreen: //戦闘画面
			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt) //TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			battle.InitPlayingBattle(win, player, time.Since(startTime))
			myUtil.UpdatePlayingTimer(myGame.CurrentGS, &startTime)
		case myGame.BattleEnemyScreen: //敵行動画面
			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt) //TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			battle.InitEnemyBattle(win, player, time.Since(startTime))
			myUtil.UpdateEnemyTimer(myGame.CurrentGS, &startTime)
		case myGame.SkillScreen: //スキル画面
			initScreenInformation(win, myUtil.BasicTxt, player)

			enemy.StartEnemyAnimation(win, &Last, &Frame)
			player.SetPlayerBattleInf(win, myUtil.BasicTxt) //TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			battle.InitSkillBattle(win, player, time.Since(startTime))
			myUtil.UpdateEnemyTimer(myGame.CurrentGS, &startTime)
		case myGame.EndScreen: //リザルト画面
			loadContent := myGame.SaveFileLoad(myGame.SaveFilePath)
			event.CreateWeaponPurchaseEvent(loadContent[2])

			myGame.InitEndScreen(win, myUtil.ScreenTxt)
			myGame.CurrentGS = battle.BattleEndScreen(win, myUtil.ScreenTxt, player, &enemy.EnemySettings[myGame.StageNum])

			if !myUtil.GetSaveReset() {
				myGame.SaveGame(myGame.SaveFilePath, 1, player)
				myUtil.SetSaveReset(true)
			}
		case myGame.TestState:
			myGame.TestMode(win, myUtil.ScreenTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
