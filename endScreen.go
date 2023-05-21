package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/battle"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/player"
)

func battleEndScreen(win *pixelgl.Window, Txt *text.Text, player *player.PlayerStatus, enemy *enemy.EnemyStatus) myGame.GameState {
	if player.HP > 0 {
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
		myPos.LineCenterAlign(win, endLines, Txt, "center")
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
		myPos.LineCenterAlign(win, endLines, Txt, "center")
	}

	//画面遷移,いろいろリセット
	if win.JustPressed(pixelgl.KeyTab) {
		myGame.CurrentGS = myGame.PlayingScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.HP = player.MaxHP
		enemy.HP = enemy.MaxHP
		battle.Shuffle(words)
		log.Println("Press:Enter -> GameState:Playing")
	} else if win.JustPressed(pixelgl.KeyBackspace) {
		myGame.CurrentGS = myGame.GoToScreen
		collectType, missType, index, score = 0, 0, 0, 0
		player.HP = player.MaxHP
		enemy.HP = enemy.MaxHP
		battle.Shuffle(words)
		log.Println("Press:Enter -> GameState:GoToScreen")
	}

	return myGame.CurrentGS
}
