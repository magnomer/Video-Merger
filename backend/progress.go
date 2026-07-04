package backend

func LPercentCalculate(current int, total int) int {
	if total <= 0 {
		return 0
	}

	return current * 100 / total
}
