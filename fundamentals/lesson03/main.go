package main

type ID int

var (
	userId   ID     = 0
	userName string = "João Azevedo"
)

func main() {
	println(userId, userName)
}
