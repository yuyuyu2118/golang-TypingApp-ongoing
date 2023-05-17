package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type GameState int

const (
	StartScreen GameState = iota
	JobSelect
	StageSelect
	PlayingScreen
	EndScreen
	TestState
)

var (
	startLines = []string{
		"This is a TypingBattleGame",
		"\n",
		"START : Press Enter",
	}
)

func initStartScreen(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Darkcyan)
	Txt.Clear()

	lineCenterAlign(win, windowHeightSize, startLines, Txt, "center")
}

func initJobSelect(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	//windowのリセットとテキストの描画
	win.Clear(colornames.Black)

	initJob(win, Txt, windowHeightSize)
}

func initStageSlect(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	win.Clear(colornames.Black)

	initStage(win, Txt, windowHeightSize)
}

func initPlayingScreen(win *pixelgl.Window, Txt *text.Text, windowHeightSize int) {
	win.Clear(colornames.Black)
	Txt.Clear()

	initPlaying(win, Txt, windowHeightSize)
}
