package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	windowHeightSize := 1440
	win, _ := initializeWindow(windowHeightSize)
	rand.Seed(time.Now().UnixNano())

	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	basicTxt := initializeAnyText(fontPath, 40, colornames.White)
	descriptionTxt := initializeAnyText(fontPath, 30, colornames.White)
	startTxt := initializeAnyText(fontPath, 80, colornames.White)
	endTxt := initializeAnyText(fontPath, 60, colornames.White)

	//playerStatusインスタンスを生成
	player := newPlayerStatus(30, 30, 1, 1, 50, 0, 2, 0, "No Job")
	stage := newStageInf(0)
	enemyKnight := newEnemyStatus(100, 100, 1, 1, 30, "knight", false, 3.0)

	var currentGameState GameState
	var Ticker *time.Ticker
	for !win.Closed() {
		switch currentGameState {
		case StartScreen:
			initStartScreen(win, startTxt, windowHeightSize)
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
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = goToClickEvent(win, win.MousePosition(), currentGameState)
			}

		case StageSelect:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				currentGameState = stageClickEvent(win, win.MousePosition(), currentGameState, stage)
				Ticker = time.NewTicker(time.Duration(time.Duration(enemyKnight.enemyAttackTick) * time.Second))
				go func() {
					for range Ticker.C {
						player.playerHP -= enemyKnight.enemyOP
						log.Println(("Attack"))
					}
				}()
				startTime = time.Now()
			}
		case TownScreen:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = townClickEvent(win, win.MousePosition(), currentGameState)
			}
		case WeaponShop:
			initScreenInformation(win, descriptionTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = weaponClickEvent(win, win.MousePosition(), currentGameState)
			}
		case ArmorShop:
			initScreenInformation(win, descriptionTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = armorClickEvent(win, win.MousePosition(), currentGameState)
			}
		case AccessoryShop:
			initScreenInformation(win, descriptionTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) || win.JustPressed(pixelgl.KeyBackspace) {
				currentGameState = accessoryClickEvent(win, win.MousePosition(), currentGameState)
			}
		case EquipmentScreen:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = equipmentClickEvent(win, win.MousePosition(), currentGameState)
			}
		case JobSelect:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				currentGameState = jobClickEvent(win, win.MousePosition(), currentGameState, player)
			}
		case SaveScreen:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
				currentGameState = saveClickEvent(win, win.MousePosition(), currentGameState, player)
			}

		case PlayingScreen:
			initScreenInformation(win, basicTxt, windowHeightSize, player, currentGameState)

			setEnemyPic(win, enemyKnight, "assets\\monster\\monster1.png", 4.0)
			setEnemyText(win, basicTxt, windowHeightSize, enemyKnight)
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加
			setPlayerBattleInf(win, basicTxt, windowHeightSize, player)

			elapsed := initBattleText(win, basicTxt, windowHeightSize)
			currentGameState = battleTypingV1(win, player, enemyKnight, elapsed, currentGameState)

		case EndScreen:
			win.Clear(colornames.Grey)
			endTxt.Clear()

			if player.playerHP > 0 {
				yourTimeString := fmt.Sprintf("%0.3f", yourTime)
				//平均キータイプ数 回/秒 Escでもう一度,Tabでタイトル
				endLines := []string{
					"YourScore : " + strconv.Itoa(score),
					"\n",
					"yourTime =" + yourTimeString,
					"collectType = " + strconv.Itoa(collectType) + " missType = " + strconv.Itoa(missType),
					"\n\n",
					"ReSTART : Press Esc | Title : Press Tab",
				}
				lineCenterAlign(win, windowHeightSize, endLines, endTxt, "center")
			} else {
				yourTimeString := fmt.Sprintf("%0.3f", yourTime)
				//平均キータイプ数 回/秒 Escでもう一度,Tabでタイトル
				endLines := []string{
					"GameOver...",
					"You have lost " + strconv.Itoa(lostGold) + " gold",
					"YourScore : " + strconv.Itoa(score),
					"\n",
					"yourTime =" + yourTimeString,
					"collectType = " + strconv.Itoa(collectType) + " missType = " + strconv.Itoa(missType),
					"\n\n",
					"ReSTART : Press Esc | Title : Press Tab",
				}
				lineCenterAlign(win, windowHeightSize, endLines, endTxt, "center")
			}

			//画面遷移,いろいろリセット
			if win.JustPressed(pixelgl.KeyEscape) {
				currentGameState = PlayingScreen
				collectType, missType, index, score = 0, 0, 0, 0
				player.playerHP = player.playerMaxHP
				enemyKnight.enemyHP = enemyKnight.enemyMaxHP
				shuffle(words)
				log.Println("Press:Enter -> GameState:Playing")
			} else if win.JustPressed(pixelgl.KeyTab) {
				currentGameState = StartScreen
				collectType, missType, index, score = 0, 0, 0, 0
				player.playerHP = player.playerMaxHP
				enemyKnight.enemyHP = enemyKnight.enemyMaxHP
				shuffle(words)
				log.Println("Press:Enter -> GameState:StartScreen")
			}
			Ticker.Stop()

		case TestState:
			testMode(win, basicTxt, windowHeightSize)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
