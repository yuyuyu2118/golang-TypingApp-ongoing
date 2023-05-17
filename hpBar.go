package main

type PlayerStatus struct {
	playerMaxHP float64
	playerHP    float64
	playerOP    float64
	playerDP    float64
	playerGold  int
	playerJob   string
}

func newPlayerStatus(MaxHP float64, HP float64, OP float64, DP float64, Gold int, Job string) *PlayerStatus {
	return &PlayerStatus{MaxHP, HP, OP, DP, Gold, Job}
}
