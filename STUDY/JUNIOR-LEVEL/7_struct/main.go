package main

import "fmt"

type User struct {
	Name        string  // ""
	Age         int     // 0
	PhoneNumber string  // ""
	IsClose     bool    // false
	Rating      float64 // 0.0
}

// (u User) - рессивер по значению
func (u User) Greeting() {
	fmt.Println("Hello, my name is", u.Name)
}

// (u *User) - рессивер по указателю
func (u *User) ChangeRating(rating float64) {
	ratingSum := u.Rating + rating
	if ratingSum < 0.0 || ratingSum > 10.0 {
		fmt.Println("Rating is invalide!")
		return
	}

	u.Rating = ratingSum

	fmt.Println("New Rating:", u.Rating)
}

func (u *User) ChangeAge(age int) {
	if age <= 0 {
		return
	}

	u.Age = age
}

func main() {
	user := User{
		Name:        "Yar",
		Age:         53,
		PhoneNumber: "0678596341",
		IsClose:     false,
		Rating:      8.0,
	}

	user.Greeting()
	user.ChangeRating(1.5)
	user.ChangeAge(54)

	fmt.Println("User:", user)
	fmt.Println("Name:", user.Name)
	fmt.Println("Rating:", user.Rating)
	fmt.Println("IsClose before:", user.IsClose)

	user.IsClose = true
	fmt.Println("IsClose after:", user.IsClose)
}
