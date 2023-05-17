package main

type GameState int

const (
	StartScreen GameState = iota
	JobSelect
	StageSelect
	Playing
	EndScreen
	TestState
)
