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
