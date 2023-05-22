package myUtil

var saveReset bool = false

func GetSaveReset() bool {
	return saveReset
}

func SetSaveReset(flag bool) {
	saveReset = flag
}
