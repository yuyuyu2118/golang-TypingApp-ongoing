package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/battle"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

const (
	winHSize = 1440
)

var startTime time.Time
var Ticker *time.Ticker

func run() {
	win, _ := initializeWindow()
	rand.Seed(time.Now().UnixNano())

	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	basicTxt := initializeAnyText(fontPath, 40, colornames.White)
	descriptionTxt := initializeAnyText(fontPath, 30, colornames.White)
	startTxt := initializeAnyText(fontPath, 80, colornames.White)
	endTxt := initializeAnyText(fontPath, 60, colornames.White)

	myPos.SetCfg(winHSize)
	//playerStatusインスタンスを生成
	stage := myGame.NewStageInf(0)
	player := player.NewPlayerStatus(30, 30, 3, 1, 50, 0, 2, 0, "No Job")
	enemyInfo := enemy.CreateEnemyInstance()
	enemyKnight := (*enemyInfo)[0]
	// for _, enemy := range *enemyInfo {
	// 	enemyKnight := enemy
	// }

	for !win.Closed() {
		switch myGame.CurrentGS {
		case myGame.StartScreen:
			myGame.InitStartScreen(win, startTxt)
			if win.JustPressed(pixelgl.KeyEnter) {
				myGame.CurrentGS = myGame.GoToScreen
				log.Println("Press:Enter -> GameState:jobSelect")
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				myGame.CurrentGS = myGame.TestState
				log.Println("TestMode")
			}
		case myGame.GoToScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				myGame.CurrentGS = myGame.GoToClickEvent(win, win.MousePosition())
			}

		case myGame.StageSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				myGame.CurrentGS = myGame.StageClickEvent(win, win.MousePosition(), stage)
			}
		case myGame.TownScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				myGame.CurrentGS = myGame.TownClickEvent(win, win.MousePosition())
			}
		case myGame.WeaponShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.WeaponClickEvent(win, win.MousePosition())
			}
		case myGame.ArmorShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.ArmorClickEvent(win, win.MousePosition())
			}
		case myGame.AccessoryShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = myGame.AccessoryClickEvent(win, win.MousePosition())
			}
		case myGame.EquipmentScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				myGame.CurrentGS = myGame.EquipmentClickEvent(win, win.MousePosition())
			}
		case myGame.JobSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				myGame.CurrentGS = myGame.JobClickEvent(win, win.MousePosition(), player)
			}
		case myGame.SaveScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
				myGame.CurrentGS = myGame.SaveClickEvent(win, win.MousePosition(), player)
			}

		case myGame.PlayingScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", enemyKnight.EnemySize)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemyKnight, elapsed)
			if myGame.CurrentGS == myGame.BattleEnemyScreen {
				startTime = time.Now()
			}
		case myGame.BattleEnemyScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", enemyKnight.EnemySize)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := time.Since(startTime)
			battle.InitBattleTextV2(win, basicTxt, elapsed)
			myGame.CurrentGS = battle.BattleTypingV2(win, player, &enemyKnight, elapsed)
			if myGame.CurrentGS == myGame.PlayingScreen {
				startTime = time.Now()
			}
		case myGame.EndScreen:
			myGame.InitEndScreen(win, endTxt)
			myGame.CurrentGS = battle.BattleEndScreen(win, endTxt, player, &enemyKnight)
		case myGame.TestState:
			myGame.TestMode(win, basicTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
