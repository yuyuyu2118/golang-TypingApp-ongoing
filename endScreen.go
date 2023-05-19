package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func battleEndScreen(win *pixelgl.Window, Txt *text.Text, player *PlayerStatus, enemy *EnemyStatus) GameState {
	if player.playerHP > 0 {
		yourTimeString := fmt.Sprintf("%0.3f", yourTime)
		//平均キータイプ数 回/秒 Escでもう一度,Tabでタイトル
		endLines := []string{
			"YourScore : " + strconv.Itoa(score),
			"\n",
			"yourTime =" + yourTimeString,
			"collectType = " + strconv.Itoa(collectType) + " missType = " + strconv.Itoa(missType),
			"\n\n",
			"ReSTART : Press Tab | Back : Press BackSpace",
		}
		lineCenterAlign(win, endLines, Txt, "center")
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
			"ReSTART : Press Tab | Back : Press BackSpace",
		}
		lineCenterAlign(win, endLines, Txt, "center")
	}

	//画面遷移,いろいろリセット
	if win.JustPressed(pixelgl.KeyTab) {
		currentGameState = PlayingScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.playerHP = player.playerMaxHP
		enemy.enemyHP = enemy.enemyMaxHP
		shuffle(words)
		log.Println("Press:Enter -> GameState:Playing")
	} else if win.JustPressed(pixelgl.KeyBackspace) {
		currentGameState = GoToScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.playerHP = player.playerMaxHP
		enemy.enemyHP = enemy.enemyMaxHP
		shuffle(words)
		log.Println("Press:Enter -> GameState:GoToScreen")
	}

	return currentGameState
}
