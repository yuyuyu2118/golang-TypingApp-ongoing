package battle

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"time"
)

func InitializeQuestion(filePath string) []string {
	words := []string{}
	file, _ := os.Open(filePath)
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		words = append(words, record[2])
	}
	Shuffle(words)
	return words
}

func Shuffle(data []string) {
	n := len(data)
	rand.Seed(time.Now().Unix())
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func ShufflePairs(pairs [][2]string) [][2]string {
	rand.Seed(time.Now().UnixNano())

	pairCount := len(pairs)
	shuffledPairs := make([][2]string, pairCount)

	perm := rand.Perm(pairCount)
	for i, v := range perm {
		shuffledPairs[v] = pairs[i]
	}

	return shuffledPairs
}

func LoadWordsFromCSV(filePath string) (map[string]string, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	pairs := make([][2]string, 0, len(records))
	for _, record := range records {
		if len(record) >= 2 {
			pairs = append(pairs, [2]string{record[1], record[2]})
		}
	}

	shuffledPairs := ShufflePairs(pairs)

	shuffledWords := make(map[string]string)
	for _, pair := range shuffledPairs {
		shuffledWords[pair[1]] = pair[0]
	}

	var tempWords []string
	for value := range shuffledWords {
		tempWords = append(tempWords, value)
	}

	return shuffledWords, tempWords, nil
}
