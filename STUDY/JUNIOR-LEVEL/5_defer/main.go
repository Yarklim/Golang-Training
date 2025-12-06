package main

import "fmt"

func main() {
	fmt.Println("Start Main")
	// defer выполнится в самом конце кода функции main
	defer func() {
		fmt.Println("DEFER MAIN")
	}()

	foo()

	hello()

	fmt.Println("End Main")
}

func hello() {
	// defer выполнится в самом конце кода функции hello
	defer func() {
		fmt.Println("DEFER HELLO")
	}()

	fmt.Println("Hello!")
}

// Вывод будет в такой последовательности: 0-3-2-1
func foo() {
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()

	fmt.Println("0")
}
