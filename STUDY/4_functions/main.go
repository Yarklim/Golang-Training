package main

import "fmt"

func main() {
	num1 := 5
	num2 := 10

	result := sum(num1, num2)

	fmt.Println(result)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
