package main

import "fmt"

func yourScore(score int, words []string, collectNum int, missNum int) {
	fmt.Println("正解単語数: ", score, "/", len(words))
	fmt.Println("正解タイプ数: ", collectNum, " ミスタイプ:", missNum)
	n := score
	switch {
	case n <= 10:
		fmt.Println("判定 F")
	case 10 < n && n <= 20:
		fmt.Println("判定 E")
	default:
		fmt.Println("判定 F")
	}
}
