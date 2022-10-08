package main

import (
	"Caculator/Math"
	"fmt"
)

func main() {
	Math.Number01 = 10.0
	Math.Number02 = 10.0

	fmt.Printf("Sum: %2.f \n", Math.Sum())
	fmt.Printf("Subtraction: %2.f \n", Math.Subtraction())
	fmt.Printf("Multiplication: %2.f \n", Math.Multiplication())
	fmt.Printf("Division: %v \n", Math.Division())
}
