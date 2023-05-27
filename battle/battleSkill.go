package battle

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/yuyuyu2118/typingGo/enemy"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
	"golang.org/x/image/colornames"
)

func BattleTypingSkill(win *pixelgl.Window, player *player.PlayerStatus, enemy *enemy.EnemyStatus) {
	if win.JustPressed(pixelgl.KeySpace) {
		log.Println("Skill!!!")
		if player.SP == 50 {
			index = 0
			player.SP = 0
			myState.CurrentGS = myState.SkillScreen
		} else {
			log.Println("skillポイントが足りない")
		}
	}
}

var (
	RookieSkillCount = 0
	RookieSkillWords = []string{"oreno", "kenngiwo", "kuraeee"}
)

func BattleTypingRookieSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	question := RookieSkillWords[RookieSkillCount]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.AttackTimer // - elapsed.Seconds()

	if myState.CurrentGS == myState.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 4
					if index == len(question) {
						index = 0
						RookieSkillCount++
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						tempWordDamage = 0.0
						if RookieSkillCount == 3 {
							RookieSkillCount = 0
							myState.CurrentGS = myState.PlayingScreen
						}
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.SkillScreen
		}
	}

	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	if myState.CurrentGS == myState.EndScreen {
		RookieSkillCount = 0
	}
	return myState.CurrentGS
}

func InitBattleTextRookieSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := RookieSkillWords[RookieSkillCount]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "俺の")
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Orange
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(RookieSkillWords)-RookieSkillCount != 1 {
			Txt.Color = colornames.Orange
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, "剣技を")
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+1])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(RookieSkillWords)-RookieSkillCount == 2 || len(RookieSkillWords)-RookieSkillCount == 1) {
			Txt.Color = colornames.Orange
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, "くらえええ!!")
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+2])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var bulletLoadingSkill = []bool{false, false, false, false, false}
var bulletDamageSkill = []int{0, 0, 0, 0, 0}

func BattleTypingHunterSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	xOffSet := 50.0
	yOffSet := myPos.TopLefPos(win, myUtil.ScreenTxt).Y - 100
	txtPos := pixel.V(xOffSet, yOffSet)
	myUtil.ScreenTxt.Color = colornames.White
	myUtil.HunterBulletTxt.Clear()
	myUtil.HunterBulletTxt.Color = colornames.White
	fmt.Fprintln(myUtil.HunterBulletTxt, "*拡張装填*")
	txtPos = pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.HunterBulletTxt.Draw(win, tempPosition)

	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.AttackTimer // - elapsed.Seconds()

	if myState.CurrentGS == myState.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= 3
					//PlayerAttack(30, pixel.Vec{X: 0, Y: 0})
					if index == len(question) {
						index = 0
						score++
						//enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						//PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						//tempWordDamage = 0.0
						if bulletLoadingSkill[3] {
							bulletLoadingSkill[4] = true
						} else if bulletLoadingSkill[2] {
							bulletLoadingSkill[3] = true
						} else if bulletLoadingSkill[1] {
							bulletLoadingSkill[2] = true
						} else if bulletLoadingSkill[0] {
							bulletLoadingSkill[1] = true
						}
						bulletLoadingSkill[0] = true

						if bulletLoadingSkill[0] {
							bulletDamageSkill[0] = int(tempWordDamage)
						}
						if bulletLoadingSkill[1] {
							bulletDamageSkill[1] = int(tempWordDamage)
						}
						if bulletLoadingSkill[2] {
							bulletDamageSkill[2] = int(tempWordDamage)
						}
						if bulletLoadingSkill[3] {
							bulletDamageSkill[3] = int(tempWordDamage)
						}
						if bulletLoadingSkill[4] {
							bulletDamageSkill[4] = int(tempWordDamage)
						}
						tempWordDamage = 0.0
						log.Println(bulletLoadingSkill)
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.SkillScreen
		}
	}

	if bulletLoadingSkill[0] && !bulletLoadingSkill[1] && !bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && !bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && !bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && bulletLoadingSkill[3] && !bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[3])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-4], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	} else if bulletLoadingSkill[0] && bulletLoadingSkill[1] && bulletLoadingSkill[2] && bulletLoadingSkill[3] && bulletLoadingSkill[4] {
		myUtil.HunterBulletTxt.Clear()
		myUtil.HunterBulletTxt.Color = colornames.White
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-1], bulletDamageSkill[4])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-2], bulletDamageSkill[3])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-3], bulletDamageSkill[2])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-4], bulletDamageSkill[1])
		fmt.Fprintln(myUtil.HunterBulletTxt, words[score-5], bulletDamageSkill[0])
		yOffSet -= myUtil.HunterBulletTxt.LineHeight + 30
		txtPos = pixel.V(xOffSet, yOffSet)
		tempPosition := pixel.IM.Moved(txtPos)
		myUtil.HunterBulletTxt.Draw(win, tempPosition)
	}

	if win.JustPressed(pixelgl.KeyEnter) {
		bulletDamageSkills := bulletDamageSkill[0] + bulletDamageSkill[1] + bulletDamageSkill[2] + bulletDamageSkill[3] + bulletDamageSkill[4]
		//enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamageSkills) //TODO: debug用
		PlayerAttack(win, bulletDamageSkill[0], win.Bounds().Center().Sub(pixel.V(50, -200)))
		PlayerAttack(win, bulletDamageSkill[1], win.Bounds().Center().Sub(pixel.V(-100, -200)))
		PlayerAttack(win, bulletDamageSkill[2], win.Bounds().Center().Sub(pixel.V(200, -200)))
		PlayerAttack(win, bulletDamageSkill[3], win.Bounds().Center().Sub(pixel.V(-200, -200)))
		PlayerAttack(win, bulletDamageSkill[4], win.Bounds().Center().Sub(pixel.V(300, -200)))
		enemy.EnemySettings[myGame.StageNum].HP += float64(bulletDamageSkills)
		for i := 0; i < 5; i++ {
			bulletDamageSkill[i] = 0
			bulletLoadingSkill[i] = false
		}
		log.Println("Bang!!")
		myState.CurrentGS = myState.PlayingScreen
	}

	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	return myState.CurrentGS
}

func InitBattleTextHunterSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := words[score]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, wordsJapanese[words[score]])
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(words)-score != 1 {
			Txt.Color = colornames.Darkgray
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[score+1]])
			fmt.Fprintln(Txt, words[score+1])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(words)-score == 2 || len(words)-score == 1) {
			Txt.Color = colornames.Gray
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[score+2]])
			fmt.Fprintln(Txt, words[score+2])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "EnemyAttack!!")
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var (
	MonkSkillWords = []string{
		"dadadadadadadadadadadadadada!!!!!!!",
		"mudamudamudamudamudamudamuda!!!!!!!",
		"oraoraoraoraoraoraoraoraoraora!!!!!",
	}
	MonkSkillWord = ""
)

func BattleTypingMonkSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	if MonkSkillWord == "" {
		MonkSkillWord = MonkSkillWords[rand.Intn(3)]
	}
	question := MonkSkillWord
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.AttackTimer // - elapsed.Seconds()

	if myState.CurrentGS == myState.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					tempWordDamage -= float64(rand.Intn(3))
					enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
					PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
					tempWordDamage = 0.0
					if index == len(question) {
						index = 0
						MonkSkillWord = ""
						myState.CurrentGS = myState.PlayingScreen
					}
				} else {
					missType++
				}
			}
		} else {
			myState.CurrentGS = myState.SkillScreen
		}
	}

	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	if myState.CurrentGS == myState.EndScreen {
		//index?
		MonkSkillWord = ""
	}
	return myState.CurrentGS
}

func InitBattleTextMonkSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := MonkSkillWord
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Gray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Red
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var magicHP int
var setTime float64
var timeBool bool

func BattleTypingMagicUserSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	question := words[score]
	temp := []byte(question)
	typed := win.Typed()

	if !timeBool {
		setTime = 0.0
		timeBool = true
	}

	setTime += 0.02
	tempCount = 5.0 - setTime
	log.Println(tempCount, setTime)

	if myState.CurrentGS == myState.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					collectType++
					magicHP++
					if index == len(question) {
						index = 0
						score++
					}
				} else {
					missType++
					magicHP--
				}
			}
		} else {
			myState.CurrentGS = myState.PlayingScreen
		}
	}

	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	if myState.CurrentGS == myState.PlayingScreen {
		if magicHP < 0 {
			magicHP = 0
		}
		player.HP += float64(magicHP)
		magicHP = 0
		if player.HP >= player.MaxHP {
			player.HP = player.MaxHP
		}
		timeBool = false
	}
	return myState.CurrentGS
}

func InitBattleTextMagicUserSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "+HP:", magicHP)
	myPos.DrawPos(win, Txt, myPos.CenLefPos(win, Txt))

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := words[score]
		Txt.Clear()
		Txt.Color = colornames.Lightgreen
		fmt.Fprintln(Txt, wordsJapanese[words[score]])
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Green
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Lightgreen
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(words)-score != 1 {
			Txt.Color = colornames.Lightgreen
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[score+1]])
			fmt.Fprintln(Txt, words[score+1])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(words)-score == 2 || len(words)-score == 1) {
			Txt.Color = colornames.Lightgreen
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[score+2]])
			fmt.Fprintln(Txt, words[score+2])
			myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	}
	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

var (
	MonsterSkillWords = []string{
		"namamuginamagomenamatamago",
		"akamakigamiaomakigamikimakigami",
		"niwanihaniwaniwatorigairu",
	}
	MonsterSkillWord = ""
)

func BattleTypingMonsterSkill(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) myState.GameState {
	if MonsterSkillWord == "" {
		MonsterSkillWord = MonsterSkillWords[rand.Intn(3)]
	}
	question := MonsterSkillWord
	temp := []byte(question)
	typed := win.Typed()

	tempCount = player.AttackTimer // - elapsed.Seconds()

	if myState.CurrentGS == myState.SkillScreen {
		if tempCount > 0 {
			if typed != "" {
				if typed[0] == temp[index] && index < len(question) {
					index++
					indexMonster += 2
					collectType++
					tempWordDamage -= 4
					if index == len(question) {
						index = 0
						indexMonster = 0
						MonsterSkillWord = ""
						enemy.EnemySettings[myGame.StageNum].HP += tempWordDamage
						PlayerAttack(win, int(tempWordDamage), win.Bounds().Center().Sub(pixel.V(50, 150)))
						tempWordDamage = 0.0
						myState.CurrentGS = myState.PlayingScreen
					}
				} else {
					missType++
					player.HP--
				}
			}
		} else {
			myState.CurrentGS = myState.SkillScreen
			indexMonster = 0
		}
	}

	myState.CurrentGS = DeathFlug(player, &enemy.EnemySettings[myGame.StageNum], elapsed, myState.CurrentGS)
	if myState.CurrentGS == myState.EndScreen {
		//index?
		MonsterSkillWord = ""
	}
	return myState.CurrentGS
}

func InitBattleTextMonsterSkill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := MonsterSkillWord
		tempWordsSlice := strings.Split(tempWords, "")
		tempWords = strings.Join(tempWordsSlice, " ")
		Txt.Clear()
		Txt.Color = colornames.Gray
		log.Println(indexMonster)
		fmt.Fprint(Txt, tempWords[:indexMonster])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[indexMonster:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}
