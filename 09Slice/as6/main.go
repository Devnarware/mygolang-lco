package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// for i := 0; i < len(msg); i++ {
	// 	for j := 0; j < len(badWords); j++ {
	// 		if(msg[i] == badWords[j]){
	// 			return i ;
	// 		}
	// 	}
	// }


	// for i, n := range msg{
	// 	for _, m := range badWords{
	// 		if(n == m){
	// 			return i ;
	// 		}
	// 	}
	// }


	badWordsMap := make(map[string]bool) ;
	for _, word := range badWords{
		badWordsMap[word] = true ;
	}

	for i, v := range msg {
		if badWordsMap[v] {
			return i ;
		}
	}


	return -1
}
