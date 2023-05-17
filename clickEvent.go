package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func jobClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState, player *PlayerStatus) GameState {

	if job1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = StageSelect
		player.playerJob = "Warrior"
	} else if job2Button.Contains(mousePos) || win.JustPressed(pixelgl.Key2) {
		currentGameState = StageSelect
		player.playerJob = "Priest"
	} else if job3Button.Contains(mousePos) || win.JustPressed(pixelgl.Key3) {
		currentGameState = StageSelect
		player.playerJob = "Wizard"
	}
	log.Println("YourJob is", player.playerJob)
	return currentGameState
}

func stageClickEvent(win *pixelgl.Window, mousePos pixel.Vec, currentGameState GameState, stage *stageInf) GameState {

	if stage1Button.Contains(mousePos) || win.JustPressed(pixelgl.Key1) {
		currentGameState = PlayingScreen
		log.Println("PlayStage is VS knight")
		stage.stageNum = 1
	}
	log.Println("YourJob is", stage.stageNum)
	return currentGameState
}
