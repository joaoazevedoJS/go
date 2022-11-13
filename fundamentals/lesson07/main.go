package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(multiple(2, 2))
	fmt.Println(multiple(12, 12))
}

// if you have two arguments with same type, you can pass de type only one timme
func sum(a, b int) int {
	return a + b
}

func multiple(a, b int) (int, bool) {
	value := a * b

	if value > 50 {
		return value, true
	}

	return value, false
}
