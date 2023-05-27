package battle

import (
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

var (
	score = 0
	//words  = InitializeQuestion()
	wordsJapanese, words, _ = LoadWordsFromCSV()
	index                   = 0
	yourTime                = 0.0

	gainGold = 0
	lostGold = 0

	AttackCount = 3.0
	tempCount   = 0.0
	lock        = false
	lock2       = false
	pressEnter  = false

	tempWordDamage = 0.0
	tempEnemySize  = 0.0
)

func DeathFlug(player *player.PlayerStatus, enemyInf *enemy.EnemyStatus, elapsed time.Duration, currentGameState myState.GameState) myState.GameState {
	if player.HP <= 0 {
		yourTime = float64(elapsed.Seconds())
		min := int(float64(enemyInf.Gold) * 0.7)
		max := int(float64(enemyInf.Gold) * 1.3)
		lostGold = rand.Intn(max-min+1) + min
		player.Gold -= lostGold
		log.Println("GameOver!!")
		currentGameState = myState.EndScreen
	}
	if enemyInf.HP <= 0 {
		//GoldRandom
		min := int(float64(enemyInf.Gold) * 0.7)
		max := int(float64(enemyInf.Gold) * 1.3)
		gainGold = rand.Intn(max-min+1) + min
		player.Gold += gainGold
		//AbilityPointの付与
		player.AP += enemyInf.DropAP
		index = 0
		score++
		currentGameState = myState.EndScreen
		yourTime = float64(elapsed.Seconds())
		for _, name := range enemy.EnemyNameSlice {
			if enemyInf.Name == name {
				myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, name)
			}
		}
		// if enemyInf.Name == "Slime" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Slime")
		// } else if enemyInf.Name == "Bird" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Bird")
		// } else if enemyInf.Name == "Plant" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Plant")
		// } else if enemyInf.Name == "Goblin" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Goblin")
		// } else if enemyInf.Name == "Zombie" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Zombie")
		// } else if enemyInf.Name == "Fairy" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Fairy")
		// } else if enemyInf.Name == "Skull" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Skull")
		// } else if enemyInf.Name == "Wizard" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Wizard")
		// } else if enemyInf.Name == "Solidier" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Solidier")
		// } else if enemyInf.Name == "Dragon" {
		// 	myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, "Dragon")
		// }
	}
	return currentGameState
}

func BattleTyping(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP - elapsed.Seconds()

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 3
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						score++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage - 1 //TODO: debug用
						PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						tempWordDamage = 0.0
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.BattleEnemyScreen
		}
	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		//攻撃判定
		//PressEnter
		if win.JustPressed(pixelgl.KeyEnter) {
			pressEnter = true
			tempEnemySize = enemy.EnemySettings[myGame.StageNum].EnemySize
			tempWordDamage = 0
		}
		if pressEnter == true {
			if enemy.EnemySettings[myGame.StageNum].EnemySize >= tempEnemySize && enemy.EnemySettings[myGame.StageNum].EnemySize < tempEnemySize*1.2 && lock == false && lock2 == false {
				enemy.EnemySettings[myGame.StageNum].EnemySize = enemy.EnemySettings[myGame.StageNum].EnemySize * 1.05
				if enemy.EnemySettings[myGame.StageNum].EnemySize > tempEnemySize*1.2 {
					lock = true
					win.Canvas().Clear(colornames.Red)
				}
			} else if enemy.EnemySettings[myGame.StageNum].EnemySize >= tempEnemySize && lock == true && lock2 == false {
				enemy.EnemySettings[myGame.StageNum].EnemySize = enemy.EnemySettings[myGame.StageNum].EnemySize * 0.95
				if enemy.EnemySettings[myGame.StageNum].EnemySize < tempEnemySize {
					lock2 = true
				}
			} else if lock == true && lock2 == true {
				enemy.EnemySettings[myGame.StageNum].EnemySize = tempEnemySize
				lock = false
				lock2 = false
				player.HP -= enemy.EnemySettings[myGame.StageNum].OP
				myState.CurrentGS = myState.PlayingScreen
				pressEnter = false
				index = 0
			}
		}
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}
