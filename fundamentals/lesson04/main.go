package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for index, value := range myArray {
		fmt.Printf("[%d]: %d\n", index, value)
	}
}
