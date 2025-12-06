package main

import "fmt"

func main() {
	number := 10

	// указатель на переменную number
	pointerNumber := &number
	fmt.Println(pointerNumber) // выводит адрес переменной в памяти

	foo(pointerNumber)

	fmt.Println(number)
}

// функция, которая принимает указатель в качестве аргумента
func foo(pointer *int) {
	result := *pointer + 5

	*pointer += 20 // изменяет переменную number из функции main

	fmt.Println(result)
}
