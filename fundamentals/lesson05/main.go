package main

import "fmt"

func main() {
	number := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d, cap=%d %v\n", len(number), cap(number), number)

	fmt.Printf("len=%d, cap=%d %v\n", len(number[:0]), cap(number[:0]), number[:0])

	fmt.Printf("len=%d, cap=%d %v\n", len(number[:4]), cap(number[:4]), number[:4])

	fmt.Printf("len=%d, cap=%d %v\n", len(number[2:]), cap(number[2:]), number[2:])

	number = append(number, 110)

	fmt.Printf("%v\n", number)

	fmt.Printf("len=%d, cap=%d %v\n", len(number[:2]), cap(number[:2]), number[:2])
}
