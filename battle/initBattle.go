package battle

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

func InitPlayingBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2(win, myUtil.BasicTxt, elapsed)
	if player.Job == "見習い剣士" {
		myGame.CurrentGS = BattleTypingRookie(win, player, elapsed)
	} else if player.Job == "狩人" {
		myGame.CurrentGS = BattleTypingHunter(win, player, elapsed)
	} else if player.Job == "モンク" {
		myGame.CurrentGS = BattleTypingMonk(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		myGame.CurrentGS = BattleTyping(win, player, elapsed)
	} else if player.Job == "化け物" {
		myGame.CurrentGS = BattleTyping(win, player, elapsed)
	}
}

func InitEnemyBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2(win, myUtil.BasicTxt, elapsed)
	if myGame.StageNum == 0 {
		myGame.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 1 {
		myGame.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 2 {
		myGame.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 3 {
		myGame.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	} else if myGame.StageNum == 4 {
		myGame.CurrentGS = BattleTypingEnemySlime(win, player, elapsed)
	}
}

func InitSkillBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	if player.Job == "見習い剣士" {
		InitBattleTextRookieSkill(win, myUtil.BasicTxt, elapsed)
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "狩人" {
		InitBattleTextHunterSkill(win, myUtil.BasicTxt, elapsed)
		myGame.CurrentGS = BattleTypingHunterSkill(win, player, elapsed)
	} else if player.Job == "モンク" {
		InitBattleTextMonkSkill(win, myUtil.BasicTxt, elapsed)
		myGame.CurrentGS = BattleTypingMonkSkill(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		InitBattleTextV2Skill(win, myUtil.BasicTxt, elapsed)
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "化け物" {
		InitBattleTextV2Skill(win, myUtil.BasicTxt, elapsed)
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	}

}
