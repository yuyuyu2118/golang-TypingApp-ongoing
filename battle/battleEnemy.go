package battle

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func BattleTypingEnemySlime(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	//eSize := enemy.EnemySettings[myGame.StageNum].EnemySize

	if myGame.CurrentGS == myGame.BattleEnemyScreen {
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
				myGame.CurrentGS = myGame.PlayingScreen
				pressEnter = false
				index = 0
			}
		}
	}
	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}
