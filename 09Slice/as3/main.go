package main

func getMessageCosts(messages []string) []float64 {
	msgCost := make([]float64, len(messages)) ;

	for i := 0; i < len(messages); i++ {
		msgCost[i] = float64(len(messages[i])) * 0.01 ;
	}

	return msgCost
}
