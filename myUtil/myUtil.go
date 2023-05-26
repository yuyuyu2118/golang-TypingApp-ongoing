package myUtil

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/yuyuyu2118/typingGo/myState"
)

func AnyKeyJustPressed(win *pixelgl.Window, keys ...pixelgl.Button) bool {
	for _, key := range keys {
		if win.JustPressed(key) {
			return true
		}
	}
	return false
}

func UpdatePlayingTimer(game myState.GameState, timer *time.Time) {
	if game == myState.BattleEnemyScreen {
		*timer = time.Now()
	}
}

func UpdateEnemyTimer(game myState.GameState, timer *time.Time) {
	if game == myState.PlayingScreen {
		*timer = time.Now()
	}
}
