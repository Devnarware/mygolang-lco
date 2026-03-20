package main



func maxMessages(thresh int) int {
	cost := 100
	i := 0
	
	for(thresh >= cost){
		cost += 100 + i
		i++ 
	}

	return i
}
