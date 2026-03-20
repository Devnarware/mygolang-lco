package main

func bulkSend(numMessages int) float64 {
	var fee float64
	for i := range numMessages {
		fee +=   1 + float64(i) * 0.01
	}
	return fee
}