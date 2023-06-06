package battle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myPlayer"
	"github.com/yuyuyu2118/typingGo/myPos"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"golang.org/x/image/colornames"
)

var battleTimeBool bool
var startTime time.Time
var DispTimer float64
<<<<<<< HEAD

func InitPlayingBattle(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration) {
=======
>>>>>>> 4715f1e (武器の固有スキルを設定)

	if !battleTimeBool {
		elapsed = time.Since(startTime)
		battleTimeBool = true
		DispTimer = player.AttackTimer - elapsed.Seconds()
	}

	xOffSet := myPos.TopRigPos(win, myUtil.ScreenTxt).X - 80
	yOffSet := myUtil.ScreenTxt.LineHeight - 20
	myUtil.ScreenTxt.Clear()
	myUtil.ScreenTxt.Color = colornames.White

	fmt.Fprintln(myUtil.ScreenTxt, "アタックタイマー", strconv.FormatFloat(DispTimer, 'f', 2, 64))
<<<<<<< HEAD

=======
	yOffSet := myUtil.ScreenTxt.LineHeight
>>>>>>> 4715f1e (武器の固有スキルを設定)
	txtPos := pixel.V(xOffSet, yOffSet)
	tempPosition := pixel.IM.Moved(txtPos)
	myUtil.ScreenTxt.Draw(win, tempPosition)

	if player.Job == "No Job" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingTest(win, player, elapsed)
	} else if player.Job == "見習い剣士" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingRookie(win, player, elapsed, &DispTimer)
	} else if player.Job == "狩人" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingHunter(win, player, elapsed, &DispTimer)
	} else if player.Job == "モンク" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonk(win, player, elapsed, &DispTimer)
	} else if player.Job == "魔法使い" {
		InitBattleTextMagicUser(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMagicUser(win, player, elapsed, &DispTimer)
	} else if player.Job == "化け物" {
		InitBattleTextMonster(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonster(win, player, elapsed, &DispTimer)
	}
}

func InitEnemyBattle(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
	if myGame.StageNum == 0 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 1 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 2 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 3 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 4 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 5 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 6 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 7 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 8 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 9 {
		myState.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	}
}

func InitSkillBattle(win *pixelgl.Window, player *myPlayer.PlayerStatus, elapsed time.Duration) {
	if player.Job == "No Job" {
		myState.CurrentGS = myState.PlayingScreen
	} else if player.Job == "見習い剣士" {
		elapsed := InitBattleTextRookieSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "狩人" {
		elapsed := InitBattleTextHunterSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingHunterSkill(win, player, elapsed)
	} else if player.Job == "モンク" {
		elapsed := InitBattleTextMonkSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonkSkill(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		elapsed := InitBattleTextMagicUserSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMagicUserSkill(win, player, elapsed)
	} else if player.Job == "化け物" {
		elapsed := InitBattleTextMonsterSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonsterSkill(win, player, elapsed)
	}

}
