package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var saveContent string

var (
	save1Button = pixel.Rect{}
	save2Button = pixel.Rect{}
)

func initSave(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "Do you want to save?")
	tempPosition = topCenterPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "1. Yes")
	tempPosition = centerLeftPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	save1Button = Txt.Bounds().Moved(tempPosition)

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "2. No")
	tempPosition = centerPos(win, Txt, windowHeightSize)
	drawPos(win, Txt, tempPosition)
	save2Button = Txt.Bounds().Moved(tempPosition)
}

func saveClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState, player *PlayerStatus) GameState {
	//TODO ページを作成したら追加
	if save1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		saveGame(player)
		currentGameState = GoToScreen
		log.Println("Save Done!")
	} else if save1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = GoToScreen
		log.Println("saveScreen->GoToScreen")
	}
	return currentGameState
}

func saveGame(player *PlayerStatus) {
	filename := "assets\\save\\save.csv"
	initialText := "- INITIAL TEXT -\nThis file stores game save data.\nGold,job,equipment,\n"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// ファイルが空の場合は初期テキストを書き込む
	fileInfo, err := file.Stat()
	log.Println(fileInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if fileInfo.Size() == 0 {
		_, err = file.WriteString(initialText)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//TODO 100行までセーブする100回からは古いデータから消える
	saveContent = strconv.Itoa(player.playerGold) + "," + player.playerJob

	_, err = file.WriteString(saveContent + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}
