package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"os"
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
var saveContent string

// TODO: Utilに入れる
var language bool

func run() {
	rand.Seed(time.Now().UnixNano())
	win, _ := initializeWindow()
	myPos.SetCfg(winHSize)
	myUtil.InitTxtFontLoading()
	loadContent := myGame.SaveFileLoad(myGame.SaveFilePath)
	player := player.NewPlayerStatus(loadContent[1], loadContent[3])
	event.CreateWeaponPurchaseEvent(loadContent[2])

	enemySettings, enemyPathBar, enemySprites := enemy.CreateEnemySettings()

	frame := 0
	last := time.Now()

	for !win.Closed() {
		switch myGame.CurrentGS {
		case myGame.StartScreen:
			myGame.InitStartScreen(win, myUtil.ScreenTxt)
			if win.JustPressed(pixelgl.KeyEnter) {
				myGame.CurrentGS = myGame.GoToScreen
				log.Println("Press:Enter -> GameState:jobSelect")
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				myGame.CurrentGS = myGame.TestState
				log.Println("TestMode")
			}
			if win.JustPressed(pixelgl.KeyEscape) {
				win.Destroy()
				os.Exit(0)
			}
		case myGame.GoToScreen:
			//TODO: Saveの削除
			initScreenInformation(win, myUtil.ScreenTxt, player)
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.GoToClickEvent(win, win.MousePosition())
			}
		case myGame.StageSelect:
			initScreenInformation(win, myUtil.ScreenTxt, player)
			//TODO: Key入力受付
			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.Key7) || win.JustPressed(pixelgl.Key8) || win.JustPressed(pixelgl.Key9) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.StageClickEvent(win, win.MousePosition())
			}
		case myGame.TownScreen:
			initScreenInformation(win, myUtil.ScreenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.TownClickEvent(win, win.MousePosition())
			}
		case myGame.WeaponShop:
			initScreenInformation(win, myUtil.DescriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.Key7) || win.JustPressed(pixelgl.Key8) || win.JustPressed(pixelgl.Key9) || win.JustPressed(pixelgl.Key0) || win.JustPressed(pixelgl.KeyB) || win.JustPressed(pixelgl.KeyS) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.WeaponClickEvent(win, win.MousePosition(), player)
			}
		case myGame.ArmorShop:
			initScreenInformation(win, myUtil.DescriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.ArmorClickEvent(win, win.MousePosition())
			}
		case myGame.AccessoryShop:
			initScreenInformation(win, myUtil.DescriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.AccessoryClickEvent(win, win.MousePosition())
			}
		case myGame.EquipmentScreen:
			initScreenInformation(win, myUtil.ScreenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.EquipmentClickEvent(win, win.MousePosition())
			}
		case myGame.JobSelect:
			initScreenInformation(win, myUtil.ScreenTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.JobClickEvent(win, win.MousePosition(), player)
				myGame.SaveGame(myGame.SaveFilePath, 1, player)
			}

		case myGame.PlayingScreen:
			initScreenInformation(win, myUtil.BasicTxt, player)
			log.Println(myGame.StageNum)

			dt := time.Since(last).Seconds()
			if dt >= 0.2 { // アニメーション速度を調整 (ここでは0.2秒ごとに更新)
				frame = (frame + 1) % len(enemySprites[myGame.StageNum])
				last = time.Now()
			}
			enemy.SetEnemySprite(win, &enemySettings[myGame.StageNum], enemyPathBar[myGame.StageNum], enemySprites[myGame.StageNum], frame)
			enemy.SetEnemySpriteText(win, myUtil.ScreenTxt, &enemySettings[myGame.StageNum])
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, myUtil.BasicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, myUtil.BasicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemySettings[myGame.StageNum], elapsed)
			if myGame.CurrentGS == myGame.BattleEnemyScreen {
				startTime = time.Now()
			}
		case myGame.BattleEnemyScreen:
			initScreenInformation(win, myUtil.BasicTxt, player)

			dt := time.Since(last).Seconds()
			if dt >= 0.2 { // アニメーション速度を調整 (ここでは0.2秒ごとに更新)
				frame = (frame + 1) % len(enemySprites[myGame.StageNum])
				last = time.Now()
			}
			enemy.SetEnemySprite(win, &enemySettings[myGame.StageNum], enemyPathBar[myGame.StageNum], enemySprites[myGame.StageNum], frame)
			enemy.SetEnemySpriteText(win, myUtil.ScreenTxt, &enemySettings[myGame.StageNum])
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, myUtil.BasicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, myUtil.BasicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemySettings[myGame.StageNum], elapsed)
			if myGame.CurrentGS == myGame.PlayingScreen {
				startTime = time.Now()
			}
		case myGame.EndScreen:
			loadContent := myGame.SaveFileLoad(myGame.SaveFilePath)
			event.CreateWeaponPurchaseEvent(loadContent[2])

			myGame.InitEndScreen(win, myUtil.ScreenTxt)
			myGame.CurrentGS = battle.BattleEndScreen(win, myUtil.ScreenTxt, player, &enemySettings[myGame.StageNum])

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
