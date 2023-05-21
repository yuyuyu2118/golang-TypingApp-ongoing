package battle

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/player"
)

var (
	score    = 0
	words    = InitializeQuestion()
	index    = 0
	yourTime = 0.0

	gainGold = 0
	lostGold = 0
)

func BattleTypingV1(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration) myGame.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	if typed != "" {
		if typed[0] == temp[index] && index < len(question) {
			index++
			collectType++
			enemy.HP -= 3
			player.SP += player.BaseSP
			//enemy Down
			log.Println("collectType = ", collectType)
			//1 word end type
			if index == len(question) {
				index = 0
				score++
				//全部打ち切っちゃった場合
				if score == len(words) {
					myGame.CurrentGS = myGame.EndScreen
					yourTime = float64(elapsed.Seconds())
				}
			}
		} else {
			missType++
			log.Println("missType = ", missType)
		}
	}

	BattleTypingSkill(win, player, enemy)
	myGame.CurrentGS = DeathFlug(player, enemy, elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}

func BattleTypingSkill(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		if player.SP == 50 {
			player.SP = 0
			if player.Job == "Warrior" {
				enemy.HP -= 15
			} else if player.Job == "Priest" {
				//TODO 僧侶の回復スキル
			} else if player.Job == "Wizard" {
				//TODO 魔法使いの時止めスキル
			}
		} else {
			log.Println("skillポイントが足りない")
		}
	}
}

func DeathFlug(player *player.PlayerStatus, enemy *enemy.EnemyStatus, elapsed time.Duration, currentGameState myGame.GameState) myGame.GameState {
	if player.HP <= 0 {
		yourTime = float64(elapsed.Seconds())
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		lostGold = rand.Intn(max-min+1) + min
		player.Gold -= lostGold
		log.Println("GameOver!!")
		currentGameState = myGame.EndScreen
	}
	if enemy.HP <= 0 {
		//GoldRandom
		min := int(float64(enemy.Gold) * 0.7)
		max := int(float64(enemy.Gold) * 1.3)
		gainGold = rand.Intn(max-min+1) + min
		player.Gold += gainGold
		index = 0
		score++
		currentGameState = myGame.EndScreen
		yourTime = float64(elapsed.Seconds())
	}
	return currentGameState
}
