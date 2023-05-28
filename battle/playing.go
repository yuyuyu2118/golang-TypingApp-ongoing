package battle

import (
	"fmt"
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
	fmt.Fprintln(Txt, "> ", words[score])
	myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt))

	offset := Txt.Bounds().W()
	TxtOrigX := Txt.Dot.X
	spacing := 100.0
	if len(words)-score != 1 {
		Txt.Color = colornames.Darkgray
		offset := Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+1])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing, 0)))
		Txt.Dot.X = TxtOrigX
	}
	if !(len(words)-score == 2 || len(words)-score == 1) {
		Txt.Color = colornames.Gray
		offset += Txt.Bounds().W()
		Txt.Clear()
		fmt.Fprintln(Txt, words[score+2])
		myPos.DrawPos(win, Txt, myPos.BottleRoundCenterPos(win, Txt).Add(pixel.V(offset+spacing*2, 0)))
	}
	return elapsed
}

func InitBattleTextV2(win *pixelgl.Window, Txt *text.Text, elapsed time.Duration) time.Duration {

	if myState.CurrentGS == myState.PlayingScreen {
		tempWords := words[score]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, wordsJapanese[words[score]])
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
		fmt.Fprintln(Txt, "敵の通常攻撃!!!")
		fmt.Fprint(Txt, "攻撃力:", enemy.EnemySettings[myGame.StageNum].OP, "防御力:", enemy.EnemySettings[myGame.StageNum].DP)
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))
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
	var tempKanjiString = "魔法力: "
	var tempKanji []rune
	for _, v := range tempKanjiString {
		tempKanji = append(tempKanji, rune(v))
	}
	magic1Txt := myUtil.InitMagicJapanText(myUtil.JapanFontPath, 40, colornames.White, tempKanji)
	magic2Txt := myUtil.InitMagicJapanText(myUtil.JapanFontPath, 45, colornames.White, tempKanji)
	magic3Txt := myUtil.InitMagicJapanText(myUtil.JapanFontPath, 50, colornames.White, tempKanji)
	magic4Txt := myUtil.InitMagicJapanText(myUtil.JapanFontPath, 55, colornames.White, tempKanji)

	if myState.CurrentGS == myState.PlayingScreen {
		Txt.Clear()
		if magicCollectType >= 0 && magicCollectType < 25 {
			Txt.Color = colornames.White
			fmt.Fprintln(magic1Txt, "魔法力:", magicCollectType)
			myPos.DrawPos(win, magic1Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(magic1Txt.TabWidth/2, magic1Txt.LineHeight/2)))
		} else if magicCollectType >= 25 && magicCollectType < 50 {
			magic2Txt.Color = colornames.Yellow
			fmt.Fprintln(magic2Txt, "魔法力:", magicCollectType)
			myPos.DrawPos(win, magic2Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(magic2Txt.TabWidth/2, magic2Txt.LineHeight/2)))
		} else if magicCollectType >= 50 && magicCollectType < 100 {
			magic3Txt.Color = colornames.Orange
			fmt.Fprintln(magic3Txt, "魔法力:", magicCollectType)
			myPos.DrawPos(win, magic3Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(magic3Txt.TabWidth/2, magic3Txt.LineHeight/2)))
		} else if magicCollectType >= 100 {
			magic4Txt.Color = colornames.Red
			fmt.Fprintln(magic4Txt, "魔法力:", magicCollectType)
			myPos.DrawPos(win, magic4Txt, myPos.CenLefPos(win, Txt).Sub(pixel.V(magic4Txt.TabWidth/2, magic4Txt.LineHeight/2)))
		}

		tempWords := words[score]
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, wordsJapanese[words[score]])
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
		tempWords := words[score]
		tempWordsSlice := strings.Split(tempWords, "")
		tempWords = strings.Join(tempWordsSlice, " ")
		Txt.Clear()
		Txt.Color = colornames.White
		for i := 0; i < len(words[score])-1; i++ {
			fmt.Fprint(Txt, "?")
		}
		fmt.Fprintln(Txt, "?")
		Txt.Color = colornames.Gray
		fmt.Fprint(Txt, tempWords[:indexMonster])
		Txt.Color = colornames.White
		fmt.Fprint(Txt, tempWords[indexMonster:])
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))

	} else if myState.CurrentGS == myState.BattleEnemyScreen {
		Txt.Clear()
		Txt.Color = colornames.White
		fmt.Fprintln(Txt, "敵の通常攻撃!!!")
		fmt.Fprint(Txt, "攻撃力:", enemy.EnemySettings[myGame.StageNum].OP, "防御力:", enemy.EnemySettings[myGame.StageNum].DP)
		myPos.DrawPos(win, Txt, myPos.CenPos(win, Txt))
	}

	Txt.Clear()
	Txt.Color = colornames.White
	fmt.Fprintln(Txt, "time = ", elapsed.Milliseconds())
	myPos.DrawPos(win, Txt, myPos.BottleLeftPos(win, Txt))

	return elapsed
}
