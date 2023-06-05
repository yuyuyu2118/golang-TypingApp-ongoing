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
)

var (
	wordsNum                = 0
	wordsJapanese, words, _ = LoadWordsFromCSV("assets\\question\\question2_6.csv")
	index                   = 0
	yourTime                = 0.0

	//別の関数で呼び出す用
	gainGold            = 0
	gainGoldCollectType = 0
	lostGold            = 0
	lostGoldMissType    = 0

	AttackCount    = 3.0
	tempCount      = 0.0
	tempWordDamage = 0.0
	tempEnemySize  = 0.0
)

func DeathFlug(player *player.PlayerStatus, enemyInf *enemy.EnemyStatus, elapsed time.Duration, currentGameState myState.GameState) myState.GameState {
	//自分がやられたとき
	if player.HP <= 0 {
		yourTime = float64(elapsed.Seconds())

		lostGoldMin := int(float64(enemyInf.Gold) * 0.7)
		lostGoldMax := int(float64(enemyInf.Gold) * 1.3)
		lostGold = rand.Intn(lostGoldMax-lostGoldMin+1) + lostGoldMin
		lostGoldMissType = int(float64(missType) * 0.1)

		player.Gold -= (lostGold + lostGoldMissType)

		log.Println("DeathFlug: GameOver")
		currentGameState = myState.EndScreen
	}
	//敵を倒したとき
	if enemyInf.HP <= 0 {

		gainGoldMin := int(float64(enemyInf.Gold) * 0.7)
		gainGoldMax := int(float64(enemyInf.Gold) * 1.3)
		gainGold = rand.Intn(gainGoldMax-gainGoldMin+1) + gainGoldMin
		gainGoldCollectType = int(float64(collectType) * 0.3)

		player.Gold += gainGold + gainGoldCollectType

		//AbilityPointの付与
		player.AP += enemyInf.DropAP

		index = 0
		wordsNum++

		log.Println("DeathFlug: DefeatEnemy")
		currentGameState = myState.EndScreen

		yourTime = float64(elapsed.Seconds())
		for _, name := range enemy.EnemyNameSlice {
			if enemyInf.Name == name {
				myGame.SaveDefeatedEnemyEvent(myGame.SaveFilePath, 2, name)
			}
		}
	}

	return currentGameState
}

func BattleTypingTest(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.AttackTimer - elapsed.Seconds()

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				// if (temp[index] == 'j' && index < len(question)) && typed[0] == 'z' {
				// 	log.Println("koko", temp[index])
				// 	question = strings.NewReplacer("j", "").Replace(words[wordsNum])
				// 	temp = []byte(question)
				// }
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 2
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						wordsNum++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, tempWordDamage, win.Bounds().Center().Sub(pixel.V(50, 150)))
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
		if myState.CurrentGS == myState.BattleEnemyScreen && win.JustPressed(pixelgl.KeyEnter) && !animationInProgress {
			animationInProgress = true
			tempEnemySize = enemy.EnemySettings[myGame.StageNum].EnemySize
			tempWordDamage = 0
		}
		if animationInProgress {
			EnemyAttackAnimation(win, player)
		}
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}
