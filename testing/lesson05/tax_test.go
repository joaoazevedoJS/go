package taxgo

import "testing"

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2.5, 0, 500, 1000, 1001, 5000}

	for _, price := range seed {
		f.Add(price)
	}

	f.Fuzz(func(t *testing.T, price float64) {
		tax := CalculateTax(price)

		if price <= 0 && tax != 0 {
			t.Errorf("Received %f but expected 0", tax)
		}

		if price >= 20000 && tax != 20 {
			t.Errorf("Received %f but expected 20", tax)
		}
	})
}
