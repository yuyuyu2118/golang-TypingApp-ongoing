package main

import (
	"github.com/eiannone/keyboard"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel/pixelgl"
)

// 入力コードを読み出すためのインターフェース
func input() <-chan byte {
	channel := make(chan byte)
	go func() {
		defer close(channel)
		keyboard.Open()
		defer keyboard.Close()

		for {
			char, _, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			channel <- byte(char)
		}
	}()
	return channel
}

func inputPixel(win *pixelgl.Window) <-chan rune {
	channel := make(chan rune)
	go func() {
		defer close(channel)

		for !win.Closed() {
			typed := win.Typed()
			if typed != "" {
				for _, ch := range typed {
					channel <- ch
				}
			}
			win.Update()
		}
	}()
	return channel
}

// キーボード音を出すためのgoroutine
func playSound(s beep.StreamSeekCloser) {
	speaker.Lock()
	s.Seek(0)
	speaker.Unlock()
	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	<-done
}
