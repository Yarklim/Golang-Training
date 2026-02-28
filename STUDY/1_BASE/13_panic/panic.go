package panic_study

import "fmt"

func Panic() {
	defer func() {
		panic := recover()
		if panic != nil {
			fmt.Println("Была паника:", panic)
		}
	}()

	// a := 0
	// b := 1 / a
	// fmt.Println(b) // panic: runtime error: integer divide by zero

	slice := []int{1, 2, 3}
	fmt.Println(slice[4]) // panic: runtime error: index out of range [4] with length 3
}
