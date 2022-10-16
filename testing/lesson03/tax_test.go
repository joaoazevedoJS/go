package taxgo

import "testing"

func TestCalculateTax(t *testing.T) {
	price := 500.0
	expectedTax := 5.0

	tax := CalculateTax(price)

	if tax != expectedTax {
		t.Errorf("Expected tax to be %2.f, but got %2.f", expectedTax, tax)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTaxTest struct {
		price       float64
		expectedTax float64
	}

	table := []calcTaxTest{
		{0.0, 0.0},
		{100.0, 5.0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, test := range table {
		tax := CalculateTax(test.price)

		if tax != test.expectedTax {
			t.Errorf("Expected tax to be %2.f, but got %2.f", test.expectedTax, tax)
		}
	}
}
