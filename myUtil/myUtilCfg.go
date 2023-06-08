package myUtil

var saveReset bool = false

func GetSaveReset() bool {
	return saveReset
}

func SetSaveReset(flag bool) {
	saveReset = flag
}

var playerReset bool = false

func GetPlayerReset() bool {
	return playerReset
}

func SetPlayerReset(flag bool) {
	playerReset = flag
}
