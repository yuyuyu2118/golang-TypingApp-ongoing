package myIo

import (
	"encoding/csv"
	"log"
	"os"
)

// CsvToSliceはcsvファイルのPathを受け取り、読み込みます。
// csvファイルは2次元スライスの文字列で返されます。
func CsvToSliceAll(path string) [][]string {
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
