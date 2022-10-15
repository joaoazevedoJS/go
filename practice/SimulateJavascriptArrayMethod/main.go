package main

import (
	"fmt"
	Array "github.com/Joaoazevedo-JS/GoArray"
)

func PrintMyArray(value int, index int, _ *Array.Array[int]) {
	fmt.Printf("[%d]: %d\n", index, value)
}

func PrintEntries(value Array.Entry, index int, _ *Array.Array[Array.Entry]) {
	fmt.Printf("[%d]: %s\n", value[0], value[1])
}

func FilterForNumbersGreaterThan9(value int, _ int, _ *Array.Array[int]) bool {
	return value >= 10
}

func main() {
	fmt.Println("------------------ CREATE ARRAY AND LIST ARRAY ------------------")

	myArray := Array.New[int](1, 2)

	myArray.ForEach(PrintMyArray)

	fmt.Println("------------------ PUSH TO ARRAY ------------------")

	pushLength := myArray.Push(10, 20, 30, 40)

	myArray.ForEach(PrintMyArray)
	fmt.Printf("New array length: %d\n", pushLength)

	fmt.Println("------------------ UNSHIFT ------------------")

	myArray.Unshift(0, 3, 4, 5)

	myArray.ForEach(PrintMyArray)

	fmt.Println("------------------ POP ------------------")

	popped := myArray.Pop()

	myArray.ForEach(PrintMyArray)
	fmt.Printf("Popped out: %d\n", popped)

	fmt.Println("------------------ FIND INDEX OF ELEMENT ------------------")

	fmt.Printf("Index of 20 is %d\n", myArray.FindIndex(func(value int, _ int, _ *Array.Array[int]) bool {
		return value == 20
	}))

	fmt.Printf("Index of 20000 is %d\n", myArray.FindIndex(func(value int, _ int, _ *Array.Array[int]) bool {
		return value == 20000
	}))

	fmt.Println("------------------ FIND ELEMENT OF ARRAY ------------------")

	elements := Array.New[Array.Array[string]]()

	elements.Push(
		Array.New[string]("a", "b"),
		Array.New[string]("c", "d"),
		Array.New[string]("e", "f"),
	)

	ab, errAB := elements.Find(func(_ Array.Array[string], index int, _ *Array.Array[Array.Array[string]]) bool {
		return index == 0
	})

	gh, errGH := elements.Find(func(_ Array.Array[string], index int, _ *Array.Array[Array.Array[string]]) bool {
		return index == 10
	})

	if errAB == nil {
		fmt.Printf("Element of index 0 is %s\n", ab.Join(", "))
	} else {
		fmt.Println("Not found element of index 0")
	}

	if errGH == nil {
		fmt.Printf("Element of index 10 is %s\n", gh.Join(", "))
	} else {
		fmt.Println("Not found element of index 10")
	}

	fmt.Println("------------------ FILTER ARRAY ------------------")

	myArray = myArray.Filter(FilterForNumbersGreaterThan9)

	myArray.ForEach(PrintMyArray)

	fmt.Println("------------------ LENGTH ------------------")

	fmt.Printf("Length: %d \n", myArray.Length())

	fmt.Println("------------------ CONCAT ARRAY ------------------")

	hundreds := Array.New[int](100, 200, 300)
	thousands := Array.New[int](1000, 2000, 3000)

	myArray = myArray.Concat(hundreds, thousands)

	myArray.ForEach(PrintMyArray)

	fmt.Println("------------------ JOIN ARRAY ------------------")

	fruits := Array.New("Banana", "Orange", "Apple", "Mango")
	decimals := Array.New(10.01, 20.2, 30.3, 40.4, 50.5)

	fmt.Println(myArray.Join())
	fmt.Println(fruits.Join(" and "))
	fmt.Println(decimals.Join(" * "))

	fmt.Println("------------------ TO STRING ------------------")

	fmt.Println(myArray.ToString())
	fmt.Println(fruits.ToString())
	fmt.Println(decimals.ToString())

	fmt.Println("------------------ ENTRIES ------------------")

	entries := fruits.Entries()

	entries.ForEach(PrintEntries)
}
