package main

import "fmt"

func main() {
	score := 11
	text1 := "You are the best!"
	text2 := "Train more"

	if score >= 10 {
		fmt.Println(text1)
	} else {
		fmt.Println(text2)
	}
}
