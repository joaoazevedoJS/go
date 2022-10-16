package taxgo

import "errors"

func CalculateTax(price float64) (float64, error) {
	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}

	if price == 0 {
		return 0, nil
	}

	if price >= 1000 && price < 20000 {
		return 10.0, nil
	}

	if price >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil
}
