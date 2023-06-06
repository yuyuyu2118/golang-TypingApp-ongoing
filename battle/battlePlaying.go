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
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

var SkillTimer float64

<<<<<<< HEAD
func BattleTypingRookie(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
=======
func BattleTypingRookie(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
>>>>>>> 4715f1e (武器の固有スキルを設定)
	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	*tempTimer = player.AttackTimer - elapsed.Seconds() + SkillTimer
	tempCount = *tempTimer

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= player.OP
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> c9826ea (武器全種のスキル追加)
						AttackRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						RecoveryRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						SkillTimer += TimerRelationWeaponSkill(win, player, tempTimer)
						DebuffRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
						AttackRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
<<<<<<< HEAD
<<<<<<< HEAD
						BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
=======
						AttackRelationSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						RecoveryRelationSkill(win, player, &tempWordDamage)
						SkillTimer += TimerRelationSkill(win, player, tempTimer)
						DebuffRelationSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
>>>>>>> 4715f1e (武器の固有スキルを設定)
=======
>>>>>>> c9826ea (武器全種のスキル追加)
=======
						BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
>>>>>>> f10e2c1 (防具のスキル追加完了)
						index = 0
						wordsNum++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, tempWordDamage, win.Bounds().Center().Sub(pixel.V(-200, -200)))
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
var bulletDamage = []float64{0.0, 0.0, 0.0}

<<<<<<< HEAD
func BattleTypingHunter(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
=======
func BattleTypingHunter(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
>>>>>>> 124dadb (スキル調整)
	xOffSet := 100.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(0, 0)
	myUtil.ScreenTxt.Color = colornames.White

	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	*tempTimer = player.AttackTimer - elapsed.Seconds() + SkillTimer
	tempCount = *tempTimer

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= player.OP + (float64(-rand.Intn(2) - 1))
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						wordsNum++

						AttackRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						RecoveryRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						SkillTimer += TimerRelationWeaponSkill(win, player, tempTimer)
						DebuffRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
						AttackRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
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
							bulletDamage[0] = tempWordDamage
						}
						if bulletLoading[1] {
							bulletDamage[1] = tempWordDamage
						}
						if bulletLoading[2] {
							bulletDamage[2] = tempWordDamage
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
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-1], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && !bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-1], bulletDamage[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-2], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoading[0] && bulletLoading[1] && bulletLoading[2] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-1], bulletDamage[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-2], bulletDamage[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[wordsNum-3], bulletDamage[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	}

	if win.JustPressed(pixelgl.KeyEnter) {
		bulletDamages := bulletDamage[0] + bulletDamage[1] + bulletDamage[2]
		//enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamages) //TODO: debug用
		PlayerAttack(win, bulletDamage[0], win.Bounds().Center().Sub(pixel.V(-200, -150)))
		PlayerAttack(win, bulletDamage[1], win.Bounds().Center().Sub(pixel.V(-200, -200)))
		PlayerAttack(win, bulletDamage[2], win.Bounds().Center().Sub(pixel.V(-200, -250)))
		enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamages)
		for i := 0; i < 3; i++ {
			bulletDamage[i] = 0
			bulletLoading[i] = false
		}
		//TODO: テキスト表示
		log.Println("Bang!!")
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

<<<<<<< HEAD
func BattleTypingMonk(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
=======
func BattleTypingMonk(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
>>>>>>> 124dadb (スキル調整)
	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	*tempTimer = player.AttackTimer - elapsed.Seconds() + SkillTimer
	tempCount = *tempTimer

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= player.OP + float64(-rand.Intn(3))

					AttackRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
					RecoveryRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
					SkillTimer += TimerRelationWeaponSkill(win, player, tempTimer)
					DebuffRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
					AttackRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
					BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])

					enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
					randomValue := (-1 * (100 + (rand.Float64() * 200)))
					PlayerAttack(win, tempWordDamage, win.Bounds().Center().Sub(pixel.V(-200, randomValue)))
					tempWordDamage = 0.0
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					player.SP += player.BaseSP
					if index == len(question) {
						index = 0
						wordsNum++
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

var magicCollectType = 0.0
var magicMissType = 0

<<<<<<< HEAD
func BattleTypingMagicUser(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
=======
func BattleTypingMagicUser(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
>>>>>>> 124dadb (スキル調整)
	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	*tempTimer = player.AttackTimer - elapsed.Seconds() + SkillTimer
	tempCount = *tempTimer

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++

					if magicCollectType >= 0 && magicCollectType < 25 {
						magicCollectType += player.OP * 0.25
					} else if magicCollectType >= 25 && magicCollectType < 50 {
						magicCollectType += player.OP * 0.5
					} else if magicCollectType >= 50 && magicCollectType < 100 {
						magicCollectType += player.OP * 1
					} else if magicCollectType >= 100 {
						magicCollectType += player.OP * 1.5
					} else if magicCollectType >= 500 {
						magicCollectType += player.OP * 3.0
					} else if magicCollectType >= 1000 {
						magicCollectType += player.OP * 15.0
					} else if magicCollectType >= 2500 {
						magicCollectType += player.OP * 30.0
					}
					log.Println(magicCollectType)
					player.SP += player.BaseSP
					if index == len(question) {

						tempWordDamage = -magicCollectType * 0.3
						AttackRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						RecoveryRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						SkillTimer += TimerRelationWeaponSkill(win, player, tempTimer)
						DebuffRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
						AttackRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])

						index = 0
						wordsNum++
						//PlayerAttack(win, tempWordDamage, win.Bounds().Center().Sub(pixel.V(50, 150)))
					}
				} else {
					missType++
					magicMissType++
					if magicCollectType >= 0 && magicCollectType < 25 {
						magicCollectType -= player.OP * 3.0
					} else if magicCollectType >= 25 && magicCollectType < 50 {
						magicCollectType -= player.OP * 5.0
					} else if magicCollectType >= 50 && magicCollectType < 100 {
						magicCollectType -= player.OP * 10.0
					} else if magicCollectType >= 100 {
						magicCollectType -= player.OP * 30.0
					} else if magicCollectType >= 500 {
						magicCollectType -= player.OP * 50.0
					} else if magicCollectType >= 1000 {
						magicCollectType -= player.OP * 100.0
					} else if magicCollectType >= 2500 {
						magicCollectType -= player.OP * 500.0
					}
					if magicCollectType < 0 {
						magicCollectType = 0
					}
				}
			}
		} else {
			myState.CurrentGS = myState.BattleEnemyScreen
		}
	}

	if win.JustPressed(pixelgl.KeyEnter) {
		PlayerAttack(win, -magicCollectType, win.Bounds().Center().Sub(pixel.V(-200, -200)))
		enemy.EnemySettings[myGame.StageNum].HP += float64(-magicCollectType)
		magicCollectType = 0
	}

	BattleTypingSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

<<<<<<< HEAD
func BattleTypingMonster(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
=======
func BattleTypingMonster(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration, tempTimer *float64) myState.GameState {
>>>>>>> 124dadb (スキル調整)
	question := words[wordsNum]
	temp := []byte(question)
	typed := win.Typed()

	*tempTimer = player.AttackTimer - elapsed.Seconds() + SkillTimer
	tempCount = *tempTimer

	if myState.CurrentGS == myState.PlayingScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= float64(rand.Intn(int(player.OP * 3.0)) /* - rand.Intn(int(player.OP))*/)
					player.SP += player.BaseSP
					if index == len(question) {

						AttackRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						RecoveryRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						SkillTimer += TimerRelationWeaponSkill(win, player, tempTimer)
						DebuffRelationWeaponSkill(win, player, &enemy.EnemySettings[myGame.StageNum])
						AttackRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum], &tempWordDamage)
						BuffRelationArmorSkill(win, player, &enemy.EnemySettings[myGame.StageNum])

						index = 0
						wordsNum++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, tempWordDamage, win.Bounds().Center().Sub(pixel.V(-200, -200)))
						tempWordDamage = 0
						if rand.Float64() <= 0.1 { // 10%の確率で自分に攻撃
							player.HP--
							player.SP--
							win.Canvas().Clear(colornames.Red)
						}
					}
				} else {
					missType++
					if rand.Float64() <= 0.3 { // 30%の確率で自分に攻撃
						player.HP--
						win.Canvas().Clear(colornames.Red)
					}
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
