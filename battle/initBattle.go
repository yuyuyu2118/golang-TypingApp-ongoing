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
		myGame.CurrentGS = BattleTypingV2(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		myGame.CurrentGS = BattleTypingV2(win, player, elapsed)
	} else if player.Job == "化け物" {
		myGame.CurrentGS = BattleTypingV2(win, player, elapsed)
	}
}

func InitSkillBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	InitBattleTextV2Skill(win, myUtil.BasicTxt, elapsed)
	if player.Job == "見習い剣士" {
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "狩人" {
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "モンク" {
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "化け物" {
		myGame.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	}

}