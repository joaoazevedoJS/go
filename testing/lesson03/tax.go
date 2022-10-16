package taxgo

func CalculateTax(price float64) float64 {
	if price == 0 {
		return 0
	}

	if price >= 1000 {
		return 10.0
	}

	return 5.0
}
