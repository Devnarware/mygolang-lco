package main

func countConnections(groupSize int) int {
	if groupSize == 0 {
		return 0
	}
	ans := groupSize - 1
	for i := groupSize - 2 ; i >= 1; i-- {
		ans += i ;
	}
	return ans
}
