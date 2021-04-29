package yahtzee_scoring

func YahtzeeScoring(dice []int) int {
	sums := make(map[int]int)
	for _, v := range dice {
		sums[v-1] += v
	}
	return max(sums)
}

func max(values map[int]int) int {
	max := 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
