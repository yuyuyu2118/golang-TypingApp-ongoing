package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

const (
	winHSize = 1440
)

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
	player := player.NewPlayerStatus(30, 30, 1, 1, 50, 0, 2, 0, "No Job")

	stage := newStageInf(0)

	enemyInfo := enemy.CreateEnemyInstance()
	enemyKnight := (*enemyInfo)[0]
	// for _, enemy := range *enemyInfo {
	// 	enemyKnight := enemy
	// }

	var Ticker *time.Ticker
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
				myGame.CurrentGS = goToClickEvent(win, win.MousePosition())
			}

		case myGame.StageSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				myGame.CurrentGS = stageClickEvent(win, win.MousePosition(), stage)
				Ticker = time.NewTicker(time.Duration(time.Duration(enemyKnight.AttackTick) * time.Second))
				go func() {
					for range Ticker.C {
						log.Println(enemyKnight.OP)
						player.HP -= enemyKnight.OP
						log.Println(("Attack"))
					}
				}()
				startTime = time.Now()
			}
		case myGame.TownScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				myGame.CurrentGS = townClickEvent(win, win.MousePosition())
			}
		case myGame.WeaponShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = weaponClickEvent(win, win.MousePosition())
			}
		case myGame.ArmorShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = armorClickEvent(win, win.MousePosition())
			}
		case myGame.AccessoryShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				myGame.CurrentGS = accessoryClickEvent(win, win.MousePosition())
			}
		case myGame.EquipmentScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				myGame.CurrentGS = equipmentClickEvent(win, win.MousePosition())
			}
		case myGame.JobSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				myGame.CurrentGS = jobClickEvent(win, win.MousePosition(), player)
			}
		case myGame.SaveScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
				myGame.CurrentGS = myUtil.SaveClickEvent(win, win.MousePosition(), player)
			}

		case myGame.PlayingScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", 4.0)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			player.SetPlayerBattleInf(win, basicTxt)

			elapsed := initBattleText(win, basicTxt)
			myGame.CurrentGS = battleTypingV1(win, player, &enemyKnight, elapsed)

		case myGame.EndScreen:
			myGame.InitEndScreen(win, endTxt)
			myGame.CurrentGS = battleEndScreen(win, endTxt, player, &enemyKnight)
			Ticker.Stop()
		case myGame.TestState:
			testMode(win, basicTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
