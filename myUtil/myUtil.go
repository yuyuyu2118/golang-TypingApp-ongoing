package myUtil

import "github.com/faiface/pixel/pixelgl"

func AnyKeyJustPressed(win *pixelgl.Window, keys ...pixelgl.Button) bool {
	for _, key := range keys {
		if win.JustPressed(key) {
			return true
		}
	}
	return false
}
