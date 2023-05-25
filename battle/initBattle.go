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
	myGame.CurrentGS = BattleTypingV2(win, player, elapsed)
}
