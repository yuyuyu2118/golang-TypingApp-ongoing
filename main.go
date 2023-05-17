package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	_ "image/png"
	"io"
	"log"
	"os"
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
	windowHeightSize := 1440
	win, _ := initializeWindow(windowHeightSize)
	face, _ := loadTTF("NotoSans-Black.ttf", 40)
	basicTxt := initializeText(face, colornames.White)

	pic, _ := openDecodePictureData("assets\\monster\\monster1.png")
	picMonster := pixel.NewSprite(pic, pic.Bounds())

	//playerStatusインスタンスを生成
	player := newPlayerStatus(30, 30, 1, 1, 0, "")

	var (
		words = []string{}
		//wordsMap    = make(map[string]string)
		index       = 0
		score       = 0
		collectType = 0
		missType    = 0

		startTime  = time.Now()
		yourTime   = 0.0
		tempString = ""
	)
	enemy := enemyStatus{}
	enemy.enemyHP = 100
	enemy.enemyMaxHP = enemy.enemyHP

	// //csvRead
	file, _ := os.Open("question2_4.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		//wordsMap[record[1]] = record[2]
		words = append(words, record[2])
	}
	shuffle(words)
	//wordsMap = shuffleMap(wordsMap)

	var currentGameState GameState

	startFace, _ := loadTTF("NotoSans-Black.ttf", 80)
	startTxt := initializeText(startFace, colornames.White)
	startLines := []string{
		"This is a TypingBattleGame",
		"\n",
		"START : Press Enter",
	}

	swordsManButton := pixel.Rect{}
	jobSelectFace, _ := loadTTF("NotoSans-Black.ttf", 80)
	jobSelectTxt := initializeText(jobSelectFace, colornames.White)
	jobSelectLines := []string{
		"Select your Job",
		"\n",
		"job1 : Swordsman",
	}

	stage1Button := pixel.Rect{}
	stageSelectFace, _ := loadTTF("NotoSans-Black.ttf", 80)
	stageSelectTxt := initializeText(stageSelectFace, colornames.White)
	stageSelectLines := []string{
		"Select play Stage",
		"\n",
		"stage1 : VS Knight",
	}

	endFace, _ := loadTTF("NotoSans-Black.ttf", 60)
	endTxt := initializeText(endFace, colornames.White)

	for !win.Closed() {
		switch currentGameState {
		case StartScreen:

			//windowのリセットとテキストの描画
			win.Clear(colornames.Darkcyan)
			startTxt.Clear()

			lineCenterAlign(win, windowHeightSize, startLines, startTxt, "center")

			//画面遷移
			if win.JustPressed(pixelgl.KeyEnter) {
				currentGameState = JobSelect
				log.Println("Press:Enter -> GameState:jobSelect")
				//TODO スタートするところに持っていく
				startTime = time.Now()
			}
			if win.JustPressed(pixelgl.KeyT) {
				currentGameState = TestState
				log.Println("TestMode")
			}

		case JobSelect:

			//windowのリセットとテキストの描画
			win.Clear(colornames.Black)
			jobSelectTxt.Clear()

			lineCenterAlign(win, windowHeightSize, jobSelectLines, jobSelectTxt, "center")
			swordsManButton = jobSelectTxt.Bounds()

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mousePos := win.MousePosition()
				// クリックがStart テキストの範囲内で発生した場合、画面遷移を実行
				if swordsManButton.Contains(mousePos) {
					currentGameState = StageSelect
					log.Println("YourJob is swordsMan")
					player.playerJob = "swordsMan"
				}
			}

		case StageSelect:

			win.Clear(colornames.Black)
			stageSelectTxt.Clear()

			lineCenterAlign(win, windowHeightSize, stageSelectLines, stageSelectTxt, "center")
			stage1Button = stageSelectTxt.Bounds()

			if win.JustPressed(pixelgl.MouseButtonLeft) {
				mousePos := win.MousePosition()
				// クリックがStart テキストの範囲内で発生した場合、画面遷移を実行
				if stage1Button.Contains(mousePos) {
					currentGameState = Playing
					log.Println("PlayStage is VS knight")
				}
			}

		case Playing:

			win.Clear(colornames.Black)
			basicTxt.Clear()

			//set Enemy
			scaleFactor := 3.0
			scaledSize := pic.Bounds().Size().Scaled(scaleFactor)
			picMonster.Draw(win, pixel.IM.Moved(win.Bounds().Center()).Scaled(win.Bounds().Center(), scaleFactor))
			setEnemyHPBarOut(win, scaledSize)
			setEnemyHPBar(win, scaledSize, enemy.enemyHP, enemy.enemyMaxHP)

			//set Time+rule
			basicTxt.Color = colornames.White
			fmt.Fprintln(basicTxt, "EnemyHP : ", enemy.enemyHP)
			drawPos(win, basicTxt, topCenterPos(win, basicTxt, windowHeightSize))

			question := words[score]
			temp := []byte(question)
			typed := win.Typed()

			basicTxt.Clear()
			basicTxt.Color = colornames.White
			fmt.Fprintln(basicTxt, "> ", words[score])
			drawPos(win, basicTxt, bottleCenterPos(win, basicTxt, windowHeightSize))

			offset := basicTxt.Bounds().W()
			basicTxtOrigX := basicTxt.Dot.X
			spacing := 50.0
			if len(words)-score != 1 {
				basicTxt.Color = colornames.Darkgray
				offset := basicTxt.Bounds().W()
				basicTxt.Clear()
				fmt.Fprintln(basicTxt, words[score+1])
				drawPos(win, basicTxt, bottleCenterPos(win, basicTxt, windowHeightSize).Add(pixel.V(offset+spacing, 0)))
				basicTxt.Dot.X = basicTxtOrigX
			}
			if !(len(words)-score == 2 || len(words)-score == 1) {
				basicTxt.Color = colornames.Gray
				offset += basicTxt.Bounds().W()
				basicTxt.Clear()
				fmt.Fprintln(basicTxt, words[score+2])
				drawPos(win, basicTxt, bottleCenterPos(win, basicTxt, windowHeightSize).Add(pixel.V(offset+spacing*2, 0)))
			}
			//basicTxt.Dot.X = basicTxtOrigX

			basicTxt.Color = colornames.White
			basicTxt.Clear()
			fmt.Fprintln(basicTxt, "\n\n", "collectType = ", collectType, " missType = ", missType)
			drawPos(win, basicTxt, bottleCenterPos(win, basicTxt, windowHeightSize))
			basicTxt.Dot.X = basicTxtOrigX

			//set Time+rule
			basicTxt.Clear()
			basicTxt.Color = colornames.White
			elapsed := time.Since(startTime)
			fmt.Fprintln(basicTxt, "time = ", elapsed.Milliseconds())
			drawPos(win, basicTxt, bottleLeftPos(win, basicTxt, windowHeightSize))

			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					enemy.enemyHP -= 3
					//enemy Down
					if enemy.enemyHP < 0 {
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
				currentGameState = Playing
				collectType, missType, index, score = 0, 0, 0, 0
				enemy.enemyHP = 100
				log.Println("Press:Enter -> GameState:Playing")
			} else if win.JustPressed(pixelgl.KeyTab) {
				currentGameState = StartScreen
				collectType, missType, index, score = 0, 0, 0, 0
				enemy.enemyHP = 100
				log.Println("Press:Enter -> GameState:StartScreen")
			}
		case TestState:
			win.Clear(colornames.Mediumblue)
			//picMonster.Draw(win, pixel.IM)

			basicTxt.Clear()
			tempString = "RightPosition"
			fmt.Fprintln(basicTxt, tempString)
			drawPos(win, basicTxt, topRightPos(win, basicTxt, windowHeightSize))

			basicTxt.Clear()
			tempString = "LeftPosition"
			fmt.Fprintln(basicTxt, tempString)
			drawPos(win, basicTxt, topLeftPos(win, basicTxt, windowHeightSize))

			basicTxt.Clear()
			tempString = "bottleCenterPosition"
			fmt.Fprintln(basicTxt, tempString)
			drawPos(win, basicTxt, bottleCenterPos(win, basicTxt, windowHeightSize))

			basicTxt.Clear()
			tempString = "bottleLeftPosition"
			fmt.Fprintln(basicTxt, tempString)
			drawPos(win, basicTxt, bottleLeftPos(win, basicTxt, windowHeightSize))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
