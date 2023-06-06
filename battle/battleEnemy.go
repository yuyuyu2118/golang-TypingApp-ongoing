package battle

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myState"
	"golang.org/x/image/colornames"
)

var animationPhase int
var animationInProgress bool

func BattleTypingEnemySlime(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration) myState.GameState {
	if myState.CurrentGS == myState.BattleEnemyScreen {
		if myState.CurrentGS == myState.BattleEnemyScreen && win.JustPressed(pixelgl.KeyEnter) && !animationInProgress {
			animationInProgress = true
			tempEnemySize = enemy.EnemySettings[myGame.StageNum].EnemySize
			tempWordDamage = 0
			SkillTimer = 0.0
<<<<<<< HEAD
<<<<<<< HEAD
			DefenceRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
=======
>>>>>>> 4715f1e (武器の固有スキルを設定)
=======
			DefenceRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
>>>>>>> c9826ea (武器全種のスキル追加)
		}
		if animationInProgress {
			EnemyAttackAnimation(win, player)
		}
	}
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

func EnemyAttackAnimation(win *pixelgl.Window, player *myPlayer.PlayerStatus) {
	switch animationPhase {
	case 0:
		if enemy.EnemySettings[myGame.StageNum].EnemySize < tempEnemySize*1.2 {
			enemy.EnemySettings[myGame.StageNum].EnemySize *= 1.05
		} else {
			win.Canvas().Clear(colornames.Red)
			animationPhase++
		}
	case 1:
		if enemy.EnemySettings[myGame.StageNum].EnemySize > tempEnemySize {
			enemy.EnemySettings[myGame.StageNum].EnemySize *= 0.95
		} else {
			enemy.EnemySettings[myGame.StageNum].EnemySize = tempEnemySize
			if enemy.EnemySettings[myGame.StageNum].OP-player.DP >= 0 {
				//Enemyの攻撃
				player.HP -= enemy.EnemySettings[myGame.StageNum].OP - player.DP
			}
			myState.CurrentGS = myState.PlayingScreen
			index = 0
			animationPhase = 0
			animationInProgress = false
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			player.OP = player.MaxOP
=======
>>>>>>> c9826ea (武器全種のスキル追加)
=======
			player.OP = player.MaxOP
>>>>>>> f10e2c1 (防具のスキル追加完了)
			player.DP = player.MaxDP
			enemy.EnemySettings[myGame.StageNum].OP = enemy.EnemySettings[myGame.StageNum].MaxOP
			RecoveryRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
			SkillTimer += TimerRelationArmorSkill(win, player, &tempWordDamage)
<<<<<<< HEAD
=======
			enemy.EnemySettings[myGame.StageNum].OP = enemy.EnemySettings[myGame.StageNum].MaxOP
>>>>>>> 4715f1e (武器の固有スキルを設定)
=======
>>>>>>> c9826ea (武器全種のスキル追加)
		}
	}
}
