package myGame

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/faiface/pixel"
)

var (
	//TODO: buttonSliceの名前変えるs
	buttonSlice           = []pixel.Rect{}
	buttonSliceArmor      = []pixel.Rect{}
	buttonSliceAccessory  = []pixel.Rect{}
	buySellSlice          = []pixel.Rect{}
	buySellSliceArmor     = []pixel.Rect{}
	buySellSliceAccessory = []pixel.Rect{}
)

func CsvToSlice(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
