package main

import "fmt"

func main() {
	salarios := map[string]int{"João": 1500, "Maria": 3000, "Pedro": 2500}

	fmt.Println(salarios["João"])

	delete(salarios, "João")

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	salarios["João"] = 1800

	fmt.Println(salarios["João"])

	sal := make(map[string]int)

	sal["João"] = 1500
}
