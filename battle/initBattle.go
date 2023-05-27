package battle

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

func InitPlayingBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
	if player.Job == "見習い剣士" {
		myState.CurrentGS = BattleTypingRookie(win, player, elapsed)
	} else if player.Job == "狩人" {
		myState.CurrentGS = BattleTypingHunter(win, player, elapsed)
	} else if player.Job == "モンク" {
		myState.CurrentGS = BattleTypingMonk(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		myState.CurrentGS = BattleTyping(win, player, elapsed)
	} else if player.Job == "化け物" {
		myState.CurrentGS = BattleTyping(win, player, elapsed)
	}
}

func InitEnemyBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2(win, myUtil.BasicTxt, elapsed)
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
	}
}

func InitSkillBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	if player.Job == "見習い剣士" {
		InitBattleTextRookieSkill(win, myUtil.BasicTxt, elapsed)
		myState.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "狩人" {
		InitBattleTextHunterSkill(win, myUtil.BasicTxt, elapsed)
		myState.CurrentGS = BattleTypingHunterSkill(win, player, elapsed)
	} else if player.Job == "モンク" {
		InitBattleTextMonkSkill(win, myUtil.BasicTxt, elapsed)
		myState.CurrentGS = BattleTypingMonkSkill(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		InitBattleTextV2Skill(win, myUtil.BasicTxt, elapsed)
		myState.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "化け物" {
		InitBattleTextV2Skill(win, myUtil.BasicTxt, elapsed)
		myState.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	}

}
