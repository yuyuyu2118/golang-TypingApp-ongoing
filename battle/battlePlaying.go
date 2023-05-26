package battle

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func BattleTypingRookie(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
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
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
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
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

var bulletLoading = []bool{false, false, false}
var bulletDamage = []int{0, 0, 0}

func BattleTypingHunter(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(0, 0)
	myUtil.ScreenTxt.Color = colornames.White

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
					tempWordDamage -= 2
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						score++
						//enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						//PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						//tempWordDamage = 0.0
						if bulletLoading[1] {
							bulletLoading[2] = true
						} else if bulletLoading[0] {
							bulletLoading[1] = true
						}
						bulletLoading[0] = true

						if bulletLoading[0] {
							bulletDamage[0] = int(tempWordDamage)
						}
						if bulletLoading[1] {
							bulletDamage[1] = int(tempWordDamage)
						}
						if bulletLoading[2] {
							bulletDamage[2] = int(tempWordDamage)
						}
						tempWordDamage = 0.0
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.BattleEnemyScreen
		}
	}

	if bulletLoading[0] && !bulletLoading[1] && !bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && !bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamage[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamage[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamage[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	}

	if win.JustPressed(pixelgl.KeyEnter) {
		bulletDamages := bulletDamage[0] + bulletDamage[1] + bulletDamage[2]
		//enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamages) //TODO: debug用
		PlayerAttack(win, bulletDamage[0], win.Bounds().Center().Sub(pixel.V(50, -200)))
		PlayerAttack(win, bulletDamage[1], win.Bounds().Center().Sub(pixel.V(-100, -200)))
		PlayerAttack(win, bulletDamage[2], win.Bounds().Center().Sub(pixel.V(200, -200)))
		enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamages)
		for i := 0; i < 3; i++ {
			bulletDamage[i] = 0
			bulletLoading[i] = false
		}
		log.Println("Bang!!")
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

func BattleTypingMonk(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
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
					tempWordDamage -= 3 + float64(rand.Intn(3)-1)
					enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
					PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, -250)))
					tempWordDamage = 0.0
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						score++
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.BattleEnemyScreen
		}
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}
