package main

import "fmt"

func main() {
	score := 17
	text1 := "You are the best!"
	text2 := "Train more"
	text3 := "You are the MEGA-BEST!"

	if score > 15 {
		fmt.Println(text3)
	} else if score > 10 {
		fmt.Println(text1)
	} else {
		fmt.Println(text2)
	}
}
