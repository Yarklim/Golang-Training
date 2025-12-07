package main

import "fmt"

// Типизация struct
type User struct {
	Name        string  // ""
	Age         int     // 0
	PhoneNumber string  // ""
	IsClose     bool    // false
	Rating      float64 // 0.0
}

// Конструктор
func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		fmt.Println("Name is invalide!")
		return User{}
	}
	if age <= 0 || age >= 150 {
		fmt.Println("Age is invalide!")
		return User{}
	}
	if phoneNumber == "" {
		fmt.Println("PhoneNumber is invalide!")
		return User{}
	}
	if rating < 0 || rating > 10 {
		fmt.Println("Rating is invalide!")
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}
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
	user := NewUser(
		"Yar",
		53,
		"0678596341",
		false,
		8.0,
	)

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
