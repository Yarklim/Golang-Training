package structures

import "fmt"

// (u User) - рессивер по значению
func (u User) greeting() {
	fmt.Println("Hello, my name is", u.name)
}

// (u *User) - рессивер по указателю
func (u *User) changeRating(rating float64) {
	ratingSum := u.rating + rating
	if ratingSum < 0.0 || ratingSum > 10.0 {
		fmt.Println("Rating is invalide!")
		return
	}

	u.rating = ratingSum

	fmt.Println("New Rating:", u.rating)
}

func (u *User) changeAge(age int) {
	if age <= 0 {
		return
	}

	u.age = age
}

func Structures() {
	user := NewUser(
		"Yar",
		53,
		"0678596341",
		false,
		8.0,
	)

	user.greeting()
	user.changeRating(1.5)
	user.changeAge(54)

	fmt.Println("User:", user)
	fmt.Println("Name:", user.name)
	fmt.Println("Rating:", user.rating)
	fmt.Println("IsClose before:", user.isClose)

	user.isClose = true
	fmt.Println("IsClose after:", user.isClose)
}
