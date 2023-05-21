package myUtil

func CheckErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
