package myIo

// func sliceNumCount(s string) map[string]int {
// 	s = "1,0,2,0,0,0,0,0,0,0,,,"
// 	counts := make(map[string]int)
// 	elements := strings.Split(s, ",")

// 	for i := 0; i < 10; i++ {
// 		key := fmt.Sprintf("weapon%d", i)
// 		counts[key] = 0
// 	}

// 	for _, e := range elements {
// 		if e == "" {
// 			continue
// 		}
// 		key := fmt.Sprintf("weapon%s", e)
// 		counts[key]++
// 	}

// 	log.Println(counts) // map[weapon0:8 weapon1:1 weapon2:1 ... weapon9:0]
// }
