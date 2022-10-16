package taxgo

func CalculateTax(price float64) float64 {
	if price < 0 {
		return 0
	}

	if price == 0 {
		return 0
	}

	if price >= 1000 && price < 20000 {
		return 10.0
	}

	if price >= 20000 {
		return 20.0
	}

	return 5.0
}
