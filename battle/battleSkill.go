package battle

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/player"
)

func BattleTypingSkill(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		if player.SP == 50 {
			index = 0
			player.SP = 0
			myGame.CurrentGS = myGame.SkillScreen
		} else {
			log.Println("skillポイントが足りない")
		}
	}
}

var (
	RookieSkillCount = 0
	RookieSkillWords = []string{"oreno", "kenngiwo", "kuraeee"}
)

func BattleTypingRookieSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	question := RookieSkillWords[RookieSkillCount]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP // - elapsed.Seconds()

	if myGame.CurrentGS == myGame.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 5
					if index == len(question) {
						index = 0
						RookieSkillCount++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage - 1 //TODO: debug用
						PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						tempWordDamage = 0.0
						if RookieSkillCount == 3 {
							RookieSkillCount = 0
							myGame.CurrentGS = myGame.PlayingScreen
						}
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.SkillScreen
		}
	}

	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}
