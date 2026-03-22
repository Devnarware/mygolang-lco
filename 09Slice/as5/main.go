package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCost := make([]float64, 0)
	for _, n := range costs {
		if(n.day == day){
			dayCost = append(dayCost, n.value) ;
		}
	}

	return dayCost 
}
