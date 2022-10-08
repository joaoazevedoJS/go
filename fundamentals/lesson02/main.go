package main

import "fmt"

const immutable = "I'm immutable"

var name = "João Azevedo"

var (
	age          int     = 404
	liveInBrazil bool    = true
	money        float64 = 0.10
)

func main() {
	println(immutable)

	// name := value is equal var name string = value
	// it's a shortcut
	country := "live in Brazil"

	if !liveInBrazil {
		// you only can use := only one time, because it's a shortcut to create variables
		country = "not live in Brazil"
	}

	fmt.Printf("Hello, my name is %s, i'm %v years old. I'm %s. I have R$ %.2f", name, age, country, money)
}
