package myIo

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
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

func GenerateCSVString(values []string) string {
	quotedValues := make([]string, len(values))
	for i, v := range values {
		quotedValues[i] = `"` + v + `"`
	}
	return strings.Join(quotedValues, ",")
}

func CsvToSlice1Line(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	var tempRecords string
	for _, value := range records {
		tempRecords += value
	}

	return tempRecords
}
