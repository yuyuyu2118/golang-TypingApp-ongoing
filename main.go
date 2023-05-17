package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"log"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var t int

func init() {
	flag.IntVar(&t, "t", 1, "時間")
	flag.Parse()
}
func run() {
	//init
	windowHeightSize := 1440
	win, _ := initializeWindow(windowHeightSize)

	basicTxt := initializeAnyText("assets\\fonts\\NotoSans-Black.ttf", 40, colornames.White)
	startTxt := initializeAnyText("assets\\fonts\\NotoSans-Black.ttf", 80, colornames.White)
	endTxt := initializeAnyText("assets\\fonts\\NotoSans-Black.ttf", 60, colornames.White)

	//playerStatusインスタンスを生成
	player := newPlayerStatus(30, 30, 1, 1, 0, "")
	stage := newStageInf(0)
	enemyKnight := newEnemyStatus(100, 100, 1, 1, 30, "knight")

	words := initializeQuestion()

	var (
		//wordsMap    = make(map[string]string)
		index       = 0
		score       = 0
		collectType = 0
		missType    = 0

		startTime = time.Now()
		yourTime  = 0.0
	)

	var currentGameState GameState
	for !win.Closed() {
		switch currentGameState {
		case StartScreen:
			initStartScreen(win, startTxt, windowHeightSize)
			if win.JustPressed(pixelgl.KeyEnter) {
				currentGameState = JobSelect
				log.Println("Press:Enter -> GameState:jobSelect")
				//TODO スタートするところに持っていく
				startTime = time.Now()
			}
			//testMode
			if win.JustPressed(pixelgl.KeyT) {
				currentGameState = TestState
				log.Println("TestMode")
			}
		case JobSelect:
			initJobSelect(win, basicTxt, windowHeightSize)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) || win.JustPressed(pixelgl.Key2) || win.JustPressed(pixelgl.Key3) {
				currentGameState = jobClickEvent(win, win.MousePosition(), currentGameState, player)
			}

		case StageSelect:
			initStageSlect(win, basicTxt, windowHeightSize)

			if win.JustPressed(pixelgl.MouseButtonLeft) || win.JustPressed(pixelgl.Key1) {
				currentGameState = stageClickEvent(win, win.MousePosition(), currentGameState, stage)
			}

		case PlayingScreen:

			initPlayingScreen(win, basicTxt, windowHeightSize)

			//set Enemy Picture&HPbar
			setEnemyPic(win, enemyKnight, "assets\\monster\\monster1.png", 4.0)
			//TODO Player HPbar(右に縦長ゲージ) & PlayerSkillBar
			//TODO 手持ちアイテムバー、攻撃力や防御力の表示UI追加

			//TODO ここからplaying.goに関数化
			//set Time+rule
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
			spacing := 100.0
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
					enemyKnight.enemyHP -= 1
					//enemy Down
					if enemyKnight.enemyHP < 0 {
						index = 0
						score++
						currentGameState = EndScreen
						yourTime = float64(elapsed.Seconds())
					}
					log.Println("collectType = ", collectType)
					if index == len(question) {
						index = 0
						score++
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

		case EndScreen:
			win.Clear(colornames.Grey)
			endTxt.Clear()
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

			//画面遷移,いろいろリセット
			if win.JustPressed(pixelgl.KeyEscape) {
				currentGameState = PlayingScreen
				collectType, missType, index, score = 0, 0, 0, 0
				enemyKnight.enemyHP = enemyKnight.enemyMaxHP
				log.Println("Press:Enter -> GameState:Playing")
			} else if win.JustPressed(pixelgl.KeyTab) {
				currentGameState = StartScreen
				collectType, missType, index, score = 0, 0, 0, 0
				enemyKnight.enemyHP = enemyKnight.enemyMaxHP
				log.Println("Press:Enter -> GameState:StartScreen")
			}
		case TestState:
			testMode(win, basicTxt, windowHeightSize)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
