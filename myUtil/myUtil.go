package myUtil

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myGame"
)

func AnyKeyJustPressed(win *pixelgl.Window, keys ...pixelgl.Button) bool {
	for _, key := range keys {
		if win.JustPressed(key) {
			return true
		}
	}
	return false
}

func UpdatePlayingTimer(game myGame.GameState, timer *time.Time) {
	if game == myGame.BattleEnemyScreen {
		*timer = time.Now()
	}
}

func UpdateEnemyTimer(game myGame.GameState, timer *time.Time) {
	if game == myGame.PlayingScreen {
		*timer = time.Now()
	}
}
