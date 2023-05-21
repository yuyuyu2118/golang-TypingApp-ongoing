package battle

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"time"
)

func InitializeQuestion() []string {
	words := []string{}
	file, _ := os.Open("assets\\question\\question2_4.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		//wordsMap[record[1]] = record[2]
		words = append(words, record[2])
	}
	Shuffle(words)
	//wordsMap = shuffleMap(wordsMap)
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

// func shuffleMap(data map[string]string) map[string]string {
// 	pairs := make([]struct{ key, value string }, 0, len(data))
// 	for k, v := range data {
// 		pairs = append(pairs, struct{ key, value string }{key: k, value: v})
// 	}
// 	//	log.Println(pairs)

// 	rand.Seed(time.Now().Unix())
// 	rand.Shuffle(len(pairs), func(i, j int) { pairs[i], pairs[j] = pairs[j], pairs[i] })

// 	shuffledData := make(map[string]string)
// 	for _, pair := range pairs {
// 		shuffledData[pair.key] = pair.value
// 		log.Println(pair.key)
// 	}

// 	return shuffledData
// }
