package battle

import (
	"fmt"
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func BattleTypingRookie(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP - elapsed.Seconds()

	if myGame.CurrentGS == myGame.PlayingScreen {
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
			myGame.CurrentGS = myGame.BattleEnemyScreen
		}
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}

var bulletLoading = []bool{false, false, false}

func BattleTypingHunter(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myGame.GameState {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(0, 0)
	myUtil.ScreenTxt.Color = colornames.White

	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.OP - elapsed.Seconds()

	if myGame.CurrentGS == myGame.PlayingScreen {
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
						if bulletLoading[1] {
							bulletLoading[2] = true
						} else if bulletLoading[0] {
							bulletLoading[1] = true
						}
						bulletLoading[0] = true
						log.Println(bulletLoading)
					}
				} else {
					missType++
				}
			}
		} else {
			myGame.CurrentGS = myGame.BattleEnemyScreen
		}
	}

	if bulletLoading[0] && !bulletLoading[1] && !bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && !bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myGame.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myGame.CurrentGS)
	return myGame.CurrentGS
}
