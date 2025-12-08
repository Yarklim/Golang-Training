package conditionals

import "fmt"

func Conditionals() {
	number := 5
	firstVal := number == 5
	secondVal := number > 12

	fmt.Println(firstVal)
	fmt.Println(secondVal)

	score := 17
	text1 := "You are the best!"
	text2 := "Train more"
	text3 := "You are the MEGA BEST!"

	if score > 15 {
		fmt.Println(text3)
	} else if score > 10 {
		fmt.Println(text1)
	} else {
		fmt.Println(text2)
	}

	val := 7

	if val < 6 || val > 16 {
		fmt.Println("You are miss")
	} else {
		fmt.Println("You are good!")
	}

	if val >= 6 && val <= 16 {
		fmt.Println("You are cool!")
	} else {
		fmt.Println("Try again")
	}

	if val != 7 {
		fmt.Println("You are lose")
	} else {
		fmt.Println("You are win!")
	}

	x := (5 < 7 && 2 < 3) || (4 < 3 && 6 < 7)

	fmt.Println(x)
}
