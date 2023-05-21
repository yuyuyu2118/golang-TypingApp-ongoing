package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
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

	//playerStatusインスタンスを生成
	player := newPlayerStatus(30, 30, 1, 1, 50, 0, 2, 0, "No Job")
	stage := newStageInf(0)

	enemy.SetCfg(winHSize)
	enemyInfo := enemy.CreateEnemyInstance()
	enemyKnight := (*enemyInfo)[0]
	// for _, enemy := range *enemyInfo {
	// 	enemyKnight := enemy
	// }

	var Ticker *time.Ticker
	for !win.Closed() {
		switch currentGameState {
		case StartScreen:
			initStartScreen(win, startTxt)
			if win.JustPressed(pixelgl.KeyEnter) {
				currentGameState = GoToScreen
				log.Println("Press:Enter -> GameState:jobSelect")
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				currentGameState = TestState
				log.Println("TestMode")
			}
		case GoToScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = goToClickEvent(win, win.MousePosition())
			}

		case StageSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				currentGameState = stageClickEvent(win, win.MousePosition(), stage)
				Ticker = time.NewTicker(time.Duration(time.Duration(enemyKnight.AttackTick) * time.Second))
				go func() {
					for range Ticker.C {
						log.Println(enemyKnight.OP)
						player.playerHP -= enemyKnight.OP
						log.Println(("Attack"))
					}
				}()
				startTime = time.Now()
			}
		case TownScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = townClickEvent(win, win.MousePosition())
			}
		case WeaponShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = weaponClickEvent(win, win.MousePosition())
			}
		case ArmorShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = armorClickEvent(win, win.MousePosition())
			}
		case AccessoryShop:
			initScreenInformation(win, descriptionTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = accessoryClickEvent(win, win.MousePosition())
			}
		case EquipmentScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = equipmentClickEvent(win, win.MousePosition())
			}
		case JobSelect:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				currentGameState = jobClickEvent(win, win.MousePosition(), player)
			}
		case SaveScreen:
			initScreenInformation(win, basicTxt, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
				currentGameState = saveClickEvent(win, win.MousePosition(), player)
			}

		case PlayingScreen:
			initScreenInformation(win, basicTxt, player)
			enemy.SetEnemyPic(win, &enemyKnight, "assets\\monster\\monster1.png", 4.0)
			enemy.SetEnemyText(win, basicTxt, &enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			setPlayerBattleInf(win, basicTxt, player)

			elapsed := initBattleText(win, basicTxt)
			currentGameState = battleTypingV1(win, player, &enemyKnight, elapsed)

		case EndScreen:
			initEndScreen(win, endTxt)
			currentGameState = battleEndScreen(win, endTxt, player, &enemyKnight)
			Ticker.Stop()
		case TestState:
			testMode(win, basicTxt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
