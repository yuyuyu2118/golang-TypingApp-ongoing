package battle

import (
	"fmt"
	"strconv"
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
	"golang.org/x/image/colornames"
)

var (
	collectType = 0
	missType    = 0
)

func InitBattleTextV1(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "> ", words[wordsNum])
	myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

	offset := Txt.Bounds().W()
	TxtOrigX := Txt.Dot.X
	spacing := 100.0
	if len(words)-wordsNum != 1 {
		Txt.Color = colornames.Darkgray
		offset := Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[wordsNum+1])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
		Txt.Dot.X = TxtOrigX
	}
	if !(len(words)-wordsNum == 2 || len(words)-wordsNum == 1) {
		Txt.Color = colornames.Gray
		offset += Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[wordsNum+2])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
	}
	return elapsed
}

func InitBattleTextV2(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.PlayingScreen {
		tempWords := words[wordsNum]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, wordsJapanese[words[wordsNum]])
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(words)-wordsNum != 1 {
			Txt.Color = colornames.Darkgray
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[wordsNum+1]])
			fmt.Fprintln(Txt, words[wordsNum+1])
			myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(words)-wordsNum == 2 || len(words)-wordsNum == 1) {
			Txt.Color = colornames.Gray
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[wordsNum+2]])
			fmt.Fprintln(Txt, words[wordsNum+2])
			myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "敵の通常攻撃!!!")
		fmt.Fprint(Txt, "攻撃力:", enemy.EnemySettings[myGame.StageNum].OP, "防御力:", enemy.EnemySettings[myGame.StageNum].DP)
		myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

func InitBattleTextV2Skill(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.SkillScreen {
		tempWords := RookieSkillWords[RookieSkillCount]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprint(Txt, "> ")
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.Orange
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(RookieSkillWords)-RookieSkillCount != 1 {
			Txt.Color = colornames.Orange
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+1])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(RookieSkillWords)-RookieSkillCount == 2 || len(RookieSkillWords)-RookieSkillCount == 1) {
			Txt.Color = colornames.Orange
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, RookieSkillWords[RookieSkillCount+2])
			myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

func InitBattleTextMagicUser(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.PlayingScreen {
		if magicCollectType >= 0 && magicCollectType < 25 {
			myUtil.Magic1Txt.Clear()
			myUtil.Magic1Txt.Color = colornames.White
			fmt.Fprintln(myUtil.Magic1Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic1Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic1Txt.TabWidth/2-50, myUtil.Magic1Txt.LineHeight/2-200)))
		} else if magicCollectType >= 25 && magicCollectType < 50 {
			myUtil.Magic2Txt.Clear()
			myUtil.Magic2Txt.Color = colornames.Yellow
			fmt.Fprintln(myUtil.Magic2Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic2Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic2Txt.TabWidth/2-50, myUtil.Magic2Txt.LineHeight/2-200)))
		} else if magicCollectType >= 50 && magicCollectType < 100 {
			myUtil.Magic3Txt.Clear()
			myUtil.Magic3Txt.Color = colornames.Orange
			fmt.Fprintln(myUtil.Magic3Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic3Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic3Txt.TabWidth/2-50, myUtil.Magic3Txt.LineHeight/2-200)))
		} else if magicCollectType >= 100 && magicCollectType < 500 {
			myUtil.Magic4Txt.Clear()
			myUtil.Magic4Txt.Color = colornames.Red
			fmt.Fprintln(myUtil.Magic4Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic4Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic4Txt.TabWidth/2-50, myUtil.Magic4Txt.LineHeight/2-200)))
		} else if magicCollectType >= 500 && magicCollectType < 1000 {
			myUtil.Magic5Txt.Clear()
			myUtil.Magic5Txt.Color = colornames.Darkred
			fmt.Fprintln(myUtil.Magic5Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic5Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic5Txt.TabWidth/2-50, myUtil.Magic5Txt.LineHeight/2-200)))
		} else if magicCollectType >= 1000 {
			myUtil.Magic6Txt.Clear()
			myUtil.Magic6Txt.Color = colornames.Purple
			fmt.Fprintln(myUtil.Magic6Txt, "魔法力:", strconv.FormatFloat(magicCollectType, 'f', 2, 64))
			myPos.DrawPos(win, myUtil.Magic6Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(myUtil.Magic6Txt.TabWidth/2-50, myUtil.Magic6Txt.LineHeight/2-200)))
		}

		tempWords := words[wordsNum]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, wordsJapanese[words[wordsNum]])
		Txt.Color = colornames.Darkslategray
		fmt.Fprint(Txt, tempWords[:index])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[index:])
		myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt))

		offset := Txt.Bounds().W()
		TxtOrigX := Txt.Dot.X
		spacing := 100.0
		if len(words)-wordsNum != 1 {
			Txt.Color = colornames.Darkgray
			offset := Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[wordsNum+1]])
			fmt.Fprintln(Txt, words[wordsNum+1])
			myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
			Txt.Dot.X = TxtOrigX
		}
		if !(len(words)-wordsNum == 2 || len(words)-wordsNum == 1) {
			Txt.Color = colornames.Gray
			offset += Txt.Bounds().W()
			Txt.Clear()
			fmt.Fprintln(Txt, wordsJapanese[words[wordsNum+2]])
			fmt.Fprintln(Txt, words[wordsNum+2])
			myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
		}
	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "敵の通常攻撃!!!")
		fmt.Fprint(Txt, "攻撃力:", enemy.EnemySettings[myGame.StageNum].OP, "防御力:", enemy.EnemySettings[myGame.StageNum].DP)
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}

func InitBattleTextMonster(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.PlayingScreen {
		tempWords := words[wordsNum]
		tempWordsSlice := strings.Split(tempWords, "")
		tempWords = strings.Join(tempWordsSlice, " ")
		Txt.Clear()
		Txt.Color = colornames.White
		for i := 0; i < len(words[wordsNum])-1; i++ {
			fmt.Fprint(Txt, "?")
		}
		fmt.Fprintln(Txt, "?")
		Txt.Color = colornames.Gray
		fmt.Fprint(Txt, tempWords[:indexMonster])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[indexMonster:])
		myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt))

	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "敵の通常攻撃!!!")
		fmt.Fprint(Txt, "攻撃力:", enemy.EnemySettings[myGame.StageNum].OP, "防御力:", enemy.EnemySettings[myGame.StageNum].DP)
		myPos.DrawPos(win, Txt, myPos.RoundCenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}
