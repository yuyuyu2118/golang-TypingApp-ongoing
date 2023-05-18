package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var t int

// func init() {
// 	flag.IntVar(&t, "t", 1, "時間")
// 	flag.Parse()
// }

func run() {
	windowHeightSize := 1440
	win, _ := initializeWindow(windowHeightSize)
	rand.Seed(time.Now().UnixNano())

	fontPath := "assets\\fonts\\NotoSans-Black.ttf"
	basicTxt := initializeAnyText(fontPath, 40, colornames.White)
	startTxt := initializeAnyText(fontPath, 80, colornames.White)
	endTxt := initializeAnyText(fontPath, 60, colornames.White)

	//playerStatusインスタンスを生成
	player := newPlayerStatus(30, 30, 1, 1, 50, 0, 2, 0, "")
	stage := newStageInf(0)
	enemyKnight := newEnemyStatus(100, 100, 1, 1, 30, "knight", false, 3.0)

	words := initializeQuestion()

	var (
		//wordsMap    = make(map[string]string)
		index       = 0
		score       = 0
		collectType = 0
		missType    = 0

		startTime = time.Now()
		yourTime  = 0.0
		gainGold  = 0
		lostGold  = 0
	)

	var currentGameState GameState
	var Ticker *time.Ticker
	for !win.Closed() {
		switch currentGameState {
		case StartScreen:
			initStartScreen(win, startTxt, windowHeightSize)
			if win.JustPressed(pixelgl.KeyEnter) {
				currentGameState = GoToScreen
				log.Println("Press:Enter -> GameState:jobSelect")
				//TODO スタートするところに持っていく
				startTime = time.Now()
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				currentGameState = TestState
				log.Println("TestMode")
			}
		case GoToScreen:
			initGoToScreen(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = goToClickEvent(win, win.MousePosition(), currentGameState)
			}

		case StageSelect:
			initStageSlect(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				currentGameState = stageClickEvent(win, win.MousePosition(), currentGameState, stage)
				Ticker = time.NewTicker(time.Duration(time.Duration(enemyKnight.enemyAttackTick) * time.Second))
				go func() {
					for range Ticker.C {
						player.playerHP -= enemyKnight.enemyOP
						log.Println(("Attack"))
					}
				}()
			}
		case TownScreen:
			initTownScreen(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = townClickEvent(win, win.MousePosition(), currentGameState)
			}
		case WeaponShop:
			initWeaponShop(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = weaponClickEvent(win, win.MousePosition(), currentGameState)
			}
		case EquipmentScreen:
			initEquipmentScreen(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) || win.JustPressed(pixelgl.Key4) || win.JustPressed(pixelgl.Key5) || win.JustPressed(pixelgl.Key6) {
				currentGameState = equipmentClickEvent(win, win.MousePosition(), currentGameState)
			}
		case JobSelect:
			initJobSelect(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				currentGameState = jobClickEvent(win, win.MousePosition(), currentGameState, player)
			}
		case SaveScreen:
			initSaveScreen(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) {
				currentGameState = saveClickEvent(win, win.MousePosition(), currentGameState, player)
			}

		case PlayingScreen:

			initPlayingScreen(win, basicTxt, windowHeightSize)
			initPlayerGold(win, basicTxt, windowHeightSize, player)

			//set Enemy Picture&HPbar
			setEnemyPic(win, enemyKnight, "assets\\monster\\monster1.png", 4.0)
			setPlayerInf(win, player)
			//TODO Player HPbar(右に縦長ゲージ) & PlayerSkillBar
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加

			//TODO ここからplaying.goに関数化
			//set Time+rule
			basicTxt.Clear()
			basicTxt.Color = colornames.White
			fmt.Fprintln(basicTxt, "EnemyHP : ", enemyKnight.enemyHP)
			drawPos(win, basicTxt, topCenterPos(win, basicTxt, windowHeightSize))

			question := words[score]
			temp := []byte(question)
			typed := win.Typed()

			basicTxt.Clear()
			basicTxt.Color = colornames.White
			fmt.Fprintln(basicTxt, "> ", words[score])
			drawPos(win, basicTxt, bottleRoundCenterPos(win, basicTxt, windowHeightSize))

			offset := basicTxt.Bounds().W()
			basicTxtOrigX := basicTxt.Dot.X
			spacing := 60.0
			if len(words)-score != 1 {
				basicTxt.Color = colornames.Darkgray
				offset := basicTxt.Bounds().W()
				basicTxt.Clear()
				fmt.Fprintln(basicTxt, words[score+1])
				drawPos(win, basicTxt, bottleRoundCenterPos(win, basicTxt, windowHeightSize).Add(pixel.V(offset+spacing, 0)))
				basicTxt.Dot.X = basicTxtOrigX
			}
			if !(len(words)-score == 2 || len(words)-score == 1) {
				basicTxt.Color = colornames.Gray
				offset += basicTxt.Bounds().W()
				basicTxt.Clear()
				fmt.Fprintln(basicTxt, words[score+2])
				drawPos(win, basicTxt, bottleRoundCenterPos(win, basicTxt, windowHeightSize).Add(pixel.V(offset+spacing*2, 0)))
			}
			//basicTxt.Dot.X = basicTxtOrigX

			basicTxt.Color = colornames.White
			basicTxt.Clear()
			fmt.Fprintln(basicTxt, "\n\n", "collectType = ", collectType, " missType = ", missType)
			drawPos(win, basicTxt, bottleRoundCenterPos(win, basicTxt, windowHeightSize))
			basicTxt.Dot.X = basicTxtOrigX

			//set Time+rule
			basicTxt.Clear()
			basicTxt.Color = colornames.White
			elapsed := time.Since(startTime)
			fmt.Fprintln(basicTxt, "time = ", elapsed.Milliseconds())
			drawPos(win, basicTxt, bottleLeftPos(win, basicTxt, windowHeightSize))
			//TODO ここまで

			//TODO タイピングの処理部分で、スペースキーによるスキルの発動処理、一定間隔での敵の攻撃の追加
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					enemyKnight.enemyHP -= 3
					player.playerSP += player.playerBaseSP
					//enemy Down
					if enemyKnight.enemyHP <= 0 {
						//GoldRandom
						min := int(float64(enemyKnight.enemyGold) * 0.7)
						max := int(float64(enemyKnight.enemyGold) * 1.3)
						gainGold = rand.Intn(max-min+1) + min
						player.playerGold += gainGold
						index = 0
						score++
						currentGameState = EndScreen
						yourTime = float64(elapsed.Seconds())
					}
					log.Println("collectType = ", collectType)
					//1 word end type
					if index == len(question) {
						index = 0
						score++
						//全部打ち切っちゃった場合
						if score == len(words) {
							currentGameState = EndScreen
							yourTime = float64(elapsed.Seconds())
						}
					}
				} else {
					missType++
					log.Println("missType = ", missType)
				}
			}

			if win.JustPressed(pixelgl.KeySpace) {
				log.Println("Skill!!!")
				player.playerSP = 0
				if player.playerJob == "Warrior" {
					enemyKnight.enemyHP -= 15
				} else if player.playerJob == "Priest" {
					//TODO 僧侶の回復スキル
				} else if player.playerJob == "Wizard" {
					//TODO 魔法使いの時止めスキル
				}
			}
			if player.playerHP <= 0 {
				yourTime = float64(elapsed.Seconds())
				min := int(float64(enemyKnight.enemyGold) * 0.7)
				max := int(float64(enemyKnight.enemyGold) * 1.3)
				lostGold = rand.Intn(max-min+1) + min
				player.playerGold -= lostGold
				log.Println("GameOver!!")
				currentGameState = EndScreen
			}
		case EndScreen:
			win.Clear(colornames.Grey)
			endTxt.Clear()
			log.Println(yourTime)

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
