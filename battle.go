package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
)

var (
	score    = 0
	words    = initializeQuestion()
	index    = 0
	yourTime = 0.0

	gainGold = 0
	lostGold = 0
)

func battleTypingV1(win *pixelgl.Window, player *PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration) GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	if typed != "" {
		if typed[0] == temp[index] && index < len(question) {
			index++
			collectType++
			enemy.HP -= 3
			player.playerSP += player.playerBaseSP
			//enemy Down
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

	battleTypingSkill(win, player, enemy)
	currentGameState = deathFlug(player, enemy, elapsed, currentGameState)
	return currentGameState
}

func battleTypingSkill(win *pixelgl.Window, player *PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		player.playerSP = 0
		if player.playerJob == "Warrior" {
			enemy.HP -= 15
		} else if player.playerJob == "Priest" {
			//TODO 僧侶の回復スキル
		} else if player.playerJob == "Wizard" {
			//TODO 魔法使いの時止めスキル
		}
	}
}

func deathFlug(player *PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration, currentGameState GameState) GameState {
	if player.playerHP <= 0 {
		yourTime = float64(elapsed.Seconds())
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		lostGold = rand.Intn(max-min+1) + min
		player.playerGold -= lostGold
		log.Println("GameOver!!")
		currentGameState = EndScreen
	}
	if enemy.HP <= 0 {
		//GoldRandom
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		gainGold = rand.Intn(max-min+1) + min
		player.playerGold += gainGold
		index = 0
		score++
		currentGameState = EndScreen
		yourTime = float64(elapsed.Seconds())
	}
	return currentGameState
}
