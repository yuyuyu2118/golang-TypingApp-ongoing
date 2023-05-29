package battle

import (
	"log"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myGame"
	"github.com/yuyuyu2118/typingGo/myState"
	"github.com/yuyuyu2118/typingGo/myUtil"
	"github.com/yuyuyu2118/typingGo/player"
)

var battleTimeBool bool
var startTime time.Time

func InitPlayingBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	if !battleTimeBool {
		elapsed = time.Since(startTime)
		battleTimeBool = true
		log.Println("koko", elapsed)
	}
	if player.Job == "見習い剣士" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingRookie(win, player, elapsed)
	} else if player.Job == "狩人" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingHunter(win, player, elapsed)
	} else if player.Job == "モンク" {
		InitBattleTextV2(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonk(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		InitBattleTextMagicUser(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMagicUser(win, player, elapsed)
	} else if player.Job == "化け物" {
		InitBattleTextMonster(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonster(win, player, elapsed)
	}
}

func InitEnemyBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
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
	}
}

func InitSkillBattle(win *pixelgl.Window, player *player.PlayerStatus, elapsed time.Duration) {
	if player.Job == "見習い剣士" {
		InitBattleTextRookieSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingRookieSkill(win, player, elapsed)
	} else if player.Job == "狩人" {
		InitBattleTextHunterSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingHunterSkill(win, player, elapsed)
	} else if player.Job == "モンク" {
		InitBattleTextMonkSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonkSkill(win, player, elapsed)
	} else if player.Job == "魔法使い" {
		InitBattleTextMagicUserSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMagicUserSkill(win, player, elapsed)
	} else if player.Job == "化け物" {
		InitBattleTextMonsterSkill(win, myUtil.ScreenTxt, elapsed)
		myState.CurrentGS = BattleTypingMonsterSkill(win, player, elapsed)
	}

}
